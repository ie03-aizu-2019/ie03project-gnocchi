package utils

import (
	"container/heap"
	"errors"

	"github.com/uzimaru0000/ie03project-gnocchi/back/model"
)

func (pq *PriorityQueue) find(place model.Place) (*Item, error) {
	for _, i := range *pq {
		if i.Place == place {
			return (*Item)(i), nil
		}
	}
	return nil, errors.New("Not found")
}

func (item *Item) update(dist float64) {
	item.Priority = dist
}

func routeUpdate(item *Item, reachItem *Item, road *model.Road, routes *(map[model.Place]([][]model.Road))) {
	if dist := item.Priority + road.Length(); reachItem.Priority > dist {
		reachItem.update(dist)
		var route [][]model.Road
		if len((*routes)[item.Place]) > 0 {
			for _, r := range (*routes)[item.Place] {
				route = append(route, append(r, *road))
			}
		} else {
			route = [][]model.Road{{*road}}
		}
		(*routes)[reachItem.Place] = route

	} else if reachItem.Priority == dist {
		for _, r := range (*routes)[item.Place] {
			(*routes)[reachItem.Place] = append((*routes)[reachItem.Place], append(r, *road))
		}
		if len((*routes)[item.Place]) == 0 {
			(*routes)[reachItem.Place] = append((*routes)[reachItem.Place], []model.Road{*road})
		}
	}
}

func Dijkstra(start *model.Place, places []*model.Place, roads []*model.Road) map[model.Place]([][]model.Road) {
	var inf float64 = 1e30
	pq := make(PriorityQueue, len(places))

	routes := make(map[model.Place]([][]model.Road), len(places))

	for i, place := range places {
		pq[i] = &Item{
			Place:    *place,
			Priority: inf,
			Index:    i,
		}
		routes[*place] = make([][]model.Road, 0, 5)

		if *place == *start {
			pq[i].Priority = 0
		}
	}
	heap.Init(&pq)

	for pq.Len() > 0 {
		i := heap.Pop(&pq).(*Item)

		for _, road := range roads {
			if *road.To == i.Place {

				if reachItem, err := pq.find(*road.From); err == nil {
					routeUpdate((*Item)(i), reachItem, road, &routes)
				}
			}
			if *road.From == i.Place {
				if reachItem, err := pq.find(*road.To); err == nil {
					routeUpdate((*Item)(i), reachItem, road, &routes)
				}
			}
		}
		heap.Init(&pq)
	}
	return routes
}
