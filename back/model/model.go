package model

import (
	"fmt"
	"math"
)

type Stringable interface {
	ToString() string
}

type Point struct {
	X float64
	Y float64
}

type Place struct {
	Id    string
	Coord Point
}

type Road struct {
	Id   int
	From *Place
	To   *Place
}

type Query struct {
	Start string
	Dest  string
	Num   int
}

func (p *Point) Length() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func (p Point) ToString() string {
	return fmt.Sprintf("%.5f %.5f", p.X, p.Y)
}

func (p Place) ToString() string {
	return p.Coord.ToString()
}
func (r Road) ToString() string {
	return fmt.Sprintf("%d %s %s", r.Id, r.From.Id, r.To.Id)
}

func (q Query) ToString() string {
	return fmt.Sprintf("%s %s %d", q.Start, q.Dest, q.Num)
}
