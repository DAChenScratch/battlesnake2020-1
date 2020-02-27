package logic

import (
	"math"

	"github.com/otonnesen/battlesnake2020/api"
)

type Direction string

const (
	Left  Direction = "left"
	Down  Direction = "down"
	Up    Direction = "up"
	Right Direction = "right"
)

func GetMove(data *api.MoveRequest) Direction {
	// state := InitState(data)
	return Up
}

func (s State) heuristic(snakeID string) float64 {
	// TODO: Just makes you not die (sometimes) for now
	you := s.Snakes[snakeID]
	head := you.Head()
	if s.outOfBounds(head) {
		return math.Inf(-1)
	}
	if s.inSnakeBody(head) {
		return math.Inf(-1)
	}
	return math.Inf(1)
}

func (s State) outOfBounds(p Point) bool {
	return p.X < 0 || p.X >= s.Width || p.Y < 0 || p.Y >= s.Height
}

func (s State) inSnakeBody(p Point) bool {
	for _, snake := range s.Snakes {
		if p.In(snake.Body) {
			return true
		}
	}
	return false
}
