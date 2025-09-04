package capture

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"strings"
)

const (
	RUNNING      string = "running"
	RUNNABLE     string = "runnable"
	SLEEP        string = "sleep"
	CHAN_SEND    string = "chan_send"
	CHAN_RECEIVE string = "chan_receive"
	SELECT       string = "select"
	IO_WAIT      string = "io_wait"
	SYSTEM_CALL  string = "system_call"
	GC_SWEEP     string = "gc_sweep"
	DEAD         string = "dead"
)

type Cluster = map[string]Object

type Object struct {
	Hash   string   `json:"hash"`
	Status string   `json:"status"`
	Name   string   `json:"name"`
	Frames []string `json:"-"`
	Count  uint64   `json:"count"`
	Ids    []uint64 `json:"ids"`
	Score  uint64   `json:"score"`
}

func Clusterize(gors []Goroutine) Cluster {
	cluster := make(Cluster)
	for _, g := range gors {
		status, err := findStatus(&g)
		if err != nil {
			log.Println(err)
			continue
		}
		name := giveName(&g)
		hsh, err := hashGoroutine(&g)
		if err != nil {
			log.Println(err)
			continue
		}

		if existing, exists := cluster[hsh]; exists {
			existing.Count += 1
			existing.Ids = append(existing.Ids, g.id)
			cluster[hsh] = existing
		} else {
			cluster[hsh] = Object{
				Hash:   hsh,
				Status: status,
				Name:   name,
				Frames: g.data[1:],
				Count:  1,
				Ids:    []uint64{g.id},
			}
		}
	}
	return cluster
}

func findStatus(g *Goroutine) (string, error) {
	fmt.Println(g.data)
	if strings.Contains(g.data[0], "running") {
		return RUNNING, nil
	}
	if strings.Contains(g.data[0], "runnable") {
		return RUNNABLE, nil
	}
	if strings.Contains(g.data[0], "sleep") {
		return SLEEP, nil
	}
	if strings.Contains(g.data[0], "chan send") {
		return CHAN_SEND, nil
	}
	if strings.Contains(g.data[0], "chan receive") {
		return CHAN_RECEIVE, nil
	}
	if strings.Contains(g.data[0], "select") {
		return SELECT, nil
	}
	if strings.Contains(g.data[0], "io wait") {
		return IO_WAIT, nil
	}
	if strings.Contains(g.data[0], "system call") {
		return SYSTEM_CALL, nil
	}
	if strings.Contains(g.data[0], "gc sweep") {
		return GC_SWEEP, nil
	}
	if strings.Contains(g.data[0], "dead") {
		return DEAD, nil
	}

	return "", errors.New("cant get status")

}

func giveName(g *Goroutine) string {
	str := strings.Split(g.data[0], ":")
	return str[1]
}

func hashGoroutine(g *Goroutine) (string, error) {
	if len(g.data) < 2 {
		return "", errors.New("cant get hash")
	}

	// make frames
	combined := strings.Join(g.data[1:], "")

	// then get hash from frames
	hasher := sha256.New()
	hasher.Write([]byte(combined))
	return hex.EncodeToString(hasher.Sum(nil)), nil
}
