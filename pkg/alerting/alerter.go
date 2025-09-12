package alerting

import (
	"fmt"
	"goss/pkg/cluster"
	"time"
)

type Alerter struct {
	Threshold uint64
	Out       chan<- Alert
}

func NewAlerter(threshold uint64, out chan<- Alert) *Alerter {
	return &Alerter{Threshold: threshold, Out: out}
}

type Alert struct {
	Hash   string
	Name   string
	Score  uint64
	Status string
	Count  uint64
	IDs    []uint64
	Time   time.Time
}

func (a *Alerter) Check(cl *cluster.Cluster) {
	for _, v := range *cl {
		if v.Score >= a.Threshold {
			alert := Alert{
				Hash:   v.Hash,
				Name:   v.Name,
				Score:  v.Score,
				Status: v.Status,
				Count:  v.Count,
				IDs:    v.Ids,
				Time:   time.Now(),
			}

			select {
			case a.Out <- alert:
			default:
				fmt.Println("alerter out channel is full")

			}
		}
	}
}
