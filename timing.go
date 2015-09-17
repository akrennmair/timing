package timing

import (
	"bytes"
	"fmt"
	"time"
)

// Measurement describes a single measurement, containing its name and duration.
type Measurement struct {
	Name     string
	Duration time.Duration
}

type Measurements []Measurement

// Recorder contains all recorded measurements.
type Recorder struct {
	Measurements Measurements
}

// Record records a measurement with a particular name. The duration is
// determined through the difference between startTime and the current time.
func (r *Recorder) Record(name string, startTime time.Time) {
	m := Measurement{
		Name:     name,
		Duration: time.Since(startTime),
	}

	r.Measurements = append(r.Measurements, m)
}

// GetTakingLongerThan returns all measurements that have taken longer
// than the duration provided.
func (r *Recorder) GetTakingLongerThan(d time.Duration) (ms Measurements) {
	for _, m := range r.Measurements {
		if m.Duration > d {
			ms = append(ms, m)
		}
	}
	return
}

func (ms Measurements) String() string {
	if len(ms) < 1 {
		return ""
	}
	var output bytes.Buffer
	output.WriteString(ms[0].String())
	for _, m := range ms[1:] {
		output.WriteString(", ")
		output.WriteString(m.String())
	}
	return output.String()
}

func (m Measurement) String() string {
	return fmt.Sprintf("%s: %s", m.Name, m.Duration)
}
