package cloudclient

import (
	"bytes"
	"context"
	"goss/pkg/cluster"
	"goss/pkg/config"
	"log"
	"net/http"
)

type CloudClient struct {
	baseURL   string
	authToken string
	client    *http.Client
	cfg       *config.Config
}

func (c *CloudClient) SendSnapshot(ctx context.Context, cluster *cluster.Cluster) error {
	snap := BuildPayload(*cluster, c.cfg.ReleaseTag, c.cfg.AgentID)

	data, err := SerializeToNDJSON(&snap)
	if err != nil {
		return err
	}

	compressed, err := CompressGzip(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/snapshot", bytes.NewBuffer(compressed))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Encoding", "gzip")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.Println("cloud ingest status:", resp.StatusCode)
	return nil
}
