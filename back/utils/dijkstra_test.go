package utils

import (
	"reflect"
	"testing"

	"github.com/uzimaru0000/ie03project-gnocchi/back/model"
)

func aTestDijkstraCase1(t *testing.T) {

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

	anser := map[*model.Place]([][]*model.Road){
		places[0]: {{roads[0]}},
		places[1]: {},
		places[2]: {{roads[2]}},
		places[3]: {{roads[2], roads[3]}},
		places[4]: {},
	}

	result := Dijkstra(places[1], places, roads)

	for k, v := range anser {
		if !reflect.DeepEqual(v, result[k]) {
			t.Fatal("testdijkstra not equal")
		}
	}

}

func aTestDijkstraCase2(t *testing.T) {
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

	anser := map[*model.Place]([][]*model.Road){
		places[0]: {{roads[0]}},
		places[1]: {},
		places[2]: {{roads[2], roads[3]}},
		places[3]: {{roads[2]}},
		places[4]: {{roads[2], roads[4]}},
		places[5]: {{roads[2], roads[7]}},
		places[6]: {{roads[2], roads[7], roads[6]}},
	}

	result := Dijkstra(places[1], places, roads)

	for k, v := range anser {
		if !reflect.DeepEqual(v, result[k]) {
			t.Fatal("testdijkstra not equal")
		}
	}

}

func aTestDijkstraCase3(t *testing.T) {
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

	result := Dijkstra(places[1], places, roads)[places[4]]

	if len(result) != 0 {
		t.Fatal("dijkstraCase3 Failed")
	}
}

func TestDijkstraCase4(t *testing.T) {
	places := []*model.Place{
		&model.Place{"1", model.Point{0.00000, 0.00000}},
		&model.Place{"2", model.Point{2.00000, 5.00000}},
		&model.Place{"3", model.Point{4.00000, 7.00000}},
		&model.Place{"4", model.Point{8.04688, 5.70312}},
		&model.Place{"5", model.Point{7.00000, 1.00000}},
		&model.Place{"6", model.Point{9.00000, 5.00000}},
		&model.Place{"7", model.Point{4.37452, 3.10038}},
		&model.Place{"8", model.Point{4.86885, 2.70492}},
		&model.Place{"9", model.Point{5.53764, 3.92473}},
		&model.Place{"10", model.Point{5.86957, 3.26087}},
	}

	roads := []*model.Road{
		&model.Road{0, places[2], places[5]},
		// &model.Road{1, places[0], places[6]},
		&model.Road{2, places[1], places[6]},
		&model.Road{3, places[9], places[5]},
		&model.Road{4, places[9], places[4]},
		&model.Road{5, places[6], places[8]},
		&model.Road{6, places[8], places[3]},
		&model.Road{7, places[2], places[8]},
		&model.Road{8, places[8], places[9]},
		&model.Road{9, places[6], places[7]},
		&model.Road{10, places[7], places[4]},
		&model.Road{11, places[0], places[7]},
		&model.Road{12, places[7], places[9]},
	}

	result := Dijkstra(places[0], places, roads)[places[3]]

	ans := [][]*model.Road{
		[]*model.Road{roads[10], roads[11], roads[7], roads[5]},
	}

	if len(ans) != len(result) {
		t.Fatal("dijkstra failed 1")
	}
	for i := range result {
		if len(result[i]) == len(ans[i]) {
			for j := range result[i] {
				// log.Printf("%d %d\n", result[i][j].Id, ans[i][j].Id)
				if result[i][j].Id != ans[i][j].Id {
					t.Fatal("dijkstra failed 2")
				}
			}
		} else {
			t.Fatal("dijkstra failedn 3")
		}
	}

}
