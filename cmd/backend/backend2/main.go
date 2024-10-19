package main

import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello From Server 2"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8082", nil)
}
