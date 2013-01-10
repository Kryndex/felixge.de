package main

import (
	"fmt"
	"net"
	"net/http"
	"path"
	"runtime"
	"github.com/felixge/felixge.de/fs"
)

var (
	_, filename, _, _ = runtime.Caller(0)
	root       = path.Join(path.Dir(filename), "../..")
)

func main() {
	fs := fs.New(http.Dir(root+"/pages"))

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Listening at: http://%s\n", listener.Addr())

	if err := http.Serve(listener, http.FileServer(fs)); err != nil {
		panic(err)
	}
}
