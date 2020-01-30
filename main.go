package main

import (
	"log"
	"net/http"
	"os"

	"github.com/otonnesen/battlesnake2020/api"
	"github.com/otonnesen/battlesnake2020/util"
)

var startResp = api.StartResponse{Color: "#75CEDD", SecondaryColor: "#7A75DD"}

func main() {
	port := os.Getenv("PORT") // Get Heroku port

	if port == "" {
		log.Printf("$PORT not set, defaulting to 8080")
		port = "8080"
	}

	http.HandleFunc("/start", util.LogRequest(start))
	http.HandleFunc("/move", util.LogRequest(move))
	http.HandleFunc("/end", util.LogRequest(end))
	http.HandleFunc("/ping", util.LogRequest(ping))

	log.Printf("Server running on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}
