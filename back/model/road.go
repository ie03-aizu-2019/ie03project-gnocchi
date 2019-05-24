package model

import "math"

func (r *Road) Length() float64 {
	return math.Sqrt(math.Pow(r.To.Coord.X-r.From.Coord.X, 2) + math.Pow(r.To.Coord.Y-r.From.Coord.Y, 2))
}

type Roads []*Road

func (rs Roads) Length() (result float64) {
	for _, r := range rs {
		result += r.Length()
	}
	return
}

func (rs Roads) Len() int { return len(rs) }

func (rs Roads) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

func (rs Roads) Less(i, j int) bool {
	return rs[i].Length() < rs[j].Length()
}

func (rs *Roads) Push(item interface{}) {
	*rs = append(*rs, item.(*Road))
}

func (rs *Roads) Pop() interface{} {
	old := *rs
	n := len(old)

	item := old[n-1]
	*rs = old[0 : n-1]

	return item
}
