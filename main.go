package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/login", LoginHandler).Methods("POST")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8082", nil))

	// http.HandleFunc("/login", Login)
	// http.HandleFunc("/home", Home)
	// http.HandleFunc("/refresh", Refresh)

	// log.Fatal(http.ListenAndServe(":8082", nil))

}
