package phase1

import (
	"testing"

	"github.com/uzimaru0000/ie03project-gnocchi/back/model"

	"github.com/uzimaru0000/ie03project-gnocchi/back/utils"
)

func task1(file string) string {
	str, err := utils.Load(file)
	if err != nil {
		return err.Error()
	}
	datas, err := utils.ParseData(str)
	if err != nil {
		return err.Error()
	}

	point, err := CheckCrossPoint(datas.Roads[0], datas.Roads[1])
	if err != nil {
		return err.Error()
	}

	return point.ToString()
}

func TestTask1Case1(t *testing.T) {
	utils.Assert("phase1/task1/case1", task1, t)
}

func TestTask1Case2(t *testing.T) {
	utils.Assert("phase1/task1/case2", task1, t)
}

func TestTask1Case3(t *testing.T) {
	utils.Assert("phase1/task1/case3", task1, t)
}

func TestTask1Case4(t *testing.T) {
	utils.Assert("phase1/task1/case4", task1, t)
}

func TestTask1CaseX(t *testing.T) {
	road1 := &model.Road{
		Id: 1,
		From: &model.Place{
			Id:    "1",
			Coord: model.Point{5.0, 5.0},
		},
		To: &model.Place{
			Id:    "2",
			Coord: model.Point{9, 5},
		},
	}

	road2 := &model.Road{
		Id: 1,
		From: &model.Place{
			Id:    "1",
			Coord: model.Point{4, 7},
		},
		To: &model.Place{
			Id:    "2",
			Coord: model.Point{5.86957, 3.26087},
		},
	}

	_, err := CheckCrossPoint(road1, road2)
	if err == nil {
		t.Fatal("Fatal")
	}
}
