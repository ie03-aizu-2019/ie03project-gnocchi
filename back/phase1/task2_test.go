package phase1

import (
	"fmt"
	"sort"
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

	// p := sorting(crossPoints)
	var result string
	for _, value := range crossPoints {
		result += fmt.Sprintf("%s\n", value.Coord.ToString())
	}

	return result
}

func sorting(points map[string]*model.Place) []*model.Place {
	p := make([]*model.Place, len(points))

	i := 0
	for _, v := range points {
		p[i] = v
		i++
	}

	sort.Slice(p, func(i, j int) bool {
		return p[i].Id < p[j].Id
	})

	return p
}

func TestTask2Case1(t *testing.T) {
	utils.Assert("phase1/task2/case1", task2, t)
}
