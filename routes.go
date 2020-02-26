package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/otonnesen/battlesnake2020/api"
	"github.com/otonnesen/battlesnake2020/logic"
)

func start(w http.ResponseWriter, req *http.Request) {
	data, err := api.NewStartRequest(req)
	if err != nil {
		log.Printf("Bad start request: %v\n", err)
	}
	if logging {
		fname := strconv.Itoa(data.GameID) + ".log"
		logfile, err := os.Create(fname)
		if err != nil {
			panic(err)
		}
		movelog = log.New(logfile, "", 0)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(startResp)
}

func move(w http.ResponseWriter, req *http.Request) {
	data, err := api.NewMoveRequest(req)
	// log.Printf("%+v\n", data)
	if err != nil {
		log.Printf("Bad move request: %v\n", err)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(api.MoveResponse{})
		return
	}

	move := logic.GetMove(data)
	log.Printf("MOVE: %s\n", move)
	if logging {
		jsonstr, _ := json.MarshalIndent(data, "", "  ")
		movelog.Printf("%s\n", jsonstr)
	}

	resp := &api.MoveResponse{Move: string(move)} // Get Move

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func end(w http.ResponseWriter, req *http.Request) {
	return
}

func ping(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Pong!")
	return
}
