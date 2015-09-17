package timing

import (
	"strings"
	"testing"
	"time"
)

func TestTiming(t *testing.T) {
	var r Recorder

	if s := r.Measurements.String(); s != "" {
		t.Fatalf("expected empty string, got %q.", s)
	}

	func() {
		defer r.Record("func1", time.Now())
		time.Sleep(100 * time.Millisecond)
	}()

	if len(r.Measurements) != 1 {
		t.Fatalf("expected 1 measurement, got %d.", len(r.Measurements))
	}

	if l := len(r.GetTakingLongerThan(10 * time.Millisecond)); l != 1 {
		t.Fatalf("expected 1 measurement taking longer than 10ms, got %d.", l)
	}

	if l := len(r.GetTakingLongerThan(200 * time.Millisecond)); l != 0 {
		t.Fatalf("expected 0 measurements taking longer than 200ms, got %d.", l)
	}

	if idx := strings.Index(r.Measurements.String(), ","); idx >= 0 {
		t.Fatalf("expected no comma in the string, found one at index %d.", idx)
	}

	func() {
		defer r.Record("func1", time.Now())
		time.Sleep(20 * time.Millisecond)
	}()

	if idx := strings.Index(r.Measurements.String(), ","); idx < 0 {
		t.Fatalf("expected a comma in the string, found none: %s", r.Measurements.String())
	}

	t.Logf("measurements = %s", r.Measurements)
}
