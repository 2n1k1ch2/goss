package sdk

import (
	"context"
	"reflect"
	"runtime"
	"runtime/pprof"
	"sync"
)

func Go(ctx context.Context, name string, fn func(ctx context.Context)) {
	if name == "" {
		name = runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	}
	labels := pprof.Labels("goss.name", name)
	go pprof.Do(ctx, labels, func(ctx context.Context) {
		fn(ctx)
	})
}
func GoGroup(ctx context.Context, name string, wg *sync.WaitGroup, fn func(ctx context.Context)) {
	if name == "" {
		name = runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	}
	wg.Add(1)
	go pprof.Do(ctx, pprof.Labels("goss.name", name), func(ctx context.Context) {
		defer wg.Done()
		fn(ctx)
	})
}
