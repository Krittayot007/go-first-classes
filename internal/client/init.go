package client

import (
	"context"
	"net/http"
	"time"
)

type Client interface{}

type clientImpl struct {
	webClient *http.Client
}

func New(ctx context.Context) (Client, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	return &clientImpl{
		webClient: client}, nil
}
