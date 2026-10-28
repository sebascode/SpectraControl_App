package main

import (
	"sync"
	"testing"
	"time"
)

func TestMapV2Update_LightOnOff(t *testing.T) {
	on := true
	u := hueV2Update{
		IDV1: "/lights/3",
		Type: "light",
		On:   &struct{ On bool `json:"on"` }{On: on},
	}
	ev, ok := mapV2Update(u)
	if !ok {
		t.Fatal("expected ok=true for on update")
	}
	if ev.Type != "light" || ev.ID != "3" {
		t.Errorf("wrong type/id: got %q %q", ev.Type, ev.ID)
	}
	if ev.On == nil || *ev.On != true {
		t.Errorf("on not set or wrong: %+v", ev.On)
	}
}

func TestMapV2Update_GroupBrightnessScale(t *testing.T) {
	// v2 brightness es 0–100, v1 es 1–254. 50% en v2 ≈ 127 ó 128 en v1.
	u := hueV2Update{
		IDV1: "/groups/2",
		Type: "grouped_light",
		Dimming: &struct {
			Brightness float64 `json:"brightness"`
		}{Brightness: 50},
	}
	ev, ok := mapV2Update(u)
	if !ok {
		t.Fatal("expected ok=true")
	}
	if ev.Type != "group" || ev.ID != "2" {
		t.Errorf("wrong type/id: got %q %q", ev.Type, ev.ID)
	}
	if ev.Bri == nil {
		t.Fatal("bri not set")
	}
	// 50/100 * 253 = 126.5 → round → 127, +1 = 128
	if *ev.Bri != 128 {
		t.Errorf("bri scale: got %d, want 128", *ev.Bri)
	}
}

func TestMapV2Update_BrightnessClampsToRange(t *testing.T) {
	cases := []struct {
		pct  float64
		want int
	}{
		{0, 1},     // 0% → 1
		{100, 254}, // 100% → 254
	}
	for _, c := range cases {
		u := hueV2Update{
			IDV1: "/lights/1",
			Dimming: &struct {
				Brightness float64 `json:"brightness"`
			}{Brightness: c.pct},
		}
		ev, _ := mapV2Update(u)
		if ev.Bri == nil || *ev.Bri != c.want {
			got := -1
			if ev.Bri != nil {
				got = *ev.Bri
			}
			t.Errorf("brightness %.1f%%: got %d, want %d", c.pct, got, c.want)
		}
	}
}

func TestMapV2Update_ColorXY(t *testing.T) {
	u := hueV2Update{
		IDV1: "/lights/4",
		Color: &struct {
			XY struct {
				X float64 `json:"x"`
				Y float64 `json:"y"`
			} `json:"xy"`
		}{XY: struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		}{X: 0.31273, Y: 0.32902}},
	}
	ev, ok := mapV2Update(u)
	if !ok {
		t.Fatal("expected ok=true")
	}
	if ev.XY == nil {
		t.Fatal("xy not set")
	}
	// Redondeo a 4 decimales.
	if ev.XY[0] != 0.3127 || ev.XY[1] != 0.329 {
		t.Errorf("xy round: got [%f, %f], want [0.3127, 0.329]", ev.XY[0], ev.XY[1])
	}
}

func TestMapV2Update_RejectsMissingIDV1(t *testing.T) {
	u := hueV2Update{IDV1: "", Type: "light"}
	if _, ok := mapV2Update(u); ok {
		t.Error("expected ok=false when id_v1 empty")
	}
}

func TestMapV2Update_RejectsUnknownResource(t *testing.T) {
	u := hueV2Update{IDV1: "/zigbee_connectivity/abc", Type: "zigbee_connectivity"}
	if _, ok := mapV2Update(u); ok {
		t.Error("expected ok=false for non light/group id_v1")
	}
}

func TestMapV2Update_RejectsEmptyPayload(t *testing.T) {
	// id_v1 válido pero sin on/dimming/color/ct → nada útil.
	u := hueV2Update{IDV1: "/lights/9", Type: "light"}
	if _, ok := mapV2Update(u); ok {
		t.Error("expected ok=false when no useful fields present")
	}
}

func TestDispatchSSEPayload_BroadcastsToHub(t *testing.T) {
	// Reemplazamos el hub global por uno limpio durante el test.
	orig := bridgeStateHub
	bridgeStateHub = newStateHub()
	defer func() { bridgeStateHub = orig }()

	ch := bridgeStateHub.subscribe()
	defer bridgeStateHub.unsubscribe(ch)

	payload := `[{
		"creationtime": "2026-05-11T20:00:00Z",
		"data": [
			{"id":"abc","id_v1":"/lights/5","type":"light","on":{"on":true},"dimming":{"brightness":40}},
			{"id":"def","id_v1":"/groups/2","type":"grouped_light","on":{"on":false}}
		],
		"id": "evt-1",
		"type": "update"
	}]`
	dispatchSSEPayload(payload)

	// Esperamos dos eventos. Drenar con timeout.
	got := make([]stateEvent, 0, 2)
	deadline := time.After(500 * time.Millisecond)
	for len(got) < 2 {
		select {
		case ev := <-ch:
			got = append(got, ev)
		case <-deadline:
			t.Fatalf("timeout: received %d events, want 2", len(got))
		}
	}
	if got[0].Type != "light" || got[0].ID != "5" || got[0].On == nil || !*got[0].On {
		t.Errorf("first event mismatch: %+v", got[0])
	}
	if got[1].Type != "group" || got[1].ID != "2" || got[1].On == nil || *got[1].On {
		t.Errorf("second event mismatch: %+v", got[1])
	}
}

func TestStateHub_DropsOnSlowConsumer(t *testing.T) {
	// Si el consumer no drena, el hub debe descartar eventos en vez de bloquear.
	hub := newStateHub()
	ch := hub.subscribe()
	defer hub.unsubscribe(ch)

	// Llenar el buffer (32) y empujar uno más; broadcast no debe bloquearse.
	done := make(chan struct{})
	go func() {
		on := true
		for i := 0; i < 100; i++ {
			hub.broadcast(stateEvent{Type: "light", ID: "1", On: &on})
		}
		close(done)
	}()
	select {
	case <-done:
		// ok
	case <-time.After(time.Second):
		t.Fatal("broadcast blocked on slow consumer")
	}
}

func TestStateHub_UnsubscribeClosesChannel(t *testing.T) {
	hub := newStateHub()
	ch := hub.subscribe()

	// Hay que evitar que dos goroutines hagan unsubscribe a la vez sobre el
	// mismo canal — el test solo verifica el camino feliz.
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for range ch {
			// drena
		}
	}()
	hub.unsubscribe(ch)
	wg.Wait() // si no se cierra, esto cuelga el test
}
