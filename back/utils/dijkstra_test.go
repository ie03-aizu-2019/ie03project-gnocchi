package utils

import (
	"log"
	"reflect"
	"testing"

	"github.com/uzimaru0000/ie03project-gnocchi/back/model"
)

func TestDijkstra(t *testing.T) {
	log.Println("in testDijkstra")

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
