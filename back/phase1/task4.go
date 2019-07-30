package phase1

import (
	"errors"
	"sort"

	"github.com/uzimaru0000/ie03project-gnocchi/back/model"
	"github.com/uzimaru0000/ie03project-gnocchi/back/utils"
)

type route struct {
	path   []*model.Road
	length float64
}

func EnumerationShortestPath(q model.Query, places []*model.Place, roads []*model.Road) (*route, error) {
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
		return nil, errors.New("NA")
	}

	routes := utils.Dijkstra(start, places, roads)
	if len(routes[dest]) == 0 {
		return nil, errors.New("NA")
	}

	route := new(route)
	paths := routes[dest]
	sortPaths(paths, dest.Id)

	route.path = paths[0]
	for _, road := range routes[dest][0] {
		route.length += road.Length()
	}

	return route, nil
}

func sortPaths(paths [][]*model.Road, startID string) [][]*model.Road {
	names := make([]string, len(paths))
	for i, path := range paths {
		names[i] = pathToString(path, startID)
	}

	sort.Slice(paths, func(i, j int) bool {
		return names[i] < names[j]
	})

	return paths
}

func pathToString(path []*model.Road, startID string) string {
	current := startID
	result := startID

	for _, p := range path {
		if p.To.Id == current {
			current = p.From.Id
			result += " " + p.From.Id
		} else {
			current = p.To.Id
			result += " " + p.To.Id
		}
	}

	return result
}
