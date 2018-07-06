package main

import (
	"log"
	"net/http"
	"time"

	"github.com/jianhan/ms-sui-http/handler"
)

func main() {
	srv := &http.Server{
		Handler:      handler.Router(),
		Addr:         "127.0.0.1:9000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
