package main

import (
    "fmt"
    "net/http"
    "os"
    "github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {

	port := os.Getenv("PORT");

	if len(port) == 0 {
		port = "3000"
	}	


	router := mux.NewRouter()
	router.HandleFunc("/", func (res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/plain")
		res.Write([]byte("This is an example server.\n"))
	})

    http.Handle("/", router)
    http.ListenAndServe(":" + port, nil)
}