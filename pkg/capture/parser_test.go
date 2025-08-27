package capture

import (
	"goss/pkg/pprof"
	"testing"
	"time"
)

func TestCapture(t *testing.T) {
	for i := 0; i < 3; i++ {
		go func() {
			time.Sleep(time.Second * 3)
		}()
	}
	time.Sleep(time.Second)
	dump := pprof.CaptureAll()

	goroutines, err := Normalize(dump)
	if err != nil {
		t.Fatal(err)
	}

	if len(goroutines) == 0 {

		t.Fatal("Goroutines trace is empty")
	}

}
