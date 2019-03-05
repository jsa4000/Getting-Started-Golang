package mngmt

import (
	"runtime"
)

// NewRuntime fetches runtime statistics
func NewRuntime() *Runtime {
	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)
	return &Runtime{
		NumGoroutine: runtime.NumGoroutine(),
		Alloc:        rtm.Alloc,
		TotalAlloc:   rtm.TotalAlloc,
		Sys:          rtm.Sys,
		Mallocs:      rtm.Mallocs,
		Frees:        rtm.Frees,
		LiveObjects:  rtm.Mallocs - rtm.Frees,
		PauseTotalNs: rtm.PauseTotalNs,
		NumGC:        rtm.NumGC,
	}
}
