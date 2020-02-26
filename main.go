package main

import (
	"log"
	"net/http"
	"os"

	"github.com/otonnesen/battlesnake2020/api"
	"github.com/otonnesen/battlesnake2020/util"
)

var startResp = api.StartResponse{Color: "#75CEDD", SecondaryColor: "#7A75DD"}
var movelog *log.Logger
var logging bool // Turns on logging JSON move data for replaying games

func main() {
	port := os.Getenv("PORT") // Get Heroku port (or non-Heroku port, I guess)

	if port == "" {
		log.Printf("$PORT not set, defaulting to 8080")
		port = "8080"
	}

	logging = false
	args := os.Args[1:]
	if len(args) > 0 {
		// Who needs argparse anyway
		if args[0] == "log" {
			log.Printf("Logging enabled")
			logging = true
		}
	}

	logfile, err := os.Open("/dev/null")
	if err != nil {
		panic(err)
	}
	movelog = log.New(logfile, "", 0)

	http.HandleFunc("/start", util.LogRequest(start))
	http.HandleFunc("/move", util.LogRequest(move))
	http.HandleFunc("/end", util.LogRequest(end))
	http.HandleFunc("/ping", util.LogRequest(ping))

	log.Printf("Server running on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}
