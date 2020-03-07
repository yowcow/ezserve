package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/yowcow/ezserve/logging"
)

var addr string
var root string

func init() {
	flag.StringVar(&addr, "addr", ":10080", "address to bind")
	flag.StringVar(&root, "root", ".", "root directory")
	flag.Parse()
}

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	logger.Println("serving static files under", root, "at address", addr)

	fs := http.FileServer(http.Dir(root))
	log.Fatalln(http.ListenAndServe(addr, logging.NewHandler(fs, logger)))
}
