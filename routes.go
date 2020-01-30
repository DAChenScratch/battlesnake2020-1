package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/otonnesen/battlesnake2020/api"
)

func start(w http.ResponseWriter, req *http.Request) {
	_, err := api.NewStartRequest(req) // TODO: Do something with data?
	if err != nil {
		log.Printf("Bad start request: %v\n", err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(startResp)
}

func move(w http.ResponseWriter, req *http.Request) {
	data, err := api.NewMoveRequest(req)
	log.Printf("%+v\n", data)
	if err != nil {
		log.Printf("Bad move request: %v\n", err)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(api.MoveResponse{})
		return
	}

	resp := &api.MoveResponse{Move: "up"} // Get Move

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
