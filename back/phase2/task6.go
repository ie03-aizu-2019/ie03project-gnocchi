package phase2

import "github.com/uzimaru0000/ie03project-gnocchi/back/model"

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
