package store

import (
	"goss/pkg/cluster"
	"time"
)

type Store struct {
	Items []item
	Size  int
}
type item struct {
	cl        cluster.Cluster
	timestamp time.Time
}

func (s *Store) Add(cl cluster.Cluster) {
	if len(s.Items) < s.Size {
		s.Items = append(s.Items, item{cl, time.Time{}})
		return
	}
	s.Items = append(s.Items[1:], item{cl, time.Time{}})
}
