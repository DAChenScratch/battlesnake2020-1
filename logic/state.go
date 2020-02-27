package logic

import (
	"container/list"

	"github.com/otonnesen/battlesnake2020/api"
)

type State struct {
	YouID  string
	Height int
	Width  int
	Food   []Point
	Snakes map[string]*Snake
}

func InitState(data *api.MoveRequest) *State {
	s := &State{}
	s.YouID = data.You.ID
	s.Height = data.Board.Height
	s.Width = data.Board.Width
	s.Food = make([]Point, len(data.Board.Food))
	for i, f := range data.Board.Food {
		s.Food[i] = Point{X: f.X, Y: f.Y}
	}
	s.Snakes = make(map[string]*Snake)
	for _, snake := range data.Board.Snakes {
		b := list.New()
		for _, p := range snake.Body {
			b.PushBack(Point{X: p.X, Y: p.Y})
		}
		s.Snakes[snake.ID] = &Snake{
			Body:   b,
			Health: snake.Health,
			ID:     snake.ID,
		}
	}
	return s
}

// Move stores a single move for a single snake.
// It stores the tail so as to be able to be undone,
// allowing a single State struct to be used for any
// number of moves, eliminating the need for duplication.
// With any luck, this will provide a meaningful speedup.
type Move struct {
	ID      string    // ID of snake to move
	Forward Direction // Direction to move
	Tail    Point     // Current position of tail to undo move
}

// m is assumed to be within the bounds of the board.
func (s *State) applyMove(m Move) {
	s.Snakes[m.ID].applyMove(m)
}

func (s *State) undoMove(m Move) {
	s.Snakes[m.ID].undoMove(m)
}

func (s *State) applyAllMoves(moves []Move) {
	for _, move := range moves {
		s.applyMove(move)
	}
}

func (s *State) undoAllMoves(moves []Move) {
	for _, move := range moves {
		s.undoMove(move)
	}
}

func (s State) getSnake(id string) *Snake {
	for k, snake := range s.Snakes {
		if k == id {
			return snake
		}
	}
	return nil
}

// Returns a map of the neighbors of each snake's head.
func (s State) allMoves() map[string][]Move {
	movemap := make(map[string][]Move)
	for _, snake := range s.Snakes {
		movemap[snake.ID] = []Move{
			Move{ID: snake.ID, Forward: Left, Tail: snake.Tail()},
			Move{ID: snake.ID, Forward: Up, Tail: snake.Tail()},
			Move{ID: snake.ID, Forward: Down, Tail: snake.Tail()},
			Move{ID: snake.ID, Forward: Right, Tail: snake.Tail()},
		}
	}
	return movemap
}

// Removes all moves that would kill the snake.
// TODO: Maybe replace with actual alhpa-beta pruning at some point.
func (s State) filterDeath(movemap map[string][]Move) map[string][]Move {
	for id, moves := range movemap {
		i := 0
		for {
			if i >= len(movemap[id]) {
				break
			}
			p := s.getSnake(id).Head().To(moves[i].Forward)
			if s.outOfBounds(p) || s.inSnakeBody(p) {
				movemap[id] = append(movemap[id][:i], movemap[id][i+1:]...)
			} else {
				i++
			}
		}
	}
	return movemap
}

func (s State) getChildMoves() [][]Move {
	moves := s.filterDeath(s.allMoves())
	var moveslice [][]Move
	for _, m := range moves {
		// Map -> slice to take cartesian product
		moveslice = append(moveslice, m)
	}
	return cartesianProduct(moveslice)
}
