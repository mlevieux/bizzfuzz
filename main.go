package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sync"
)

func main() {


	sh := statsCalls{}
	sh.info = make(map[string]int)
	sh.mux = new(sync.Mutex)

	r := mux.NewRouter()
	r.HandleFunc("/fizzbuzz", sh.handleFizzBuzz)
	r.HandleFunc("/statistics", sh.handleStatistics)

	log.Fatalln(http.ListenAndServe(":8080", r))
}