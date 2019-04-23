package phase1

import (
	"testing"

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

	point, err := CheckCrossPoint(&datas.Roads[0], &datas.Roads[1])
	if err != nil {
		return err.Error()
	}

	return point.ToString()
}

func TestTask1Case1(t *testing.T) {
	utils.Assert("phase1/task1/case1", task1, t)
}

func TestTask2Case2(t *testing.T) {
	utils.Assert("phase1/task1/case2", task1, t)
}

func TestTask3Case3(t *testing.T) {
	utils.Assert("phase1/task1/case3", task1, t)
}
