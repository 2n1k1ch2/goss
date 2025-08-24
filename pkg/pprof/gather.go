package pprof

import (
	"runtime"
	"strings"
	"time"
)

type GoroutineDump struct {
	Timestamp time.Time
	Stacks    []string //
}

func CaptureAll() *GoroutineDump {
	buf := make([]byte, 1<<20) // 1 MB буфер
	n := runtime.Stack(buf, true)
	raw := string(buf[:n])

	stacks := strings.Split(raw, "\n\n")

	return &GoroutineDump{
		Timestamp: time.Now(),
		Stacks:    stacks,
	}
}
