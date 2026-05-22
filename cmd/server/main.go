package main

import (
	"log"
	"net/http"
	httpapi "backgammon-teacher/internal/api/http"
	coachapp "backgammon-teacher/internal/modules/coach/app"
)

func main() {
	server := httpapi.NewServer(coachapp.DefaultGenerateHint)
	log.Fatal(http.ListenAndServe(":8080", server))
}
