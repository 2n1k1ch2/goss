package cluster

import (
	"goss/pkg/parser"
	"goss/pkg/pprof"
	"testing"
	"time"
)

func TestCluserize(t *testing.T) {
	for i := 0; i < 5; i++ {
		go func() {
			time.Sleep(time.Second * 5)
		}()
	}
	time.Sleep(time.Second)
	dump := pprof.CaptureAll()

	goroutines, err := parser.Normalize(dump)

	if err != nil {
		t.Fatal(err)
	}

	cluster := Clusterize(goroutines)
	if len(cluster) == 0 {
		t.Fatal("cluster is empty")
	}

}
