package logic

import "container/list"

type Point struct {
	X int
	Y int
}

func (p Point) To(d Direction) *Point {
	switch d {
	case Left:
		return &Point{X: p.X - 1, Y: p.Y}
	case Down:
		return &Point{X: p.X, Y: p.Y + 1}
	case Up:
		return &Point{X: p.X, Y: p.Y - 1}
	case Right:
		return &Point{X: p.X + 1, Y: p.Y}
	default:
		return &Point{-1, -1}
	}
}

func (p Point) In(points *list.List) bool {
	for o := points.Front(); o != nil; o = o.Next() {
		if p.X == o.Value.(*Point).X && p.Y == o.Value.(*Point).Y {
			return true
		}
	}
	return false
}

func (p Point) Neighbors() []*Point {
	return []*Point{
		p.To(Left),
		p.To(Down),
		p.To(Up),
		p.To(Right),
	}
}
