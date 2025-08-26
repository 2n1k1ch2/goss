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

func get_stack() (int, []byte) {

	buf := make([]byte, 1024*1024) // 1MB обычно достаточно

	for {
		n := runtime.Stack(buf, true) // all = true
		if n < len(buf) {
			return n, buf[:n]
		}
		// Увеличиваем если не хватило
		buf = make([]byte, 2*len(buf))
	}
}
func CaptureAll() *GoroutineDump {

	n, buf := get_stack()
	raw := string(buf[:n])

	stacks := strings.Split(raw, "\n")

	return &GoroutineDump{
		Timestamp: time.Now(),
		Stacks:    stacks,
	}
}
