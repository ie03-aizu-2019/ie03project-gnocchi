package phase1

import (
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

	roads := EnumerateCrossPoints(datas.Roads)

	var result string
	crossPoints := make(map[string]*model.Place)
	for _, road := range roads {
		if road.From.Id[0] == 'C' {
			crossPoints[road.From.Id] = road.From
		}
	}

	p := sorting(crossPoints)
	for _, value := range p {
		result += value.Coord.ToString() + "\n"
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
