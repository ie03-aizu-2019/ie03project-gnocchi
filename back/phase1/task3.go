package phase1

import (
	"container/heap"
	"errors"
	"fmt"

	"github.com/uzimaru0000/ie03project-gnocchi/back/model"
	"github.com/uzimaru0000/ie03project-gnocchi/back/utils"
)

type priorityQueue utils.PriorityQueue
type item utils.Item

func (pq *priorityQueue) find(place model.Place) (*item, error) {
	for _, i := range *pq {
		if i.Place == place {
			return (*item)(i), nil
		}
	}
	return nil, errors.New("Not found")
}

func (item *item) update(dist float64) {
	item.Priority = dist
}

func routeUpdate(item *item, reachItem *item, road *model.Road, routes *(map[model.Place]([][]model.Road))) {
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

	routes := dijkstra(start, places, roads)
	if len(routes[*dest]) == 0 {
		return 0, errors.New("NA")
	}

	dist := 0.0
	for _, road := range routes[*dest][0] {
		dist += road.Length()
	}

	return dist, nil
}

func dijkstra(start *model.Place, places []*model.Place, roads []*model.Road) map[model.Place]([][]model.Road) {
	var inf float64 = 1e30
	pq := make(utils.PriorityQueue, len(places))

	routes := make(map[model.Place]([][]model.Road), len(places))

	for i, place := range places {
		pq[i] = &utils.Item{
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
		i := heap.Pop(&pq).(*utils.Item)

		for _, road := range roads {
			if *road.To == i.Place {

				q := priorityQueue(pq)
				if reachItem, err := q.find(*road.From); err == nil {
					routeUpdate((*item)(i), reachItem, road, &routes)
				}
			}
			if *road.From == i.Place {
				q := priorityQueue(pq)
				if reachItem, err := q.find(*road.To); err == nil {
					routeUpdate((*item)(i), reachItem, road, &routes)
				}
			}
		}
		heap.Init(&pq)
	}
	return routes
}

func roadsToString(rs []*model.Road) string {
	var result string
	for _, r := range rs {
		result += fmt.Sprintf("%s\n", r.ToString())
	}
	return result
}
