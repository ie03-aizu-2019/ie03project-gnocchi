package phase1

import (
	"testing"

	"github.com/uzimaru0000/ie03project-gnocchi/back/model"
	"github.com/uzimaru0000/ie03project-gnocchi/back/utils"
)

func task2(file string) ([]*model.Place, error) {
	path := utils.CreatePath(file)
	str, err := utils.Load(path)
	if err != nil {
		return nil, err
	}

	datas, err := utils.ParseData(str)
	if err != nil {
		return nil, err
	}

	places := EnumerateCrossPoint(datas.Roads)
	return places, nil
}

func TestTask2Case1(t *testing.T) {
	places, err := task2("phase1/task2/case1.txt")
	if err != nil {
		t.Fatal(err)
	}

	ans := []*model.Place{
		{
			Id: "C1",
			Coord: model.Point{
				X: 3.66667,
				Y: 3.66667,
			},
		},
		{
			Id: "C2",
			Coord: model.Point{
				X: 4.86885,
				Y: 2.70492,
			},
		},
		{
			Id: "C3",
			Coord: model.Point{
				X: 5.86957,
				Y: 3.26087,
			},
		},
	}

	for i, p := range places {
		if p.Id == ans[i].Id &&
			!(utils.NearEqual(p.Coord.X, ans[i].Coord.X) && utils.NearEqual(p.Coord.Y, ans[i].Coord.Y)) {
			t.Fatal("not equal")
			return
		}
	}

}
