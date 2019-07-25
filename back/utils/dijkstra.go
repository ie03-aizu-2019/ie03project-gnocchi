package utils

import (
	"container/heap"
	"errors"
	"log"

	"github.com/uzimaru0000/ie03project-gnocchi/back/model"
)

type Item struct {
	Place    model.Place
	Priority float64
	Index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

// we want the ascending-prioirty queue
// so we use less than here
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq PriorityQueue) Push(x interface{}) {
	n := len(pq)
	item := x.(*Item)
	item.Index = n
	pq = append(pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

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

// What is do this function?
// road:{ to:item, from:reachItem}
// from:item -> to:reachItem
// Update to point distance(priority), if it is shorter than the distance to next point.
// if to point distance equal the form distance, push to routes.
func routeUpdate(item *Item, reachItem *Item, road *model.Road, routes *(map[model.Place]([][]*model.Road))) {
	if dist := item.Priority + road.Length(); reachItem.Priority > dist {
		reachItem.update(dist)
		var route [][]*model.Road
		if len((*routes)[item.Place]) > 0 {
			for _, r := range (*routes)[item.Place] {
				route = append(route, append(r, road))
			}
		} else {
			route = [][]*model.Road{{road}}
		}
		(*routes)[reachItem.Place] = route

	} else if reachItem.Priority == dist {
		for _, r := range (*routes)[item.Place] {
			(*routes)[reachItem.Place] = append((*routes)[reachItem.Place], append(r, road))
		}
		if len((*routes)[item.Place]) == 0 {
			(*routes)[reachItem.Place] = append((*routes)[reachItem.Place], []*model.Road{road})
		}
	}
}

// Dijkstra is algorithm search shortest path to each node from start node
// args: start, places, roads, return:  map that key is each place which value is shortest route from start
func Dijkstra(start *model.Place, places []*model.Place, roads []*model.Road) map[model.Place]([][]*model.Road) {
	var inf = 1e30
	pq := make(PriorityQueue, len(places))

	routes := make(map[model.Place]([][]*model.Road), len(places))

	flg := true
	for i, place := range places {
		pq[i] = &Item{
			Place:    *place,
			Priority: inf,
			Index:    i,
		}
		routes[*place] = make([][]*model.Road, 0)

		if place.Id == start.Id {
			flg = false
			pq[i].Priority = 0
		}
	}
	if flg {
		log.Println("start not found")
		return routes
	}
	heap.Init(&pq)

	for pq.Len() > 0 {
		i := heap.Pop(&pq).(*Item)
		if i.Priority == inf {
			break
		}
		for _, road := range roads {
			if road.To.Id == i.Place.Id {

				if reachItem, err := pq.find(*road.From); err == nil {
					routeUpdate((*Item)(i), reachItem, road, &routes)
				}
			}
			if road.From.Id == i.Place.Id {
				if reachItem, err := pq.find(*road.To); err == nil {
					routeUpdate((*Item)(i), reachItem, road, &routes)
				}
			}
		}
		heap.Init(&pq)
	}
	return routes
}
