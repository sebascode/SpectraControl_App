package main

import "testing"

func mkScene(kfs []sceneKeyframe) *scenePreset {
	return &scenePreset{Duration: 1, Loop: true, Keyframes: kfs}
}

func TestInterpolateScene_AtKeyframes(t *testing.T) {
	s := mkScene([]sceneKeyframe{
		{T: 0.0, Color: [3]uint8{0, 0, 0}},
		{T: 0.5, Color: [3]uint8{100, 150, 200}},
		{T: 1.0, Color: [3]uint8{255, 255, 255}},
	})

	cases := []struct {
		phase      float64
		wantR      uint8
		wantG      uint8
		wantB      uint8
	}{
		{0.0, 0, 0, 0},
		{0.5, 100, 150, 200},
		// phase 1.0 wraps to 0.0 → first keyframe
		{1.0, 0, 0, 0},
	}
	for _, c := range cases {
		r, g, b := interpolateScene(s, c.phase)
		if r != c.wantR || g != c.wantG || b != c.wantB {
			t.Errorf("phase %v: got (%d,%d,%d), want (%d,%d,%d)",
				c.phase, r, g, b, c.wantR, c.wantG, c.wantB)
		}
	}
}

func TestInterpolateScene_Midway(t *testing.T) {
	s := mkScene([]sceneKeyframe{
		{T: 0.0, Color: [3]uint8{0, 0, 0}},
		{T: 1.0, Color: [3]uint8{200, 100, 50}},
	})
	r, g, b := interpolateScene(s, 0.5)
	// linear: (100, 50, 25)
	if r != 100 || g != 50 || b != 25 {
		t.Errorf("midway lerp: got (%d,%d,%d), want (100,50,25)", r, g, b)
	}
}

func TestInterpolateScene_WrapsNegativePhase(t *testing.T) {
	s := mkScene([]sceneKeyframe{
		{T: 0.0, Color: [3]uint8{10, 20, 30}},
		{T: 1.0, Color: [3]uint8{200, 200, 200}},
	})
	// -0.75 → +0.25 after wrap
	gotR, gotG, gotB := interpolateScene(s, -0.75)
	wantR, wantG, wantB := interpolateScene(s, 0.25)
	if gotR != wantR || gotG != wantG || gotB != wantB {
		t.Errorf("negative wrap: got (%d,%d,%d), want (%d,%d,%d)",
			gotR, gotG, gotB, wantR, wantG, wantB)
	}
}

func TestInterpolateScene_WrapsLargePhase(t *testing.T) {
	s := mkScene([]sceneKeyframe{
		{T: 0.0, Color: [3]uint8{10, 20, 30}},
		{T: 1.0, Color: [3]uint8{200, 200, 200}},
	})
	// 3.25 → 0.25 after wrap
	gotR, gotG, gotB := interpolateScene(s, 3.25)
	wantR, wantG, wantB := interpolateScene(s, 0.25)
	if gotR != wantR || gotG != wantG || gotB != wantB {
		t.Errorf("large wrap: got (%d,%d,%d), want (%d,%d,%d)",
			gotR, gotG, gotB, wantR, wantG, wantB)
	}
}

func TestInterpolateScene_EmptyKeyframesReturnsBlack(t *testing.T) {
	s := mkScene(nil)
	r, g, b := interpolateScene(s, 0.5)
	if r != 0 || g != 0 || b != 0 {
		t.Errorf("empty keyframes: got (%d,%d,%d), want (0,0,0)", r, g, b)
	}
}

func TestInterpolateScene_SingleKeyframeIsConstant(t *testing.T) {
	s := mkScene([]sceneKeyframe{
		{T: 0.5, Color: [3]uint8{77, 88, 99}},
	})
	for _, phase := range []float64{0.0, 0.25, 0.5, 0.75, 0.999} {
		r, g, b := interpolateScene(s, phase)
		if r != 77 || g != 88 || b != 99 {
			t.Errorf("single keyframe phase %v: got (%d,%d,%d), want (77,88,99)",
				phase, r, g, b)
		}
	}
}
