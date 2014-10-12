package main

import (
    "fmt"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {

	port := os.Getenv("PORT");

	if len(port) == 0 {
		port = "3000"
	}	

    http.HandleFunc("/", handler)
    http.ListenAndServe(":" + port, nil)
}