package main

import (
	"log"
	"net/http"
	"sybstring-finder/internal/server"
)

func main() {
	http.HandleFunc("/api/substring", server.GetSubstring)

	log.Print("start server")
	panic(http.ListenAndServe(":8080", nil))
}
