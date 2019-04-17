package phase1

import (
	"testing"

	"github.com/uzimaru0000/ie03project-gnocchi/back/utils"

	"github.com/uzimaru0000/ie03project-gnocchi/back/model"
)

func pathSelectCase(file string) (*model.Point, error) {
	path := utils.CreatePath(file)
	datas, err := model.Load(path)
	if err != nil {
		return nil, err
	}

	point, err := CheckCrossPoint(&datas.Roads[0], &datas.Roads[1])
	if err != nil {
		return nil, err
	}

	return point, err
}

func TestCase1(t *testing.T) {
	point, err := pathSelectCase("phase1/case1.txt")
	ans := model.Point{X: 3.66667, Y: 3.66667}
	if err != nil {
		t.Fatal(err)
	}

	if !utils.NearEqual(point.X, ans.X) || !utils.NearEqual(point.Y, ans.Y) {
		t.Fatal("not equal")
	}
}

func TestCase2(t *testing.T) {
	_, err := pathSelectCase("phase1/case2.txt")
	if err == nil {
		t.Fatal("Failed")
	}

	if err.Error() != "NA" {
		t.Fatal("not equal")
	}
}

func TestCase3(t *testing.T) {
	_, err := pathSelectCase("phase1/case3.txt")
	if err == nil {
		t.Fatal("Failed")
	}

	if err.Error() != "NA" {
		t.Fatal("not equal")
	}
}
