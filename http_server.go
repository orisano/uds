package uds

import (
	"fmt"
	"net"
	"net/http"
)

func ListenAndServe(sockPath string, handler http.Handler) error {
	if handler == nil {
		handler = http.DefaultServeMux
	}
	l, err := net.Listen("unix", sockPath)
	if err != nil {
		return fmt.Errorf("listen unix: %w", err)
	}
	return http.Serve(l, handler)
}
