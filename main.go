package main

import (
	"github.com/anhgelus/go.anhgelus.world-backend/src"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	err := os.Mkdir("/config", 0777)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	src.LoadConfig("/config/redirections.toml")

	r := mux.NewRouter()
	r.HandleFunc("/{id}/{redirect:.+}", src.HandleBasic)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":80",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Default().Println("Starting...")
	log.Fatal(srv.ListenAndServe())
}
