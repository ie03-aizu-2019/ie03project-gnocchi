package phase2

import (
	"fmt"
	"math"
	"sort"

	"github.com/uzimaru0000/ie03project-gnocchi/back/utils"

	"github.com/uzimaru0000/ie03project-gnocchi/back/model"
)

// refs : https://qiita.com/boiledorange73/items/bcd4e150e7caa0210ee6
func calcLineRate(road *model.Road, place *model.Place) float64 {
	x1 := road.From.Coord.X
	y1 := road.From.Coord.Y
	a := road.To.Coord.X - x1
	b := road.To.Coord.Y - y1

	x0 := place.Coord.X
	y0 := place.Coord.Y
	rate := -(a*(x1-x0) + b*(y1-y0)) / (a*a + b*b)
	return math.Max(0, math.Min(rate, 1))
}

func rateToPoint(road *model.Road, rate float64) *model.Point {
	x1 := road.From.Coord.X
	y1 := road.From.Coord.Y
	a := road.To.Coord.X - x1
	b := road.To.Coord.Y - y1
	x0 := a*rate + x1
	y0 := b*rate + y1

	return &model.Point{X: x0, Y: y0}
}

func distance(place1 *model.Point, place2 *model.Point) float64 {
	return (&model.Point{X: place1.X - place2.X, Y: place1.Y - place2.Y}).Length()
}

// RecomendClossPoint : 現在の道から最短になる点を探す関数
func RecomendClossPoint(roads []*model.Road, places []*model.Place) []*model.Point {
	result := make([]*model.Point, len(places))

	for i, p := range places {
		result[i] = calcMinRoad(roads, p)
	}

	return result
}

// CreateRecomendRoads : 現在の道から最短になる点を探し，その点から道を作成する関数
func CreateRecomendRoads(places []*model.Place, roads []*model.Road, addedPlaces []*model.Place) ([]*model.Road, []*model.Place) {
	recomendRoads := []*model.Road{}
	recomendPlaces := []*model.Place{}

	points := RecomendClossPoint(roads, addedPlaces)
	recomendPlaceID := 1
	for i, p := range points {
		if same, place := isSamePoint(p, places); same {
			recomendRoads = append(recomendRoads, &model.Road{
				Id:   len(roads) + i,
				From: place,
				To:   addedPlaces[i],
			})
		} else {
			recomendPlace := &model.Place{
				Id:    fmt.Sprintf("R%d", recomendPlaceID),
				Coord: *p,
			}

			recomendRoads = append(recomendRoads, &model.Road{
				Id:   len(roads) + i,
				From: recomendPlace,
				To:   addedPlaces[i],
			})
			recomendPlaces = append(recomendPlaces, recomendPlace)

			recomendPlaceID++
		}
	}

	return recomendRoads, recomendPlaces
}

func calcMinRoad(roads []*model.Road, place *model.Place) *model.Point {
	// 全部の道に対して，最短経路の交点の座標を計算する
	minRoads := make([]*model.Point, len(roads))
	for i, r := range roads {
		rate := calcLineRate(r, place)
		minRoads[i] = rateToPoint(r, rate)
	}
	// placeと交点の距離を使ってソートする
	sort.SliceStable(minRoads, func(i, j int) bool {
		return distance(&place.Coord, minRoads[i]) < distance(&place.Coord, minRoads[j])
	})
	// 0番目の値を採用する（最小値）
	return minRoads[0]
}

func isSamePoint(point *model.Point, places []*model.Place) (bool, *model.Place) {
	for _, p := range places {
		if utils.NearEqual(point.X, p.Coord.X) && utils.NearEqual(point.Y, p.Coord.Y) {
			return true, p
		}
	}
	return false, nil
}
