package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	port = flag.String("port", "5000", "srv port")
	dir  = flag.String("dir", ".", "directory to srv")
)

func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Println(req.Method, req.URL.String())
		h.ServeHTTP(w, req)
	})
}

func main() {
	flag.Parse()

	fileServer := http.FileServer(http.Dir(*dir))
	host := "localhost:" + *port

	fmt.Println("=== srv", *dir, "->", host, "===")
	if err := http.ListenAndServe(host, logger(fileServer)); err != nil {
		panic("could not srv" + host)
	}
}
