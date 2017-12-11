package main

import (
	"io"
	"log"
	"os"

	"github.com/orisano/uds"
)

const (
	sockPath = "../server/sample.sock"
)

func main() {
	client := uds.NewClient(sockPath)
	resp, err := client.Get("http://unix/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
