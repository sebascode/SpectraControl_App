package main

import "testing"

func TestSetBri_ClampsLow(t *testing.T) {
	setBri(0)
	if got := getBri(); got != 1 {
		t.Errorf("setBri(0): got %d, want 1 (clamp)", got)
	}
	setBri(-50)
	if got := getBri(); got != 1 {
		t.Errorf("setBri(-50): got %d, want 1 (clamp)", got)
	}
}

func TestSetBri_ClampsHigh(t *testing.T) {
	setBri(256)
	if got := getBri(); got != 255 {
		t.Errorf("setBri(256): got %d, want 255 (clamp)", got)
	}
	setBri(999)
	if got := getBri(); got != 255 {
		t.Errorf("setBri(999): got %d, want 255 (clamp)", got)
	}
}

func TestSetBri_InRangePassesThrough(t *testing.T) {
	cases := []int{1, 50, 128, 200, 255}
	for _, c := range cases {
		setBri(c)
		if got := getBri(); int(got) != c {
			t.Errorf("setBri(%d): got %d, want %d", c, got, c)
		}
	}
}
