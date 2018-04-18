package uds

import "os"

func Get(sockPath, address string) (string, string) {
	if _, err := os.Stat(sockPath); err == nil {
		return "unix", sockPath
	} else {
		return "tcp", address
	}
}
