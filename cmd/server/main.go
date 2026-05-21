package main

import (
	"log"
	"net/http"
	"os"

	httpapi "backgammon-teacher/internal/api/http"
)

func main() {
	addr := os.Getenv("HTTP_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	server := httpapi.NewServer()
	log.Printf("backgammon-teacher server listening on %s", addr)
	if err := http.ListenAndServe(addr, server); err != nil {
		log.Fatal(err)
	}
}
