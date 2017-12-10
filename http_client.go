package uds

import (
	"context"
	"net"
	"net/http"
)

func NewClient(sockPath string) *http.Client {
	t := &http.Transport{
		DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
			return net.Dial("unix", sockPath)
		},
	}
	return &http.Client{
		Transport: t,
	}
}
