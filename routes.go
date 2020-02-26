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
		movelog_file[id] = logfile
		movelog[id] = log.New(logfile, "", 0)
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
		movelog[data.Game.ID].Printf("%s\n", jsonstr)
	}

	resp := &api.MoveResponse{Move: string(move)} // Get Move

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func end(w http.ResponseWriter, req *http.Request) {
	data, err := api.NewMoveRequest(req)
	// log.Printf("%+v\n", data)
	if err != nil {
		log.Printf("Bad move request: %v\n", err)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(api.MoveResponse{})
		return
	}
	if logging {
		id := data.Game.ID
		movelog_file[id].Close()
		delete(movelog_file, id)
		delete(movelog, id)
	}
}

func ping(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Pong!")
	return
}
