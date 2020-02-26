package api

import (
	"encoding/json"
	"net/http"
)

// StartRequest is specified by the
// 2019 battlesnake API
// NOTE: Deprecated for 2020, now uses same schema
// as MoveRequest. Changes reflected in routes.go.
type StartRequest struct {
	Game   game
	Width  int
	Height int
}

// StartResponse is specified by the
// 2019 battlesnake API
type StartResponse struct {
	Color          string `json:"color"`
	SecondaryColor string `json:"secondary_color"`
	HeadURL        string `json:"head_url"`
	Taunt          string `json:"taunt"`
	HeadType       string `json:"head_type"`
	TailType       string `json:"tail_type"`
}

// MoveRequest is specified by the
// 2019 battlesnake API
type MoveRequest struct {
	Board board `json:"board"`
	Game  game  `json:"game"`
	Turn  int   `json:"turn"`
	You   Snake `json:"you"`
}

// MoveResponse is specified by the
// 2019 battlesnake API
type MoveResponse struct {
	Move string `json:"move"`
}

type board struct {
	Food   []Point `json:"food"`
	Height int     `json:"height"`
	Width  int     `json:"width"`
	Snakes []Snake `json:"snakes"`
}

// Point represents a pair of coordinates on
// the battlesnake board.
type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Snake represents the JSON snake object
// sent in battlesnake 2019 move requests
type Snake struct {
	Body   []Point `json:"body"`
	Health int     `json:"health"`
	ID     string  `json:"id"`
	Name   string  `json:"name"`
}

type game struct {
	ID string `json:"id"`
}

// NewStartRequest unmarshals JSON from an http.Request into a StartRequest
func NewStartRequest(req *http.Request) (*StartRequest, error) {
	d := StartRequest{}
	err := json.NewDecoder(req.Body).Decode(&d)
	return &d, err
}

// NewMoveRequest unmarshals JSON from an http.Request into a MoveRequest
func NewMoveRequest(req *http.Request) (*MoveRequest, error) {
	d := MoveRequest{}
	err := json.NewDecoder(req.Body).Decode(&d)
	return &d, err
}
