package utils

import (
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/uzimaru0000/ie03project-gnocchi/back/model"
)

func TestDijkstraCase1(t *testing.T) {

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

	anser := map[model.Place]([][]*model.Road){
		*places[0]: {{roads[0]}},
		*places[1]: {},
		*places[2]: {{roads[2]}},
		*places[3]: {{roads[2], roads[3]}},
		*places[4]: {},
	}

	result := Dijkstra(places[1], places, roads)

	for k, v := range anser {
		if !reflect.DeepEqual(v, result[k]) {
			t.Fatal("testdijkstra not equal")
		}
	}

}

func TestDijkstraCase2(t *testing.T) {
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

	anser := map[model.Place]([][]*model.Road){
		*places[0]: {{roads[0]}},
		*places[1]: {},
		*places[2]: {{roads[2], roads[3]}},
		*places[3]: {{roads[2]}},
		*places[4]: {{roads[2], roads[4]}},
		*places[5]: {{roads[2], roads[7]}},
		*places[6]: {{roads[2], roads[7], roads[6]}},
	}

	result := Dijkstra(places[1], places, roads)

	for k, v := range anser {
		if !reflect.DeepEqual(v, result[k]) {
			t.Fatal("testdijkstra not equal")
		}
	}

}

func TestDijkstraCase3(t *testing.T) {
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
		&model.Road{1, places[2], places[0]},
		&model.Road{3, places[3], places[2]},
		&model.Road{4, places[3], places[4]},
		&model.Road{5, places[4], places[6]},
		&model.Road{6, places[5], places[6]},
		&model.Road{7, places[5], places[3]},
	}

	result := Dijkstra(places[1], places, roads)[*places[1]]

	for _, rs := range result {
		str := ""
		for _, r := range rs {
			str += fmt.Sprintf("%d, ", r.Id)
		}
		log.Println("result: ", str)
	}

	if len(result) != 0 {
		t.Fatal("dijkstraCase3 Failed")
	}
}
