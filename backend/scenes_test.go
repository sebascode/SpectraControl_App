package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/go-chi/chi/v5"
)

func mkScene(kfs []sceneKeyframe) *scenePreset {
	return &scenePreset{Duration: 1, Loop: true, Keyframes: kfs}
}

// withTempUserScenesDir aísla cada test del config dir real apuntando
// XDG_CONFIG_HOME a un tmpdir. userScenesDir() respeta esa variable vía
// os.UserConfigDir(), así que todos los reads/writes caen en el tmpdir.
func withTempUserScenesDir(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", dir)
	return filepath.Join(dir, "spectracontrol", "scenes")
}

func validScenePreset() scenePreset {
	return scenePreset{
		ID:       "my-scene",
		Name:     "My Scene",
		Duration: 5,
		Loop:     true,
		Stagger:  0.1,
		Keyframes: []sceneKeyframe{
			{T: 0, Color: [3]uint8{255, 0, 0}},
			{T: 1, Color: [3]uint8{0, 0, 255}},
		},
	}
}

func postScene(t *testing.T, body any) *httptest.ResponseRecorder {
	t.Helper()
	buf, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	req := httptest.NewRequest(http.MethodPost, "/api/scenes", bytes.NewReader(buf))
	rec := httptest.NewRecorder()
	handleCreateScene(rec, req)
	return rec
}

func TestCreateScene_Valid(t *testing.T) {
	userDir := withTempUserScenesDir(t)
	rec := postScene(t, validScenePreset())
	if rec.Code != http.StatusOK {
		t.Fatalf("got %d, want 200; body=%s", rec.Code, rec.Body.String())
	}
	path := filepath.Join(userDir, "my-scene.json")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("expected file at %s: %v", path, err)
	}
	var s scenePreset
	if err := json.Unmarshal(data, &s); err != nil {
		t.Fatalf("file is not valid JSON: %v", err)
	}
	if s.Source != "user" {
		t.Errorf("source should be forced to 'user', got %q", s.Source)
	}
	if s.ID != "my-scene" || s.Name != "My Scene" {
		t.Errorf("round-trip mismatch: %+v", s)
	}
}

func TestCreateScene_RejectsBadInput(t *testing.T) {
	withTempUserScenesDir(t)
	cases := []struct {
		name   string
		mutate func(*scenePreset)
	}{
		{"empty id", func(s *scenePreset) { s.ID = "" }},
		{"id with slash", func(s *scenePreset) { s.ID = "foo/bar" }},
		{"id with dot", func(s *scenePreset) { s.ID = "foo.bar" }},
		{"id with uppercase", func(s *scenePreset) { s.ID = "MyScene" }},
		{"empty name", func(s *scenePreset) { s.Name = "  " }},
		{"zero duration", func(s *scenePreset) { s.Duration = 0 }},
		{"negative duration", func(s *scenePreset) { s.Duration = -1 }},
		{"one keyframe", func(s *scenePreset) { s.Keyframes = s.Keyframes[:1] }},
		{"keyframe t below 0", func(s *scenePreset) { s.Keyframes[0].T = -0.1 }},
		{"keyframe t above 1", func(s *scenePreset) { s.Keyframes[1].T = 1.1 }},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s := validScenePreset()
			c.mutate(&s)
			rec := postScene(t, s)
			if rec.Code != http.StatusBadRequest {
				t.Fatalf("got %d, want 400; body=%s", rec.Code, rec.Body.String())
			}
		})
	}
}

func TestCreateScene_OverwriteAllowed(t *testing.T) {
	userDir := withTempUserScenesDir(t)
	postScene(t, validScenePreset())

	updated := validScenePreset()
	updated.Name = "Renamed"
	rec := postScene(t, updated)
	if rec.Code != http.StatusOK {
		t.Fatalf("second write should succeed; got %d", rec.Code)
	}
	data, _ := os.ReadFile(filepath.Join(userDir, "my-scene.json"))
	var s scenePreset
	json.Unmarshal(data, &s)
	if s.Name != "Renamed" {
		t.Errorf("expected updated name, got %q", s.Name)
	}
}

func deleteScene(t *testing.T, id string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(http.MethodDelete, "/api/scenes/"+id, nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", id)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rec := httptest.NewRecorder()
	handleDeleteScene(rec, req)
	return rec
}

func TestDeleteScene_RemovesUserFile(t *testing.T) {
	userDir := withTempUserScenesDir(t)
	postScene(t, validScenePreset())

	rec := deleteScene(t, "my-scene")
	if rec.Code != http.StatusOK {
		t.Fatalf("got %d, want 200; body=%s", rec.Code, rec.Body.String())
	}
	if _, err := os.Stat(filepath.Join(userDir, "my-scene.json")); !os.IsNotExist(err) {
		t.Errorf("file should be deleted, stat err=%v", err)
	}
}

func TestDeleteScene_MissingReturns404(t *testing.T) {
	withTempUserScenesDir(t)
	rec := deleteScene(t, "nonexistent")
	if rec.Code != http.StatusNotFound {
		t.Fatalf("got %d, want 404; body=%s", rec.Code, rec.Body.String())
	}
}

func TestDeleteScene_RejectsBadID(t *testing.T) {
	withTempUserScenesDir(t)
	rec := deleteScene(t, "../etc/passwd")
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("got %d, want 400; body=%s", rec.Code, rec.Body.String())
	}
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
