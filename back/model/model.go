package model

type Point struct {
	X float32
	Y float32
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
