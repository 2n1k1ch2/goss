package prometheus

import (
	"goss/pkg/cluster"
)

type Exporter struct {
	Cluster cluster.Cluster
}

func NewExporter() *Exporter {
	return &Exporter{
		Cluster: map[string]cluster.Object{},
	}
}

func (e *Exporter) UpdateMetrics(cluster cluster.Cluster) {
	e.Cluster = cluster
}
