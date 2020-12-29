package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/fizzbuzz", handleFizzBuzz)

	log.Fatalln(http.ListenAndServe(":8080", r))
}