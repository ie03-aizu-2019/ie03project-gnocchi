package phase1

import (
	"errors"
	"fmt"

	"github.com/uzimaru0000/ie03project-gnocchi/back/model"
	"github.com/uzimaru0000/ie03project-gnocchi/back/utils"
)

func CalcShortestPath(q model.Query, places []*model.Place, roads []*model.Road) (float64, error) {
	var start, dest *model.Place
	for _, p := range places {
		if q.Start == p.Id {
			start = p
		}
		if q.Dest == p.Id {
			dest = p
		}
	}
	if start == nil || dest == nil {
		return 0, errors.New("NA")
	}

	routes := utils.Dijkstra(start, places, roads)
	if len(routes[dest]) == 0 {
		return 0, errors.New("NA")
	}

	dist := 0.0
	for _, road := range routes[dest][0] {
		dist += road.Length()
	}

	return dist, nil
}

func roadsToString(rs []*model.Road) string {
	var result string
	for _, r := range rs {
		result += fmt.Sprintf("%s\n", r.ToString())
	}
	return result
}
