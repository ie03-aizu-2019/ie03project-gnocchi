package phase2

import (
	"reflect"
	"testing"

	"github.com/uzimaru0000/ie03project-gnocchi/back/model"
)

func TestRoadFilter(t *testing.T) {

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

	rmRoad := &model.Road{2, places[1], places[2]}

	ans := []*model.Road{
		&model.Road{0, places[0], places[1]},
		&model.Road{1, places[0], places[2]},
		&model.Road{3, places[2], places[3]},
	}

	result := roadFilter(roads, rmRoad)

	if !reflect.DeepEqual(ans, result) {
		t.Fatal("fail roadsfilter")
	}

}

func TestPlaceFilter(t *testing.T) {
	places := []*model.Place{
		&model.Place{"0", model.Point{0, 0}},
		&model.Place{"1", model.Point{1, 0}},
		&model.Place{"2", model.Point{0, 1}},
		&model.Place{"3", model.Point{0, 2}},
		&model.Place{"4", model.Point{2, 2}},
	}

	rmRoads := []*model.Road{
		&model.Road{0, places[0], places[1]},
		&model.Road{2, places[1], places[2]},
	}

	ans := []*model.Place{
		&model.Place{"3", model.Point{0, 2}},
		&model.Place{"4", model.Point{2, 2}},
	}
	result := placeFilter(places, rmRoads)

	if !reflect.DeepEqual(ans, result) {
		t.Fatal("not equal place filter")
	}
}
