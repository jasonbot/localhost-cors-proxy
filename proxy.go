package main

import (
	"flag"

	corsproxy "./corsproxy"
)

func main() {
	var listenport int
	var forwardurl string

	flag.IntVar(&listenport, "listenport", 1234, "the port to bind to (bound to 0.0.0.0)")
	flag.StringVar(&forwardurl, "forwardurl", "http://localhost:5000", "the url to forward to")
	flag.Parse()

	proxy, err := corsproxy.NewProxy(listenport, forwardurl)
	if err == nil {
		proxy.Serve()
	}
}
