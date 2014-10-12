package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "3000"
	}

	// gorilla/mux is a powerful URL router and dispatcher.
	router := mux.NewRouter().PathPrefix("/v1").Subrouter()

	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/plain")
		res.Write([]byte("This is an example server."))
	})

	router.HandleFunc("/users/{userID}", func(res http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)
		userID := vars["userID"]

		res.Header().Set("Content-Type", "text/plain")
		res.Write([]byte("user id: " + userID))
	})

	http.Handle("/", router)
	http.ListenAndServe(":"+port, nil)
	fmt.Println("Starting server on port", port)
}
