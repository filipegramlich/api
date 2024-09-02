package main

import (
	"log"
	"net/http"
)

func main() {
	// handler := http.HandlerFunc(server.PlayerServer)
	server := &PlayerServer{}
	log.Fatal(http.ListenAndServe(":5000", server))
}
