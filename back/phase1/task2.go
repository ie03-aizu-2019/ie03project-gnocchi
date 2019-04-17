package phase1

import "github.com/uzimaru0000/ie03project-gnocchi/back/model"

func EnumerateCrossPoint(roads []model.Road) []*model.Point {
	points := []*model.Point{}

	for i, road := range roads {
		for j := i + 1; j < len(roads); j++ {
			p, err := CheckCrossPoint(&road, &roads[j])
			if err == nil {
				points = append(points, p)
			}
		}
	}

	return points
}
