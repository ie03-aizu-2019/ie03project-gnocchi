package phase2

import (
	"fmt"
	"strings"
	"testing"

	"github.com/uzimaru0000/ie03project-gnocchi/back/model"
	"github.com/uzimaru0000/ie03project-gnocchi/back/phase1"
	"github.com/uzimaru0000/ie03project-gnocchi/back/utils"
)

func toFormat(num float64) string {
	k := len(fmt.Sprintf("%d", int(num)))
	temp := fmt.Sprintf("%%.%df", 6-k)
	return strings.TrimRight(fmt.Sprintf(temp, num), "0")
}

func task5(file string) string {

	str, err := utils.Load(file)
	if err != nil {
		return err.Error()
	}

	datas, err := utils.ParseData(str)
	if err != nil {
		return err.Error()
	}

	roads := phase1.ConnectOnRoadPoints(datas.Roads, datas.Places)

	roads, places := phase1.EnumerateCrossPoints(roads)

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
				sum += r.Length()
			}
			result += toFormat(sum) + "\n"
		}
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

	if len(ans) != len(result) {
		t.Fatal("task5 Case2 Not Equal")
	} else {
		for i := range ans {
			if len(ans[i]) == len(result[i]) {
				for j := range ans[i] {
					if ans[i][j].Id != result[i][j].Id {
						t.Fatal("task5 Case2 Not Equal")
					}
				}
			} else {
				t.Fatal("task5 Case2 Not Equal")

			}
		}
	}

}

func TestTask5Case3(t *testing.T) {
	places := []*model.Place{
		&model.Place{"1", model.Point{0, 0}},
		&model.Place{"2", model.Point{0, 1}},
		&model.Place{"3", model.Point{2, 0}},
		&model.Place{"4", model.Point{1, 1}},
		&model.Place{"5", model.Point{1, 3}},
		&model.Place{"6", model.Point{3, 1}},
		&model.Place{"7", model.Point{3, 2}},
	}

	roads := []*model.Road{
		&model.Road{0, places[0], places[1]},
		&model.Road{1, places[2], places[0]},
		&model.Road{2, places[3], places[1]},
		&model.Road{3, places[3], places[2]},
		&model.Road{4, places[3], places[4]},
		&model.Road{5, places[4], places[6]},
		&model.Road{6, places[5], places[6]},
		&model.Road{7, places[5], places[3]},
	}

	q := model.Query{Start: "1", Dest: "5", Num: 4}

	ans := [][]*model.Road{
		[]*model.Road{roads[0], roads[2], roads[4]},
		[]*model.Road{roads[1], roads[3], roads[4]},
		[]*model.Road{roads[0], roads[2], roads[7], roads[6], roads[5]},
		[]*model.Road{roads[1], roads[3], roads[7], roads[6], roads[5]},
	}

	result := calcKthShortestPath(q, places, roads)

	if len(ans) != len(result) {
		t.Fatal("not equal task5Case3")
	} else {
		for i := range ans {
			if len(ans[i]) == len(result[i]) {
				for j := range ans[i] {
					if ans[i][j].Id != result[i][j].Id {
						t.Fatal("not equal task5Case3")
					}
				}
			} else {
				t.Fatal("not equal task5Case3")

			}
		}
	}

}

func TestSetVisit(t *testing.T) {
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
	mp := map[string]([]*model.Road){}
	ans := map[string]([]*model.Road){}

	key1 := roads[1:3]
	val1 := []*model.Road{roads[0], roads[3]}

	key2 := roads[2:4]
	val2 := []*model.Road{roads[4]}

	key3 := []*model.Road{
		&model.Road{Id: 2, To: places[1], From: places[2]},
		&model.Road{Id: 3, To: places[2], From: places[3]},
	}
	val3 := roads[4]

	mp[road2String(key1)] = val1
	ans[road2String(key1)] = val1
	ans[road2String(key2)] = val2

	mp = setVisited(mp, key3, val3)

	for k := range ans {
		if len(ans[k]) == len(mp[k]) {
			for i := range ans[k] {
				if ans[k][i].Id != mp[k][i].Id {
					t.Fatal("not Equal!")
				}
			}
		} else {
			t.Fatal("not Equal!")
		}
	}
}

