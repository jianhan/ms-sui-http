package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"log"
)

func main() {
	r := mux.NewRouter()
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
