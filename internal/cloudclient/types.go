package cloudclient

import (
	"goss/pkg/cluster"
	"time"
)

type Snapshot struct {
	ReleaseTag string
	Timestamp  time.Time
	AgentID    string
	Cluster    cluster.Cluster
	Success    bool
	StatusCode int
	Error      string
}
