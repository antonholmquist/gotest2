package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
	"os"
)

func main() {

	db, err := sql.Open("postgres", "host=localhost port=5432 dbname=test sslmode=disable")
	err = db.Ping()

	if err != nil {
		fmt.Println("error opening database:", err)
	} else if db != nil {
		fmt.Println("opened database:", db)
	}

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

	router.Methods("GET").Path("/user").HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/plain")
		res.Write([]byte("GET USER"))
	})

	router.Methods("POST").Path("/user").HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		_, queryError := db.Exec("INSERT INTO \"user\" (name) VALUES ($1);", "anton holmquist")

		if queryError != nil {
			res.Header().Set("Content-Type", "text/plain")
			res.Write([]byte("POST FAILED: " + queryError.Error()))
		} else {

			res.Header().Set("Content-Type", "text/plain")
			res.Write([]byte("POST USER"))
		}

	})

	http.Handle("/", router)
	http.ListenAndServe(":"+port, nil)
	fmt.Println("Starting server on port", port)
}
