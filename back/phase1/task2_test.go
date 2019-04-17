package phase1

import (
	"testing"

	"github.com/uzimaru0000/ie03project-gnocchi/back/model"
	"github.com/uzimaru0000/ie03project-gnocchi/back/utils"
)

func task2(file string) ([]*model.Point, error) {
	path := utils.CreatePath(file)
	datas, err := model.Load(path)
	if err != nil {
		return nil, err
	}

	point := EnumerateCrossPoint(datas.Roads)
	return point, nil
}

func TestTask2Case1(t *testing.T) {
	points, err := task2("phase1/task2/case1.txt")
	if err != nil {
		t.Fatal(err)
	}

	ans := []model.Point{
		model.Point{
			X: 3.66667,
			Y: 3.66667,
		},
		model.Point{
			X: 4.86885,
			Y: 2.70492,
		},
		model.Point{
			X: 5.86957,
			Y: 3.26087,
		},
	}

	for i, p := range points {
		if !(utils.NearEqual(p.X, ans[i].X) && utils.NearEqual(p.Y, ans[i].Y)) {
			t.Fatal("not equal")
			return
		}
	}

}
