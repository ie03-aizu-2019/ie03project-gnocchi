package model

type Point struct {
	X float64
	Y float64
}

type Place struct {
	Id    int
	Coord Point
}

type Road struct {
	Id   int
	From Place
	To   Place
}

type datas struct {
	Places []Place
	Roads  []Road
}
