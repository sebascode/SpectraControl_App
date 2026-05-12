package main

import (
	"encoding/binary"
	"testing"
)

func TestBuildHueStreamPacket_HeaderLayout(t *testing.T) {
	cfgID := "6eaf3b98-418d-48f3-89e4-a374cf9ef290" // 36 chars
	channels := []chanColor{
		{id: 0, r: 0x1234, g: 0x5678, b: 0x9abc},
	}
	pkt := buildHueStreamPacket(0x42, cfgID, channels)

	if got := string(pkt[:9]); got != "HueStream" {
		t.Errorf("magic: got %q, want %q", got, "HueStream")
	}
	if pkt[9] != 0x02 {
		t.Errorf("version major: got 0x%02x, want 0x02", pkt[9])
	}
	if pkt[10] != 0x00 {
		t.Errorf("version minor: got 0x%02x, want 0x00", pkt[10])
	}
	if pkt[11] != 0x42 {
		t.Errorf("sequence: got 0x%02x, want 0x42", pkt[11])
	}
	if pkt[14] != 0x00 {
		t.Errorf("color space: got 0x%02x, want 0x00 (RGB)", pkt[14])
	}
	if got := string(pkt[16:52]); got != cfgID {
		t.Errorf("config UUID: got %q, want %q", got, cfgID)
	}
}

func TestBuildHueStreamPacket_ChannelPayload(t *testing.T) {
	cfgID := "6eaf3b98-418d-48f3-89e4-a374cf9ef290"
	channels := []chanColor{
		{id: 0, r: 0x1234, g: 0x5678, b: 0x9abc},
		{id: 7, r: 0xffff, g: 0x0000, b: 0x8000},
	}
	pkt := buildHueStreamPacket(0, cfgID, channels)

	wantLen := 52 + len(channels)*7
	if len(pkt) != wantLen {
		t.Fatalf("packet length: got %d, want %d", len(pkt), wantLen)
	}

	for i, ch := range channels {
		base := 52 + i*7
		if pkt[base] != ch.id {
			t.Errorf("channel %d id: got %d, want %d", i, pkt[base], ch.id)
		}
		if r := binary.BigEndian.Uint16(pkt[base+1:]); r != ch.r {
			t.Errorf("channel %d R: got 0x%04x, want 0x%04x", i, r, ch.r)
		}
		if g := binary.BigEndian.Uint16(pkt[base+3:]); g != ch.g {
			t.Errorf("channel %d G: got 0x%04x, want 0x%04x", i, g, ch.g)
		}
		if b := binary.BigEndian.Uint16(pkt[base+5:]); b != ch.b {
			t.Errorf("channel %d B: got 0x%04x, want 0x%04x", i, b, ch.b)
		}
	}
}

func TestBuildHueStreamPacket_NoChannels(t *testing.T) {
	cfgID := "6eaf3b98-418d-48f3-89e4-a374cf9ef290"
	pkt := buildHueStreamPacket(0, cfgID, nil)
	if len(pkt) != 52 {
		t.Errorf("empty channels packet length: got %d, want 52", len(pkt))
	}
}

func TestLerpColors_Endpoints(t *testing.T) {
	from := []chanColor{{id: 1, r: 0, g: 0, b: 0}}
	to := []chanColor{{id: 1, r: 0xffff, g: 0xffff, b: 0xffff}}

	at0 := lerpColors(from, to, 0)
	if at0[0].r != 0 || at0[0].g != 0 || at0[0].b != 0 {
		t.Errorf("t=0: got (%v,%v,%v), want (0,0,0)", at0[0].r, at0[0].g, at0[0].b)
	}
	at1 := lerpColors(from, to, 1)
	if at1[0].r != 0xffff || at1[0].g != 0xffff || at1[0].b != 0xffff {
		t.Errorf("t=1: got (%v,%v,%v), want (ffff,ffff,ffff)", at1[0].r, at1[0].g, at1[0].b)
	}
}

func TestLerpColors_Midway(t *testing.T) {
	from := []chanColor{{id: 1, r: 0, g: 0, b: 0}}
	to := []chanColor{{id: 1, r: 1000, g: 2000, b: 3000}}
	mid := lerpColors(from, to, 0.5)
	if mid[0].r != 500 || mid[0].g != 1000 || mid[0].b != 1500 {
		t.Errorf("midway: got (%v,%v,%v), want (500,1000,1500)",
			mid[0].r, mid[0].g, mid[0].b)
	}
}

func TestLerpColors_UnmatchedChannelPassesThrough(t *testing.T) {
	// Channel 1 in `from`, but `to` asks for channel 9 — should return channel 9 as-is.
	from := []chanColor{{id: 1, r: 100, g: 100, b: 100}}
	to := []chanColor{{id: 9, r: 500, g: 500, b: 500}}
	out := lerpColors(from, to, 0.5)
	if out[0].id != 9 || out[0].r != 500 || out[0].g != 500 || out[0].b != 500 {
		t.Errorf("unmatched: got id=%d (%v,%v,%v), want id=9 (500,500,500)",
			out[0].id, out[0].r, out[0].g, out[0].b)
	}
}

func TestSmoothstep_Endpoints(t *testing.T) {
	if got := smoothstep(0); got != 0 {
		t.Errorf("smoothstep(0): got %v, want 0", got)
	}
	if got := smoothstep(1); got != 1 {
		t.Errorf("smoothstep(1): got %v, want 1", got)
	}
	// Midpoint of smoothstep is exactly 0.5 (3·0.5²−2·0.5³ = 0.75−0.25)
	if got := smoothstep(0.5); got != 0.5 {
		t.Errorf("smoothstep(0.5): got %v, want 0.5", got)
	}
}

func TestSmoothstep_Monotonic(t *testing.T) {
	prev := smoothstep(0)
	for i := 1; i <= 100; i++ {
		cur := smoothstep(float64(i) / 100)
		if cur < prev {
			t.Errorf("not monotonic at %d: %v < %v", i, cur, prev)
		}
		prev = cur
	}
}
