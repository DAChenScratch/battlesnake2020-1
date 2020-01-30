package logic

import "github.com/otonnesen/battlesnake2020/api"

type Direction string

const (
	Left  Direction = "left"
	Down  Direction = "down"
	Up    Direction = "up"
	Right Direction = "right"
)

func GetMoves(data *api.MoveRequest) Direction {
	return Up
}
