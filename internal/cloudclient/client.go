package cloudclient

import (
	"context"
	"goss/pkg/cluster"
	"net/http"
)

type CloudClient struct {
	baseURL   string
	authToken string
	client    *http.Client
}

func (c CloudClient) SendSnapshot(ctx context.Context, cluster *cluster.Cluster) error {
	snap := BuildPayload(*cluster, ctx.Value(), ctx)
	data, err := SerializeToNDJSON(&snap)
	if err != nil {
		return err
	}
	compressdata, err := CompressGzip(data)
	if err != nil {
		return err
	}

	return nil
}
