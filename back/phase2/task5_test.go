package phase2

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/uzimaru0000/ie03project-gnocchi/back/model"
	"github.com/uzimaru0000/ie03project-gnocchi/back/phase1"
	"github.com/uzimaru0000/ie03project-gnocchi/back/utils"
)

func task5(file string) string {

	str, err := utils.Load(file)
	if err != nil {
		return err.Error()
	}

	datas, err := utils.ParseData(str)
	if err != nil {
		return err.Error()
	}

	roads, places := phase1.EnumerateCrossPoints(datas.Roads)

	result := ""

	for _, q := range datas.Queries {
		routes := calcKthShortestPath(*q, append(datas.Places, places...), roads)
		if len(routes) == 0 {
			result += fmt.Sprintln("NA")
			continue
		}
		for _, rs := range routes {
			sum := 0.0
			for _, r := range rs {
				// log.Println(r.To.Id, "->", r.From.Id)
				sum += r.Length()
			}
			result += fmt.Sprintf("%.5f\n", sum)
			// log.Println("----")
		}
		// log.Println("=====")
	}
	return result
}

func TestTask5Case1(t *testing.T) {
	utils.Assert("phase2/task5/case1", task5, t)
}

func TestTask5Case2(t *testing.T) {
	places := []*model.Place{
		&model.Place{Id: "0", Coord: model.Point{X: 0, Y: 0}},
		&model.Place{Id: "1", Coord: model.Point{X: 2, Y: 0}},
		&model.Place{Id: "2", Coord: model.Point{X: 0, Y: 1}},
		&model.Place{Id: "3", Coord: model.Point{X: 1, Y: 1}},
	}

	roads := []*model.Road{
		&model.Road{Id: 0, To: places[0], From: places[1]},
		&model.Road{Id: 1, To: places[0], From: places[2]},
		&model.Road{Id: 2, To: places[0], From: places[3]},
		&model.Road{Id: 3, To: places[2], From: places[3]},
		&model.Road{Id: 4, To: places[3], From: places[1]},
	}

	ans := [][]*model.Road{
		[]*model.Road{roads[2]},
		[]*model.Road{roads[1], roads[3]},
		[]*model.Road{roads[0], roads[4]},
	}

	result := calcKthShortestPath(model.Query{Start: "0", Dest: "3", Num: 3}, places, roads)

	if !reflect.DeepEqual(ans, result) {
		t.Fatal("task5 Case2 Not Equal")
	}
}

func TestSetVisit(t *testing.T) {
	mp := map[model.Place]([]*model.Road){}
	places := []*model.Place{
		&model.Place{Id: "0", Coord: model.Point{X: 0, Y: 0}},
		&model.Place{Id: "1", Coord: model.Point{X: 1, Y: 0}},
		&model.Place{Id: "2", Coord: model.Point{X: 0, Y: 1}},
		&model.Place{Id: "3", Coord: model.Point{X: 0, Y: 2}},
		&model.Place{Id: "4", Coord: model.Point{X: 2, Y: 2}},
	}
	roads := []*model.Road{
		&model.Road{Id: 0, To: places[0], From: places[1]},
		&model.Road{Id: 1, To: places[0], From: places[2]},
		&model.Road{Id: 2, To: places[1], From: places[2]},
		&model.Road{Id: 3, To: places[2], From: places[3]},
		&model.Road{Id: 4, To: places[3], From: places[4]},
	}
	mp[*places[0]] = []*model.Road{roads[0], roads[3]}
	mp[*places[2]] = []*model.Road{roads[1], roads[3]}

	add := []*model.Road{
		&model.Road{Id: 1, To: places[0], From: places[2]},
		&model.Road{Id: 3, To: places[2], From: places[3]},
	}

	ans := map[model.Place]([]*model.Road){
		*places[0]: []*model.Road{roads[0], roads[3], roads[1]},
		*places[2]: {roads[1], roads[3]},
		*places[3]: {roads[3]},
	}

	mp = setVisited(mp, add)

	if !reflect.DeepEqual(mp, ans) {
		t.Fatal("failed set visit")
	}
}

func TestAvoidPlaces(t *testing.T) {
	places := []*model.Place{
		&model.Place{Id: "0", Coord: model.Point{X: 0, Y: 0}},
		&model.Place{Id: "1", Coord: model.Point{X: 1, Y: 0}},
		&model.Place{Id: "2", Coord: model.Point{X: 0, Y: 1}},
		&model.Place{Id: "3", Coord: model.Point{X: 0, Y: 2}},
		&model.Place{Id: "4", Coord: model.Point{X: 2, Y: 2}},
	}

	avoid := []*model.Road{
		&model.Road{Id: 0, To: places[0], From: places[1]},
		&model.Road{Id: 1, To: places[0], From: places[2]},
	}

	anser := []*model.Place{
		&model.Place{Id: "0", Coord: model.Point{X: 0, Y: 0}},
		&model.Place{Id: "3", Coord: model.Point{X: 0, Y: 2}},
		&model.Place{Id: "4", Coord: model.Point{X: 2, Y: 2}},
	}

	result := avoidPlaces(places, avoid, places[0])

	if !reflect.DeepEqual(result, anser) {
		t.Fatal("test avoidPlace failed!")
	}
}

