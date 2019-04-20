package phase1

import (
	"fmt"
	"sort"

	"github.com/uzimaru0000/ie03project-gnocchi/back/model"
)

func EnumerateCrossPoint(roads []model.Road) []*model.Place {
	places := []*model.Place{}

	for i, road := range roads {
		for j := i + 1; j < len(roads); j++ {
			p, err := CheckCrossPoint(&road, &roads[j])
			if err == nil {
				places = append(places, &model.Place{Id: "", Coord: *p})
			}
		}
	}

	sort.Slice(places, func(i, j int) bool {
		if places[i].Coord.X == places[j].Coord.X {
			return places[i].Coord.Y < places[j].Coord.Y
		}
		return places[i].Coord.X < places[j].Coord.X
	})

	for i := range places {
		places[i].Id = fmt.Sprintf("C%d", i+1)
	}

	return places
}
