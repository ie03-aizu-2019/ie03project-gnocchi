package phase1

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/uzimaru0000/ie03project-gnocchi/back/model"

	"github.com/uzimaru0000/ie03project-gnocchi/back/utils"
)

func task4(file string) string {
	str, err := utils.Load(file)
	if err != nil {
		return err.Error()
	}

	datas, err := utils.ParseData(str)
	if err != nil {
		return err.Error()
	}

	roads, places := EnumerateCrossPoints(datas.Roads)

	result := ""

	for _, q := range datas.Queries {
		route, err := EnumerationShortestPath(*q, append(datas.Places, places...), roads)
		if err != nil {
			result += fmt.Sprintln("NA")
			continue
		}

		result += fmt.Sprintf("%.5f\n%s\n", route.length, pathToString(route.path, q.Start))
	}

	return result
}

func TestTask4Case1(t *testing.T) {
	utils.Assert("phase1/task4/case1", task4, t)
}

func TestSortPaths(t *testing.T) {
	path := [][]model.Road{
		{
			model.Road{
				To:   &model.Place{Id: "A"},
				From: &model.Place{Id: "B"},
			},
			model.Road{
				To:   &model.Place{Id: "B"},
				From: &model.Place{Id: "D"},
			},
			model.Road{
				To:   &model.Place{Id: "D"},
				From: &model.Place{Id: "F"},
			},
		},
		{
			model.Road{
				To:   &model.Place{Id: "A"},
				From: &model.Place{Id: "B"},
			},
			model.Road{
				To:   &model.Place{Id: "B"},
				From: &model.Place{Id: "C"},
			},
			model.Road{
				To:   &model.Place{Id: "C"},
				From: &model.Place{Id: "D"},
			},
			model.Road{
				To:   &model.Place{Id: "D"},
				From: &model.Place{Id: "F"},
			},
		},
		{
			model.Road{
				To:   &model.Place{Id: "A"},
				From: &model.Place{Id: "E"},
			},
			model.Road{
				To:   &model.Place{Id: "E"},
				From: &model.Place{Id: "D"},
			},
			model.Road{
				To:   &model.Place{Id: "D"},
				From: &model.Place{Id: "F"},
			},
		},
	}

	ans := [][]model.Road{
		path[1],
		path[0],
		path[2],
	}

	if !reflect.DeepEqual(sortPaths(path, "F"), ans) {
		t.Fatal("Sort Failed")
	}
}
