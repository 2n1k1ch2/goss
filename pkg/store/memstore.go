package store

import (
	"goss/pkg/cluster"
	"time"
)

type Store struct {
	SnapShots []SnapShot
	Size      int
}
type SnapShot struct {
	cl        cluster.Cluster
	timestamp time.Time
}

func (s *Store) Add(cl cluster.Cluster) {
	s.ComputeDrift(&cl)
	if len(s.SnapShots) < s.Size {
		s.SnapShots = append(s.SnapShots, SnapShot{cl, time.Time{}})
		return
	}
	s.SnapShots = append(s.SnapShots[1:], SnapShot{cl, time.Time{}})

}
func (s *Store) ComputeDrift(cl *cluster.Cluster) {
	last := s.SnapShots[len(s.SnapShots)-1]
	currentTime := time.Now()

	for hash := range *cl {
		if val, exists := last.cl[hash]; exists {
			timeDiff := currentTime.Sub(last.timestamp).Seconds()

			if timeDiff > 0 {
				countDiff := float64((*cl)[hash].Count - val.Count)
				drift := countDiff / timeDiff

				obj := (*cl)[hash]
				obj.Drift_rate = drift
				(*cl)[hash] = obj
			}
		}
	}
}
