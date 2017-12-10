package uds

import (
	"net"
	"github.com/pkg/errors"
	"net/http"
)

func ListenAndServe(sockPath string, handler http.Handler) error {
	if handler == nil {
		handler = http.DefaultServeMux
	}
	l, err := net.Listen("unix", sockPath)
	if err != nil {
		return errors.Wrap(err, "failed to listen")
	}
	return http.Serve(l, handler)
}