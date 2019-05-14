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

func (pq *PriprityQueue) update(dist float64, place model.Place) {
	for _, item := range *pq {
		if item.place == place {
			if item.priority > dist {
				item.priority = dist
			}
		}
	}
}

func CalcShortedtPath(q model.Query, places []*model.Place, roads []*model.Road) (float64, error) {
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

	route := dijkstra(start, places, roads)
	if *route[len(roads)-1].To != *dest {
		return 0, errors.New("NA")
	}

	dist := 0.0
	for _, r := range route {
		dist += r.Length()
	}
	return dist, nil
}

func dijkstra(start *model.Place, places []*model.Place, roads []*model.Road) []*model.Road {
	var inf float64 = 1e30
	pq := make(PriprityQueue, len(places))
	shortests := make(map[model.Place]float64, len(places))
	for i, place := range places {
		pq[i] = &Item{
			place:    *place,
			priority: inf,
			index:    i,
		}
		shortests[*place] = inf
		if *place == *start {
			pq[i].priority = 0
		}
	}
	// shortests[*start] = 0
	heap.Init(&pq)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		// log.Printf(toString(item))
		shortests[item.place] = item.priority
		for _, road := range roads {
			if *road.To == item.place {
				pq.update(item.priority+road.Length(), *road.From)
			}
			if *road.From == item.place {
				pq.update(item.priority+road.Length(), *road.To)
			}
		}
		heap.Init(&pq)
	}
	return shortests
}

func roadsToString(rs []*model.Road) string {
	var result string
	for _, r := range rs {
		result += fmt.Sprintf("%s\n", r.ToString())
	}
	return result
}
