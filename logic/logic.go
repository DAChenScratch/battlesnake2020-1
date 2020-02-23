package logic

import "github.com/otonnesen/battlesnake2020/api"

type Direction string

const (
	Left  Direction = "left"
	Down  Direction = "down"
	Up    Direction = "up"
	Right Direction = "right"
)

func GetMove(data *api.MoveRequest) Direction {
	state := InitState(data)
	return Up
}
