package phase1

import (
	"container/heap"
	"errors"
	"fmt"

	"github.com/uzimaru0000/ie03project-gnocchi/back/model"
)

type Item struct {
	place    model.Place
	priority float64
	index    int
}

type PriprityQueue []*Item

func (pq PriprityQueue) Len() int { return len(pq) }

func (pq PriprityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

// we want the ascending-prioirty queue
// so we use less than here
func (pq PriprityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq PriprityQueue) Push(x interface{}) {
	n := len(pq)
	item := x.(*Item)
	item.index = n
	pq = append(pq, item)
}
func (pq *PriprityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriprityQueue) find(place model.Place) (*Item, error) {
	for _, item := range *pq {
		if item.place == place {
			return item, nil
		}
	}
	return nil, errors.New("Not found")
}

func (item *Item) update(dist float64) {
	item.priority = dist
}

func routeUpdate(item *Item, reachItem *Item, road *model.Road, routes *(map[model.Place]([][]model.Road))) {
	if dist := item.priority + road.Length(); reachItem.priority > dist {
		reachItem.update(dist)
		var route [][]model.Road
		if len((*routes)[item.place]) > 0 {
			for _, r := range (*routes)[item.place] {
				route = append(route, append(r, *road))
			}
		} else {
			route = [][]model.Road{{*road}}
		}
		(*routes)[reachItem.place] = route

	} else if reachItem.priority == dist {
		for _, r := range (*routes)[item.place] {
			(*routes)[reachItem.place] = append((*routes)[reachItem.place], append(r, *road))
		}
		if len((*routes)[item.place]) == 0 {
			(*routes)[reachItem.place] = append((*routes)[reachItem.place], []model.Road{*road})
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
	pq := make(PriprityQueue, len(places))

	routes := make(map[model.Place]([][]model.Road), len(places))

	for i, place := range places {
		pq[i] = &Item{
			place:    *place,
			priority: inf,
			index:    i,
		}
		routes[*place] = make([][]model.Road, 0, 5)

		if *place == *start {
			pq[i].priority = 0
		}
	}
	heap.Init(&pq)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)

		for _, road := range roads {
			if *road.To == item.place {
				if reachItem, err := pq.find(*road.From); err == nil {
					routeUpdate(item, reachItem, road, &routes)
				}
			}
			if *road.From == item.place {
				if reachItem, err := pq.find(*road.To); err == nil {
					routeUpdate(item, reachItem, road, &routes)
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
