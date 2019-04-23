package phase1

import (
	"testing"

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

	places := EnumerateCrossPoint(datas.Roads)

	var result string
	for _, p := range places {
		result += p.ToString() + "\n"
	}

	return result
}

func TestTask2Case1(t *testing.T) {
	utils.Assert("phase1/task2/case1", task2, t)
}
