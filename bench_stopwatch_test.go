package main

import (
	"testing"
	"time"
)

type Recorder interface {
	Record(d time.Duration)
}

type timer struct {
	x time.Duration
}

func (t *timer) Start() Stopwatch {
	return Stopwatch{started: time.Now(), recorder: t}
}

func (t *timer) Record(d time.Duration) {
	t.x = d
}

type Stopwatch struct {
	started  time.Time
	recorder Recorder
}

func (s Stopwatch) Stop() {
	s.recorder.Record(time.Since(s.started))
}

func BenchmarkStopwatch(b *testing.B) {
	timer := &timer{}
	for i := 0; i < b.N; i++ {
		stopwatch := timer.Start()
		defer stopwatch.Stop()
	}
}
