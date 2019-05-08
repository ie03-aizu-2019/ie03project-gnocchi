package phase1

import (
	"fmt"
	"testing"

	"github.com/uzimaru0000/ie03project-gnocchi/back/model"

	"github.com/uzimaru0000/ie03project-gnocchi/back/utils"
)

// TODO: 関数の実行結果をテストケースの答えの文字列になるように整形する
func task3(file string) string {
	str, err := utils.Load(file)
	if err != nil {
		return err.Error()
	}

	datas, err := utils.ParseData(str)
	if err != nil {
		return err.Error()
	}

	roads, places := EnumerateCrossPoints(datas.Roads)

	var result string
	for _, query := range datas.Queries {
		s, err := CalcShortedtPath(*query, append(datas.Places, places...), roads)
		if err != nil {
			result += "NA\n"
		} else {
			result += fmt.Sprintf("%.5f\n", s)
		}
	}
	return result
}

func TestTask3Case1(t *testing.T) {
	utils.Assert("phase1/task3/case1", task3, t)
}

func TestDijkstra(t *testing.T) {

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

	anser := map[model.Place]float64{
		*places[0]: 1.0,
		*places[1]: 0,
		*places[2]: 1.41421356237,
		*places[3]: 2.41421356237,
		*places[4]: 1e30,
	}

	result := dijkstra(places[1], places, roads)

	for k, v := range anser {
		if !utils.NearEqual(v, result[k]) {
			t.Fatal("testdijkstra not equal")
		}
	}

}
