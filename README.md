# localhost-cors-proxy
A proxy that binds to `127.0.0.1:N` and reverse proxies to `forwardurl`, adding a set of various CORS headers to enable CORS (and therefore the javascript fetch api in browser). I developed this to help with testing local react apps that use fetch to talk to APIs elsewhere.

## Security notice

**DO NOT RUN THIS ON ANY UNTRUSTED NETWORK, YOU FOOLS**.

## Using

* `go build proxy.go`
* `./proxy -listenport 8000 -forwardurl http://localhost:8888`
