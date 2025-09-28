package cloudclient

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"goss/pkg/cluster"
	"net/http"
	"time"
)

func SerializeToNDJSON(snapshot *Snapshot) ([]byte, error) {
	snapshotJSON, err := json.Marshal(snapshot)
	if err != nil {
		return nil, err
	}
	return snapshotJSON, nil

}
func CompressGzip(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	gzipWriter := gzip.NewWriter(&buf)

	_, err := gzipWriter.Write(data)
	if err != nil {
		return nil, err
	}

	err = gzipWriter.Close()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
func BuildPayload(cluster cluster.Cluster, releaseTag string, agentID string) Snapshot {
	snap := Snapshot{
		ReleaseTag: releaseTag,
		Timestamp:  time.Now(),
		Cluster:    cluster,
		AgentID:    agentID,
		Success:    false,
		StatusCode: http.StatusCreated,
		Error:      "",
	}
	return snap
}
