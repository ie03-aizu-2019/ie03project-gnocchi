package phase1

import (
	"fmt"
	"sort"

	"github.com/uzimaru0000/ie03project-gnocchi/back/utils"

	"github.com/uzimaru0000/ie03project-gnocchi/back/model"
)

func ConnectOnRoadPoints(roads []*model.Road, places []*model.Place) []*model.Road {
	i := 0
	for i < len(roads) {
		flag := false

		for _, p := range places {
			if onTheRoad(p, roads[i]) {
				r := roads[i]
				remove(&roads, i)
				roads = append(roads, makeRoads(r, p, len(roads))...)
			}
		}

		if !flag {
			i++
		}

	}

	roadIdReregistration(&roads)

	return roads
}

func EnumerateCrossPoints(roads []*model.Road) ([]*model.Road, []*model.Place) {
	crossPoints := []*model.Place{}

	i := 0
	for i < len(roads) {
		flag := false
		for j := i + 1; j < len(roads); j++ {
			road1 := *roads[i]
			road2 := *roads[j]
			p, err := CheckCrossPoint(&road1, &road2)
			if err == nil {
				// 交差している道を削除
				remove(&roads, j)
				remove(&roads, i)

				crossPoint := &model.Place{Id: "X", Coord: *p}
				crossPoints = append(crossPoints, crossPoint)

				roads = append(roads, makeRoads(&road1, crossPoint, len(roads))...)
				roads = append(roads, makeRoads(&road2, crossPoint, len(roads))...)

				flag = true
				break
			}
		}

		if !flag {
			i++
		}
	}

	idReregistration(crossPoints)
	roadIdReregistration(&roads)

	return roads, crossPoints
}

func remove(s *[]*model.Road, i int) {
	if i >= len(*s) {
		return
	}

	*s = append((*s)[:i], (*s)[i+1:]...)
}

func makeRoads(r *model.Road, p *model.Place, n int) []*model.Road {
	return []*model.Road{
		{
			Id:   n,
			From: r.From,
			To:   p,
		},
		{
			Id:   n + 1,
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

func onTheRoad(place *model.Place, road *model.Road) bool {
	placeToStart := &model.Point{
		X: road.From.Coord.X - place.Coord.X,
		Y: road.From.Coord.Y - place.Coord.Y,
	}
	placeToEnd := &model.Point{
		X: road.To.Coord.X - place.Coord.X,
		Y: road.To.Coord.Y - place.Coord.Y,
	}

	// どちらかの長さが０ならばその地点は始点か終点
	if utils.NearEqual(placeToStart.Length(), 0.0) || utils.NearEqual(placeToEnd.Length(), 0.0) {
		return false
	}

	// normalize
	p2sLen := placeToStart.Length()
	p2eLen := placeToEnd.Length()
	dot := (placeToStart.X*placeToEnd.X + placeToStart.Y*placeToEnd.Y) / (p2sLen * p2eLen)
	return utils.NearEqual(dot, -1)
}

func roadIdReregistration(roads *[]*model.Road) {
	for i, r := range *roads {
		r.Id = i
	}
}
