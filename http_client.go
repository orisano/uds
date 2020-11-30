package uds

import (
	"context"
	"net"
	"net/http"
)

func NewClient(sockPath string) *http.Client {
	t := &http.Transport{
		DialContext: func(ctx context.Context, _, _ string) (net.Conn, error) {
			var d net.Dialer
			return d.DialContext(ctx, "unix", sockPath)
		},
	}
	return &http.Client{
		Transport: t,
	}
}
