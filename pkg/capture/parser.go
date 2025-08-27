package capture

import (
	"errors"
	"goss/pkg/pprof"
	"regexp"
	"strings"
)

type Goroutine []string

func Normalize(gd *pprof.GoroutineDump) ([]Goroutine, error) {
	if gd == nil {
		return nil, errors.New("gd is empty")
	}
	var Goroutines []Goroutine
	var goroutine Goroutine
	for _, v := range gd.Stacks {

		if v == "" {
			continue
		}

		if strings.Contains(v, "created by") {
			continue
		}
		v = strings.Replace(v, "+0x", "", -1)
		if strings.Contains(v, "(") {
			v = func(string) string {
				re := regexp.MustCompile(`\([^()]*\)$`)
				return re.ReplaceAllString(v, "")
			}(v)

		}
		if strings.Contains(v, "goroutine") {

			re := regexp.MustCompile(`^goroutine\s+\d+`)
			v = re.ReplaceAllString(v, "")
			if len(goroutine) != 0 {

				Goroutines = append(Goroutines, goroutine)
				goroutine = Goroutine{}
			}
		}

		goroutine = append(goroutine, v)

	}
	if goroutine != nil {
		Goroutines = append(Goroutines, goroutine)
	}

	return Goroutines, nil
}
