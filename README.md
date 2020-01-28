# localhost-cors-proxy
A proxy that binds to `0.0.0.0:N` and reverse proxies to `127.0.0.1:M`, adding a `Access-Control-Allow-Origin: *` header so you can use the fetch API to talk to an API service that would not normally allow for CORS.

## Security notice

*DO NOT RUN THIS ON ANY UNTRUSTED NETWORK, YOU FOOLS*.

## Using

* `go build proxy.go`
* `./proxy -listenport 8000 -forwardport 8888`
