package phase2

import (
	"github.com/uzimaru0000/ie03project-gnocchi/back/model"
	"github.com/uzimaru0000/ie03project-gnocchi/back/utils"
)

func roadFilter(list []*model.Road, removeList []*model.Road) (result []*model.Road) {
	for _, l := range list {
		flg := true
		for _, rl := range removeList {
			if l == rl {
				flg = false
			}
		}
		if flg {
			result = append(result, l)
		}
	}
	return
}

func placeFilter(list []*model.Place, removeList []*model.Place) (result []*model.Place) {
	for _, l := range list {
		flg := true
		for _, rl := range removeList {
			if l == rl {
				flg = false
			}
		}
		if flg {
			result = append(result, l)
		}
	}
	return
}

func calcKthShortestPath(q model.Query, places []*model.Place, roads []*model.Road) (result [][]*model.Road) {
	var start *model.Place
	var dest *model.Place

	for _, p := range places {
		if p.Id == q.Start {
			start = p
		}
		if p.Id == q.Dest {
			dest = p
		}
	}

	result = utils.Dijkstra(start, places, roads)[*dest]

	return
}
