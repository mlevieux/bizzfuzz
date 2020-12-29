package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	sh := newStatistics()

	r := mux.NewRouter()
	r.HandleFunc("/fizzbuzz", sh.handleFizzBuzz)
	r.HandleFunc("/statistics", sh.handleStatistics)

	log.Fatalln(http.ListenAndServe(":8080", r))
}