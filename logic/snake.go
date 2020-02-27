package logic

import "container/list"

// We use a linked list for the body for quick
// insertion/deletion of head/tail.
type Snake struct {
	Body   *list.List // Doubly linked list of Points.
	Health int
	ID     string
}

func (s Snake) Head() Point {
	return s.Body.Front().Value.(Point)
}

func (s Snake) Tail() Point {
	return s.Body.Back().Value.(Point)
}

// m is assumed to be within the bounds of the board.
func (s *Snake) applyMove(m Move) {
	// TODO:
	// * Check if a snake eats food and if so,
	// leave its tail next turn
	// 		**** Actually just PushBack the point its tail
	// 		**** is at so there are two points stacked on the
	// 		**** tail.
	// * Decrement snake health or reset it to 100
	// if it eats food
	s.Body.PushFront(s.Head().To(m.Forward))
	s.Body.Remove(s.Body.Back())
}

func (s *Snake) undoMove(m Move) {
	s.Body.PushBack(m.Tail)
	s.Body.Remove(s.Body.Front())
}
