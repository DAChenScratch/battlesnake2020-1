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
	id := data.You.ID
	state := InitState(data)
	var dir Direction
	value := math.Inf(-1)
	moveslice := state.getChildMoves()
	for _, moves := range moveslice {
		state.applyAllMoves(moves)

		if m := state.minimax(4, id, true); m > value {
			value = m
			for _, move := range moves {
				if move.ID == id {
					dir = move.Forward
				}
			}
		}

		state.undoAllMoves(moves)
	}
	return dir
}

func (s *State) minimax(depth int, snakeID string, maximizing bool) float64 {
	if depth == 0 /* or if game is over */ {
		return s.heuristic(snakeID)
	}
	children := s.getChildMoves()
	if maximizing {
		value := math.Inf(-1)
		for _, child := range children {
			s.applyAllMoves(child)

			value = math.Max(value, s.minimax(depth-1, snakeID, false))

			s.undoAllMoves(child)
		}
		return value
	} else {
		value := math.Inf(1)
		for _, child := range children {
			s.applyAllMoves(child)

			value = math.Min(value, s.minimax(depth-1, snakeID, true))

			s.undoAllMoves(child)
		}
		return value
	}
}

func (s State) heuristic(snakeID string) float64 {
	// TODO: Just makes you not die (sometimes) for now
	you := s.Snakes[snakeID]
	head := you.Head()
	if s.outOfBounds(head) {
		return math.Inf(-1)
	}
	for _, snake := range s.Snakes {
		if snake.ID == you.ID {
			continue
		}
		if head.In(snake.Body) {
			// TODO: Check all snakes as below (without the head) and
			// check heads separately (in the cas where the other snake
			// is smaller, we can consider its head safe).

			return math.Inf(-1)
		}
		for p := you.Body.Front().Next(); p != nil; p = p.Next() {
			if head.X == p.Value.(Point).X && head.Y == p.Value.(Point).Y {
				return math.Inf(-1)
			}
		}
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
