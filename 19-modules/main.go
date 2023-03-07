package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to mod in go")

	greeter()

	// mux router
	r := mux.NewRouter()
	r.HandleFunc("/", HomeServe).Methods("GET")

	// listening on Port 4000
	log.Fatal(http.ListenAndServe(":4000", r))

}

func greeter() {
	fmt.Println("Hey, there I am mod in go")
}

func HomeServe(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<html><head></head><body style='background: black;'><div style='display: flex;justify-content: center;'><h1 style='color: white;'> Welcome to golang series </h1></div></gdiv></body></html>"))
}
