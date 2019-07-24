package phase2

import (
	"fmt"
	"testing"

	"github.com/uzimaru0000/ie03project-gnocchi/back/phase1"
	"github.com/uzimaru0000/ie03project-gnocchi/back/utils"
)

func task8(file string) string {
	str, err := utils.Load(file)
	if err != nil {
		return err.Error()
	}

	datas, err := utils.ParseData(str)
	if err != nil {
		return err.Error()
	}

	roads, _ := phase1.EnumerateCrossPoints(datas.Roads)
	bridges := DetectBridge(roads)

	var result string
	for from, dests := range bridges {
		for _, to := range dests {
			if from.Id < to.Id {
				result += fmt.Sprintf("%s %s\n", from.Id, to.Id)
			} else {
				result += fmt.Sprintf("%s %s\n", to.Id, from.Id)
			}
		}
	}

	return result
}

func TestTask8Case1(t *testing.T) {
	utils.Assert("phase2/task8/case1", task8, t)
}

func TestTask8Case2(t *testing.T) {
	utils.Assert("phase2/task8/case2", task8, t)
}
