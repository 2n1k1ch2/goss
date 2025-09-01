package prometheus

import (
	"goss/pkg/capture"
)

type Exporter struct {
	Cluster capture.Cluster
}

func NewExporter() *Exporter {
	return &Exporter{
		Cluster: map[string]capture.Object{},
	}
}

func (e *Exporter) UpdateMetrics(cluster capture.Cluster) {
	e.Cluster = cluster
}
