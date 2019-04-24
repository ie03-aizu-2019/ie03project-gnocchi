package phase1

import (
	"fmt"
	"sort"

	"github.com/uzimaru0000/ie03project-gnocchi/back/model"
)

func EnumerateCrossPoints(roads []*model.Road) []*model.Road {
	crossPoints := []*model.Place{}

	i := 0
	for i < len(roads) {
		flag := false
		for j := i + 1; j < len(roads); j++ {
			road1 := roads[i]
			road2 := roads[j]
			p, err := CheckCrossPoint(road1, road2)
			if err == nil {
				crossPoint := &model.Place{Id: "X", Coord: *p}
				crossPoints = append(crossPoints, crossPoint)

				roads = append(roads, makeRoads(road1, crossPoint)...)
				roads = append(roads, makeRoads(road2, crossPoint)...)

				roads = remove(roads, i)
				roads = remove(roads, j-1)

				flag = true
				break
			}
		}

		if !flag {
			i++
		}
	}

	idReregistration(crossPoints)

	return roads
}

func remove(s []*model.Road, i int) []*model.Road {
	if i >= len(s) {
		return s
	}

	return append(s[:i], s[i+1:]...)
}

func makeRoads(r *model.Road, p *model.Place) []*model.Road {
	return []*model.Road{
		{
			Id:   0,
			From: r.From,
			To:   p,
		},
		{
			Id:   0,
			From: p,
			To:   r.To,
		},
	}
}

// IDの再登録
func idReregistration(points []*model.Place) {
	sort.Slice(points, func(i, j int) bool {
		if points[i].Coord.X == points[j].Coord.X {
			return points[i].Coord.Y < points[j].Coord.Y
		}
		return points[i].Coord.X < points[j].Coord.X
	})

	for i := range points {
		points[i].Id = fmt.Sprintf("C%d", i+1)
	}
}
