package main

import (
	"errors"
	"testing"
)

func TestCheckHueBody(t *testing.T) {
	cases := []struct {
		name string
		body string
		want error
	}{
		{"ok map", `{"1":{"name":"x"}}`, nil},
		{"ok success array", `[{"success":{"/lights/1/state/on":true}}]`, nil},
		{"unauthorized", `[{"error":{"type":1,"address":"/","description":"unauthorized user"}}]`, errHueUnauthorized},
		{"empty", ``, nil},
	}
	for _, c := range cases {
		err := checkHueBody([]byte(c.body))
		if !errors.Is(err, c.want) {
			t.Errorf("%s: got %v, want %v", c.name, err, c.want)
		}
	}
	if err := checkHueBody([]byte(`[{"error":{"type":3,"description":"resource not available"}}]`)); err == nil || errors.Is(err, errHueUnauthorized) {
		t.Errorf("type 3 should be a generic error, got %v", err)
	}
}
