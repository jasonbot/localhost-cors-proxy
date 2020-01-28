package main

import (
	"flag"

	corsproxy "./corsproxy"
)

func main() {
	var listenport, forwardport int

	flag.IntVar(&listenport, "listenport", 1234, "the port to bind to (bound to 0.0.0.0)")
	flag.IntVar(&forwardport, "forwardport", 5000, "the port to forward to (bound to 127.0.0.1)")
	flag.Parse()

	proxy := corsproxy.NewProxy(listenport, forwardport)
	proxy.Serve()
}
