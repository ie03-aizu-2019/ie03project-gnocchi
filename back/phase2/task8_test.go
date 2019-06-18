package phase2

import (
	"fmt"
	"testing"

	"github.com/uzimaru0000/ie03project-gnocchi/back/model"
	"github.com/uzimaru0000/ie03project-gnocchi/back/utils"
)

var road = model.Road{
	From: &model.Place{
		Coord: model.Point{X: 4, Y: 7},
	},
	To: &model.Place{
		Coord: model.Point{X: 7, Y: 1},
	},
}

var place = model.Place{
	Coord: model.Point{X: 5, Y: 4},
}

func task8(file string) string {
	str, err := utils.Load(file)
	if err != nil {
		return err.Error()
	}

	datas, err := utils.ParseData(str)
	if err != nil {
		return err.Error()
	}

	places := RecomendClossPoint(datas.Roads, datas.AddPlaces)

	var result string
	for _, p := range places {
		result += fmt.Sprintln(p.ToString())
	}

	return result
}

func TestCalcLineRate(t *testing.T) {
	expected := 0.466666666667
	actual := calcLineRate(&road, &place)

	if !utils.NearEqual(expected, actual) {
		t.Fatal("not equal", actual)
	}
}

func TestRateToPoint(t *testing.T) {
	expected := model.Point{X: 5.4, Y: 4.2}
	actual := rateToPoint(&road, 0.466666666667)

	if !utils.NearEqual(expected.X, actual.X) || !utils.NearEqual(expected.Y, actual.Y) {
		t.Fatal("not equal", actual)
	}
}

func TestDistance(t *testing.T) {
	expected := 1.414213562
	actual := distance(
		&model.Point{X: 0, Y: 0},
		&model.Point{X: 1, Y: 1},
	)

	if !utils.NearEqual(expected, actual) {
		t.Fatal("not qeual")
	}
}

func TestCase1(t *testing.T) {
	utils.Assert("phase2/task8/case1", task8, t)
}
