package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/otonnesen/battlesnake2020/api"
	"github.com/otonnesen/battlesnake2020/logic"
)

func start(w http.ResponseWriter, req *http.Request) {
	data, err := api.NewMoveRequest(req)
	if err != nil {
		log.Printf("Bad start request: %v\n", err)
	}
	if logging {
		id := data.Game.ID
		logfile, err := os.Create("./log/" + id + ".log")
		if err != nil {
			panic(err)
		}

		logger := log.New(logfile, "", 0)

		lf := LogFile{logger, logfile}

		movelog[id] = lf
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
		jsonstr, _ := json.Marshal(data)
		movelog[data.Game.ID].logger.Printf("%s\n", jsonstr)
	}

	resp := &api.MoveResponse{Move: string(move)} // Get Move

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func end(w http.ResponseWriter, req *http.Request) {
	// data, err := api.NewMoveRequest(req)
	_, err := api.NewMoveRequest(req)
	// log.Printf("%+v\n", data)
	if err != nil {
		log.Printf("Bad move request: %v\n", err)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(api.MoveResponse{})
		return
	}
	if logging {
		// For some reason the battlesnake server is sending the
		// end request before any move requests, so the snake
		// crashes becase the log file is closed.

		// id := data.Game.ID
		// movelog[id].file.Close()
		// delete(movelog, id)
	}
}

func ping(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Pong!")
	return
}
