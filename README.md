# uds
unix domain socket utilities for golang.

## How to Use
examples/server/main.go

```go
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

```

examples/client/main.go

```go
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
```

## LICENSE
MIT