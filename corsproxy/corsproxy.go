package corsproxy

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type corsProxyStruct struct {
	reverseproxy *httputil.ReverseProxy
	listenport   int
	forwardurl   string
}

// CorsProxy is the entry point to the HTTP proxy
type CorsProxy interface {
	Serve()
}

func (p *corsProxyStruct) Serve() {
	log.Printf("Listening on port %v; forwarding port %v\n", p.listenport, p.forwardurl)

	corsCombiner := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowedOrigin := r.Header.Get("Origin")
		if allowedOrigin == "" {
			allowedOrigin = fmt.Sprintf("%v://%v", r.URL.Scheme, r.Host)
		}
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Headers", "authorization, origin, x-requested-with")
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PATCH, DELETE")

		// Fetch API sends an OPTIONS call that may not be supported
		if r.Method == "OPTIONS" {
			log.Printf("Intercepting OPTIONS on %v", r.URL)
			w.Header().Set("Allow", "OPTIONS, GET, POST, PATCH, DELETE")
			w.WriteHeader(http.StatusNoContent)
		} else {
			log.Printf("Request %v %v", r.Method, r.URL)
			p.reverseproxy.ServeHTTP(w, r)
		}
	})

	err := http.ListenAndServe(fmt.Sprintf("127.0.0.1:%v", p.listenport), corsCombiner)
	if err != nil {
		log.Fatalf("Error serving: %v", err)
	}
}

// NewProxy returns a new HTTP proxy
func NewProxy(listenport int, forwardurl string) (CorsProxy, error) {
	proxyURL, err := url.Parse(forwardurl)

	if err != nil {
		log.Fatalf("Can't parse url %v: %v", forwardurl, err)
		return nil, err
	}

	proxy := httputil.NewSingleHostReverseProxy(proxyURL)

	return &corsProxyStruct{
		reverseproxy: proxy,
		listenport:   listenport,
		forwardurl:   forwardurl,
	}, nil
}
