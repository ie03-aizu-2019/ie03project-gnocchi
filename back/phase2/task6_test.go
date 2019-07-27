package phase2

import (
	"fmt"
	"testing"

	"github.com/uzimaru0000/ie03project-gnocchi/back/phase1"
	"github.com/uzimaru0000/ie03project-gnocchi/back/utils"
)

func task6(file string) string {

	str, err := utils.Load(file)
	if err != nil {
		return err.Error()
	}

	datas, err := utils.ParseData(str)
	if err != nil {
		return err.Error()
	}

	roads, crossPoints := phase1.EnumerateCrossPoints(datas.Roads)
	roads = phase1.ConnectOnRoadPoints(roads, append(datas.Places, crossPoints...))

	result := ""
	for _, q := range datas.Queries {
		routes := calcKthShortestPath(*q, append(datas.Places, crossPoints...), roads)
		if len(routes) == 0 {
			result += fmt.Sprintln("NA")
			continue
		}
		for _, rs := range routes {
			sum := 0.0
			for _, r := range rs {
				sum += r.Length()
			}
			result += fmt.Sprintf("%s\n%s\n", toFormat(sum), pathToString(rs, q.Start))
		}
	}
	return result
}

func TestTask6Case1(t *testing.T) {
	utils.Assert("phase2/task6/case1", task6, t)
}
