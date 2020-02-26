package main

import (
	"log"
	"net/http"
	"os"

	"github.com/otonnesen/battlesnake2020/api"
	"github.com/otonnesen/battlesnake2020/util"
)

var startResp = api.StartResponse{Color: "#75CEDD", SecondaryColor: "#7A75DD"}
var movelog map[string]*log.Logger
var movelog_file map[string]*os.File
var logging bool // Turns on logging JSON move data for replaying games

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		// Who needs argparse anyway
		if args[0] == "log" {
			movelog = make(map[string]*log.Logger)
			movelog_file = make(map[string]*os.File)
			log.Printf("Logging enabled")
			logging = true
		}
	}

	port := os.Getenv("PORT") // Get Heroku port (or non-Heroku port, I guess)

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
