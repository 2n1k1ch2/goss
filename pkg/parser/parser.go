package parser

import (
	"errors"
	"goss/pkg/pprof"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Goroutine struct {
	Data []string
	Id   uint64
}

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

			re := regexp.MustCompile(`^goroutine\s+(\d+)`)
			matches := re.FindStringSubmatch(v)
			if len(matches) > 1 {
				id, err := strconv.ParseUint(matches[1], 10, 64)
				if err == nil {
					goroutine.Id = id
				} else {
					log.Printf("failed to parse goroutine id: %v", err)
				}
			}
			v = re.ReplaceAllString(v, "")
			if len(goroutine.Data) != 0 {

				Goroutines = append(Goroutines, goroutine)
				goroutine = Goroutine{}
			}
		}

		goroutine.Data = append(goroutine.Data, v)

	}
	if goroutine.Data != nil {
		Goroutines = append(Goroutines, goroutine)
	}

	return Goroutines, nil
}
