package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/yowcow/ezserve/handler"
)

type Options struct {
	addr        string
	allowOrigin string
	root        string
	quiet       bool
}

var opt Options

func init() {
	flag.StringVar(&opt.addr, "addr", ":10080", "address to bind")
	flag.StringVar(&opt.allowOrigin, "allow-origin", "", "access-control-allow-origin")
	flag.StringVar(&opt.root, "root", ".", "root directory")
	flag.BoolVar(&opt.quiet, "quiet", false, "quiet output")
	flag.Parse()
}

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	logger.Println("serving static files under", opt.root, "at address", opt.addr)

	handlers := []http.Handler{
		handler.NewCORSHandler(opt.allowOrigin),
		http.FileServer(http.Dir(opt.root)),
	}
	if !opt.quiet {
		handlers = append(handlers, handler.NewLoggingHandler(logger))
	}

	log.Fatal(http.ListenAndServe(opt.addr, handler.NewMiddleware(handlers)))
}
