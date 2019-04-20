package model

import "math"

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
	From Place
	To   Place
}

type Query struct {
	Start string
	Dest  string
	Num   int
}

type datas struct {
	Places  []Place
	Roads   []Road
	Queries []Query
}

func (p *Point) Length() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}
