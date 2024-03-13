package main

import (
	"HTTP/api"
	"net/http"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}

// http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Hello World"))
// })
// http.ListenAndServe(":8080", nil)
