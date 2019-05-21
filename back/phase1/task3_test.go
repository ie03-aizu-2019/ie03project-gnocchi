package phase1

import (
	"fmt"
	"testing"

	"github.com/uzimaru0000/ie03project-gnocchi/back/utils"
)

// TODO: 関数の実行結果をテストケースの答えの文字列になるように整形する
func task3(file string) string {
	str, err := utils.Load(file)
	if err != nil {
		return err.Error()
	}

	datas, err := utils.ParseData(str)
	if err != nil {
		return err.Error()
	}

	roads, places := EnumerateCrossPoints(datas.Roads)

	var result string
	for _, query := range datas.Queries {
		s, err := CalcShortestPath(*query, append(datas.Places, places...), roads)
		if err != nil {
			result += "NA\n"
		} else {
			result += fmt.Sprintf("%.5f\n", s)
		}
	}
	return result
}

func TestTask3Case1(t *testing.T) {
	utils.Assert("phase1/task3/case1", task3, t)
}
