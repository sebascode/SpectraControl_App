// Global brightness state. Applied to both screen sync and dynamic scenes.
// HTTP path uses the "bri" field in the Hue PUT body. DTLS path scales R/G/B
// before encoding (HueStream v2 has no separate brightness channel).
package main

import (
	"encoding/json"
	"net/http"
	"sync/atomic"
)

var globalBri atomic.Uint32 // stored as uint8 0..255 widened to uint32

func init() {
	globalBri.Store(200)
}

func getBri() uint8 {
	return uint8(globalBri.Load())
}

func setBri(v int) {
	if v < 1 {
		v = 1
	}
	if v > 255 {
		v = 255
	}
	globalBri.Store(uint32(v))
}

func handleGetBrightness(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, map[string]int{"value": int(getBri())})
}

func handleSetBrightness(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Value int `json:"value"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeErr(w, http.StatusBadRequest, err.Error())
		return
	}
	setBri(body.Value)
	writeJSON(w, map[string]int{"value": int(getBri())})
}
