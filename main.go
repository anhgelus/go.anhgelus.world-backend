package main

import (
	"github.com/anhgelus/go.anhgelus.world-backend/src"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{id}/{return:.+}", src.HandleBasic)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":80",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Default().Println("Starting...")
	log.Fatal(srv.ListenAndServe())
}