func TestAvoidRoads(t *testing.T) {
	places := []*model.Place{
		&model.Place{Id: "0", Coord: model.Point{X: 0, Y: 0}},
		&model.Place{Id: "1", Coord: model.Point{X: 1, Y: 0}},
		&model.Place{Id: "2", Coord: model.Point{X: 0, Y: 1}},
		&model.Place{Id: "3", Coord: model.Point{X: 0, Y: 2}},
		&model.Place{Id: "4", Coord: model.Point{X: 2, Y: 2}},
	}
	roads := []*model.Road{
		&model.Road{Id: 0, To: places[0], From: places[1]},
		&model.Road{Id: 1, To: places[0], From: places[2]},
		&model.Road{Id: 2, To: places[1], From: places[2]},
		&model.Road{Id: 3, To: places[2], From: places[3]},
		&model.Road{Id: 4, To: places[3], From: places[4]},
	}

	avoid := []*model.Road{
		&model.Road{Id: 0, To: places[0], From: places[1]},
		&model.Road{Id: 2, To: places[1], From: places[2]},
		&model.Road{Id: 4, To: places[3], From: places[4]},
	}

	anser := []*model.Road{
		&model.Road{Id: 0, To: places[0], From: places[1]},
		&model.Road{Id: 1, To: places[0], From: places[2]},
		&model.Road{Id: 3, To: places[2], From: places[3]},
	}

	result := avoidRoads(roads, avoid, places[0])

	if !reflect.DeepEqual(result, anser) {
		t.Fatal("test avoidRoads failed!")
	}
}

func TestNextPlace(t *testing.T) {
	places := []*model.Place{
		&model.Place{Id: "0", Coord: model.Point{X: 0, Y: 0}},
		&model.Place{Id: "1", Coord: model.Point{X: 1, Y: 0}},
	}
	r := &model.Road{Id: 0, To: places[0], From: places[1]}

	result := nextPlace(places[0], r)

	if !reflect.DeepEqual(result, places[1]) {
		t.Fatal("test nextPlace faild!")
	}
}

func TestAvoidRoad(t *testing.T) {
	places := []*model.Place{
		&model.Place{Id: "0", Coord: model.Point{X: 0, Y: 0}},
		&model.Place{Id: "1", Coord: model.Point{X: 1, Y: 0}},
		&model.Place{Id: "2", Coord: model.Point{X: 0, Y: 1}},
		&model.Place{Id: "3", Coord: model.Point{X: 0, Y: 2}},
		&model.Place{Id: "4", Coord: model.Point{X: 2, Y: 2}},
	}
	roads := []*model.Road{
		&model.Road{Id: 0, To: places[0], From: places[1]},
		&model.Road{Id: 1, To: places[0], From: places[2]},
		&model.Road{Id: 2, To: places[1], From: places[2]},
		&model.Road{Id: 3, To: places[2], From: places[3]},
		&model.Road{Id: 4, To: places[3], From: places[4]},
	}

	avoid := &model.Road{Id: 0, To: places[0], From: places[1]}

	ans := []*model.Road{
		&model.Road{Id: 1, To: places[0], From: places[2]},
		&model.Road{Id: 2, To: places[1], From: places[2]},
		&model.Road{Id: 3, To: places[2], From: places[3]},
		&model.Road{Id: 4, To: places[3], From: places[4]},
	}

	result := avoidRoad(roads, avoid)

	if !reflect.DeepEqual(result, ans) {
		t.Fatal("test avoidRoad failed!")
	}

}

func TestJoinRoads(t *testing.T) {
	places := []*model.Place{
		&model.Place{Id: "0", Coord: model.Point{X: 0, Y: 0}},
		&model.Place{Id: "1", Coord: model.Point{X: 1, Y: 0}},
		&model.Place{Id: "2", Coord: model.Point{X: 0, Y: 1}},
		&model.Place{Id: "3", Coord: model.Point{X: 0, Y: 2}},
		&model.Place{Id: "4", Coord: model.Point{X: 2, Y: 2}},
	}

	roads := [][]*model.Road{
		[]*model.Road{
			&model.Road{Id: 0, To: places[0], From: places[1]},
			&model.Road{Id: 1, To: places[0], From: places[2]},
			&model.Road{Id: 2, To: places[1], From: places[2]},
			&model.Road{Id: 3, To: places[2], From: places[3]},
		},
	}
	head := []*model.Road{
		&model.Road{Id: 0, To: places[0], From: places[1]},
		&model.Road{Id: 1, To: places[0], From: places[2]},
	}

	tail := [][]*model.Road{
		[]*model.Road{
			&model.Road{Id: 2, To: places[1], From: places[2]},
			&model.Road{Id: 3, To: places[2], From: places[3]},
		},
	}

	result := joinRoads(head, tail)
	if result == nil || !reflect.DeepEqual(roads, result) {
		t.Fatal("test joinRoads failed!")
	}
}
