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
				sum += r.Length()
			}
			result += fmt.Sprintf("%.5f\n", sum)

		}

	}

	return result
}

func TestTask5Case1(t *testing.T) {
	utils.Assert("phase2/task5/case1", task5, t)
}

func TestAvoidPlaces(t *testing.T) {
	places := []*model.Place{
		&model.Place{"0", model.Point{0, 0}},
		&model.Place{"1", model.Point{1, 0}},
		&model.Place{"2", model.Point{0, 1}},
		&model.Place{"3", model.Point{0, 2}},
		&model.Place{"4", model.Point{2, 2}},
	}

	avoid := []*model.Road{
		&model.Road{2, places[1], places[2]},
		&model.Road{3, places[2], places[3]},
	}

	anser := []*model.Place{
		&model.Place{"0", model.Point{0, 0}},
		&model.Place{"4", model.Point{2, 2}},
	}

	result := avoidPlaces(places, avoid)

	if !reflect.DeepEqual(result, anser) {
		t.Fatal("test avoidPlace failed!")
	}
}

func TestAvoidRoads(t *testing.T) {
	places := []*model.Place{
		&model.Place{"0", model.Point{0, 0}},
		&model.Place{"1", model.Point{1, 0}},
		&model.Place{"2", model.Point{0, 1}},
		&model.Place{"3", model.Point{0, 2}},
		&model.Place{"4", model.Point{2, 2}},
	}
	roads := []*model.Road{
		&model.Road{0, places[0], places[1]},
		&model.Road{1, places[0], places[2]},
		&model.Road{2, places[1], places[2]},
		&model.Road{3, places[2], places[3]},
		&model.Road{4, places[3], places[4]},
	}

	avoid := []*model.Road{
		&model.Road{1, places[0], places[2]},
		&model.Road{2, places[1], places[2]},
	}

	anser := []*model.Road{
		&model.Road{4, places[3], places[4]},
	}

	result := avoidRoads(roads, avoid)

	if !reflect.DeepEqual(result, anser) {
		t.Fatal("test avoidRoads railed!")
	}
}

func TestJoinRoads(t *testing.T) {
	places := []*model.Place{
		&model.Place{"0", model.Point{0, 0}},
		&model.Place{"1", model.Point{1, 0}},
		&model.Place{"2", model.Point{0, 1}},
		&model.Place{"3", model.Point{0, 2}},
		&model.Place{"4", model.Point{2, 2}},
	}

	roads := []*model.Road{
		&model.Road{0, places[0], places[1]},
		&model.Road{1, places[0], places[2]},
		&model.Road{2, places[1], places[2]},
		&model.Road{3, places[2], places[3]},
	}

	tail := [][]*model.Road{
		roads[2:],
	}

	result := joinRoads(roads[:2], tail)
	if !reflect.DeepEqual(roads, result[0]) {
		t.Fatal("test joinRoads failed!")
	}
}
