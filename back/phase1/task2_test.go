package phase1

import (
	"fmt"
	"testing"

	"github.com/uzimaru0000/ie03project-gnocchi/back/model"
	"github.com/uzimaru0000/ie03project-gnocchi/back/utils"
)

func task2(file string) string {
	str, err := utils.Load(file)
	if err != nil {
		return err.Error()
	}

	datas, err := utils.ParseData(str)
	if err != nil {
		return err.Error()
	}

	_, crossPoints := EnumerateCrossPoints(datas.Roads)

	var result string
	for _, value := range crossPoints {
		result += fmt.Sprintf("%s\n", value.Coord.ToString())
	}

	return result
}

func TestOnTheRoad(t *testing.T) {
	point := &model.Place{
		Coord: model.Point{X: 5, Y: 5},
		Id:    "1",
	}
	road := &model.Road{
		From: &model.Place{
			Coord: model.Point{X: 4, Y: 7},
			Id:    "2",
		},
		To: &model.Place{
			Coord: model.Point{X: 7, Y: 1},
			Id:    "3",
		},
		Id: 0,
	}

	if !onTheRoad(point, road) {
		t.Fatalf("This point is on the road")
	}
}

func TestConnectOnRoadPoints(t *testing.T) {
	str, err := utils.Load(utils.CreatePath("phase1/task2/case1.txt"))
	if err != nil {
		t.Fatal(err.Error())
	}

	datas, err := utils.ParseData(str)
	if err != nil {
		t.Fatal(err.Error())
	}

	roads := ConnectOnRoadPoints(datas.Roads, datas.Places)
	for _, r := range roads {
		t.Logf("%s -> %s", r.From.Id, r.To.Id)
	}
}

func TestTask2Case1(t *testing.T) {
	utils.Assert("phase1/task2/case1", task2, t)
}
