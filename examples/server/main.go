package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/orisano/uds"
)

const (
	sockPath = "./sample.sock"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html><body><h1>It works!</h1></body></html>")
	})
	os.Remove(sockPath)
	log.Fatal(uds.ListenAndServe("./sample.sock", nil))
}
