package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/yowcow/ezserve/cors"
	"github.com/yowcow/ezserve/logging"
)

var addr string
var root string
var cert string
var key string
var allowCORS bool
var quiet bool

func init() {
	flag.StringVar(&addr, "addr", ":10080", "address to bind")
	flag.StringVar(&root, "root", ".", "root directory")
	flag.StringVar(&cert, "cert", "", "certificate file")
	flag.StringVar(&key, "key", "", "key file")
	flag.BoolVar(&allowCORS, "allow-cors", false, "allow CORS policy")
	flag.BoolVar(&quiet, "quiet", false, "quiet output")
	flag.Parse()
}

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	logger.Println("serving static files under", root, "at address", addr)

	handler := http.FileServer(http.Dir(root))

	if allowCORS {
		handler = cors.NewHandler(handler, allowCORS)
	}

	if !quiet {
		handler = logging.NewHandler(handler, logger)
	}

	if cert != "" && key != "" {
		log.Fatalln(http.ListenAndServeTLS(addr, cert, key, handler))
	} else {
		log.Fatalln(http.ListenAndServe(addr, handler))
	}
}
