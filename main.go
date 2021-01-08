package main

import (
	"flag"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sync"
)

var (
	flagListenAddr = flag.String("a", ":8080", "specifies what address to listen to to the server")
)

func main() {

	flag.Parse()

	sh := statsCalls{}
	sh.info = make(map[string]int)
	sh.mux = new(sync.Mutex)

	r := mux.NewRouter()
	r.HandleFunc("/fizzbuzz", sh.handleFizzBuzz)
	r.HandleFunc("/statistics", sh.handleStatistics)

	log.Fatalln(http.ListenAndServe(*flagListenAddr, r))
}