func TestRoad2String(t *testing.T) {
	places := []*model.Place{
		&model.Place{Id: "0", Coord: model.Point{X: 0, Y: 0}},
		&model.Place{Id: "1", Coord: model.Point{X: 1, Y: 0}},
		&model.Place{Id: "2", Coord: model.Point{X: 0, Y: 1}},
		&model.Place{Id: "3", Coord: model.Point{X: 0, Y: 2}},
		&model.Place{Id: "4", Coord: model.Point{X: 2, Y: 2}},
	}
	roads1 := []*model.Road{
		&model.Road{Id: 0, To: places[0], From: places[1]},
		&model.Road{Id: 1, To: places[0], From: places[2]},
		&model.Road{Id: 2, To: places[1], From: places[2]},
		&model.Road{Id: 3, To: places[2], From: places[3]},
		&model.Road{Id: 4, To: places[3], From: places[4]},
	}
	roads2 := []*model.Road{
		&model.Road{Id: 0, To: places[0], From: places[1]},
		&model.Road{Id: 1, To: places[0], From: places[2]},
		&model.Road{Id: 2, To: places[1], From: places[2]},
		&model.Road{Id: 3, To: places[2], From: places[3]},
		&model.Road{Id: 4, To: places[3], From: places[4]},
	}

	if road2String(roads1) != road2String(roads2) {
		t.Fatal("not Equal!")
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

	ext := &model.Place{Id: "0", Coord: model.Point{X: 0, Y: 0}}

	result := avoidPlaces(places, avoid, ext)

	for i := range anser {
		if i > len(result) || anser[i].Id != result[i].Id {
			t.Fatal("test avoidPlace failed!")
		}
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
		&model.Road{Id: 1, To: places[0], From: places[2]},
		&model.Road{Id: 3, To: places[2], From: places[3]},
	}

	result := avoidRoads(roads, avoid)

	for i := range anser {
		if i > len(result) || anser[i].Id != result[i].Id {
			t.Fatal("test avoidPlace failed!")
		}
	}
}

func TestAvoidSpurRoot(t *testing.T) {
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
		&model.Road{Id: 2, To: places[1], From: places[2]},
		&model.Road{Id: 3, To: places[2], From: places[3]},
	}

	anser := []*model.Road{
		&model.Road{Id: 3, To: places[2], From: places[3]},
		&model.Road{Id: 4, To: places[3], From: places[4]},
	}

	result := avoidSpurRoot(roads, avoid, places[3])

	for i := range anser {
		if i > len(result) || anser[i].Id != result[i].Id {
			t.Fatal("avoid SpurRoot failed")
		}
	}

}

func TestNextPlace(t *testing.T) {
	places := []*model.Place{
		&model.Place{Id: "0", Coord: model.Point{X: 0, Y: 0}},
		&model.Place{Id: "1", Coord: model.Point{X: 1, Y: 0}},
	}
	r := &model.Road{Id: 0, To: places[0], From: places[1]}

	result := nextPlace(places[0], r)

	if result.Id != places[1].Id {
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

	anser := []*model.Road{
		&model.Road{Id: 1, To: places[0], From: places[2]},
		&model.Road{Id: 2, To: places[1], From: places[2]},
		&model.Road{Id: 3, To: places[2], From: places[3]},
		&model.Road{Id: 4, To: places[3], From: places[4]},
	}

	result := avoidRoad(roads, avoid)

	for i := range anser {
		if i > len(result) || anser[i].Id != result[i].Id {
			t.Fatal("test avoidRoad failed!")
		}
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

	for i := range roads {
		if i < len(result) {
			for j := range roads[i] {
				if len(result[i]) < j || result[i][j].Id != roads[i][j].Id {
					t.Fatal("test joinRoads failed!")
				}
			}
		} else {
			t.Fatal("test joinRoads failed!")
		}
	}
}

func TestIsUniq(t *testing.T) {
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

	aRoads := []*model.Road{
		&model.Road{Id: 0, To: places[0], From: places[1]},
		&model.Road{Id: 1, To: places[0], From: places[2]},
		&model.Road{Id: 2, To: places[1], From: places[2]},
		&model.Road{Id: 3, To: places[2], From: places[3]},
	}

	if isUniq(roads, aRoads) {
		t.Fatal("isUnique failed")
	}

}
