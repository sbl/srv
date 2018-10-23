package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main() {
	var (
		addr = flag.String("addr", ":8080", "srv addr")
		dir  = flag.String("dir", ".", "directory to srv")
	)
	flag.Parse()

	fileServer := http.FileServer(http.Dir(*dir))
	http.Handle("/", handlers.LoggingHandler(os.Stdout, fileServer))

	log.Println("=== srv", *dir, "->", *addr, "===")
	log.Fatal(http.ListenAndServe(*addr, nil))
}
