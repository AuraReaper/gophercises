package main

import (
	"log"
	"net/http"

	"github.com/AuraReaper/choose_your_own_adventure/internal"
)

const filePath = "data/gopher.json"

func main() {
	svc := internal.NewService(filePath)
	handler := internal.NewHandler(svc)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.RootHandler)
	mux.HandleFunc("/{arc}", handler.LoopHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
