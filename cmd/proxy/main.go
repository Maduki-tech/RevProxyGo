package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
)

var backenURLS = []string{
	"http://localhost:8081",
	"http://localhost:8082",
}

var currentBackend uint64

func getNextBackend() string {
	next := atomic.AddUint64(&currentBackend, 1)
	return backenURLS[next%uint64(len(backenURLS))]
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		backendURL, _ := url.Parse(getNextBackend())

		proxy := httputil.NewSingleHostReverseProxy(backendURL)

		proxy.ServeHTTP(w, r)
	})

	log.Println("Starting reverse proxy on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))

}
