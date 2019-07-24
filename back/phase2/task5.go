package phase2

import (
	"container/heap"
	"log"
	"reflect"

	"github.com/uzimaru0000/ie03project-gnocchi/back/model"
	"github.com/uzimaru0000/ie03project-gnocchi/back/utils"
)

type Item struct {
	roads    []*model.Road
	priority float64
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

// we want the ascending-prioirty queue
// so we use less than here
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq PriorityQueue) Push(x interface{}) {
	n := len(pq)
	item := x.(*Item)
	item.index = n
	pq = append(pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, roads []*model.Road, priority float64) {
	item.roads = roads
	item.priority = priority
	heap.Fix(pq, item.index)
}

func avoidPlaces(base []*model.Place, avoids []*model.Road) (result []*model.Place) {
	for _, bs := range base {
		flg := true
		for _, av := range avoids {
			if reflect.DeepEqual(bs, av.To) || reflect.DeepEqual(bs, av.From) {
				flg = false
				break
			}
		}
		if flg {
			result = append(result, bs)
		}
	}
	return
}

func avoidRoads(base []*model.Road, avoids []*model.Road) (result []*model.Road) {
	for _, bs := range base {
		flg := true
		for _, av := range avoids {
			if reflect.DeepEqual(bs, av) {
				flg = false
				break
			}
		}
		if flg {
			result = append(result, bs)
		}
	}
	return
}

func avoidRoad(base []*model.Road, r *model.Road) (result []*model.Road) {
	for _, rs := range base {
		if !reflect.DeepEqual(rs, r) {
			result = append(result, rs)
		}
	}
	return
}

func joinRoads(spurRoads []*model.Road, roads [][]*model.Road) (result [][]*model.Road) {
	for _, rs := range roads {
		result = append(result, append(spurRoads, rs...))
	}
	return
}

func nextPlace(p *model.Place, r *model.Road) *model.Place {
	if r.To == p {
		return r.From
	} else {
		return r.To
	}
}

func roadsLen(rs []*model.Road) (result float64) {
	for _, r := range rs {
		result += r.Length()
	}
	return
}

func road2String(rs []*model.Road) (str string) {
	for _, r := range rs {
		str += (r.From).Id + "-" + r.To.Id
	}
	return
}

func calcKthShortestPath(q model.Query, places []*model.Place, roads []*model.Road) (result [][]*model.Road) {
	k := q.Num
	var start, dest *model.Place
	for _, p := range places {
		if q.Start == p.Id {
			start = p
		}
		if q.Dest == p.Id {
			dest = p
		}
	}

	shortests := utils.Dijkstra(start, places, roads)[*dest]
	pq := make(PriorityQueue, len(shortests))
	for i, ss := range shortests {
		pq[i] = &Item{
			roads:    ss,
			priority: roadsLen(ss),
			index:    i,
		}
	}
	heap.Init(&pq)

	for pq.Len() > 0 && len(result) < k {
		baseItem, ok := pq.Pop().(*Item)
		if !ok {
			log.Print("pop failed")
		}
		baseRoute := baseItem.roads
		result = append(result, baseRoute)

		spurRoads := []*model.Road{}
		spurNode := start
		for _, r := range baseRoute {
			notWork := r
			dijp := avoidPlaces(places, spurRoads)
			dijr := avoidRoads(roads, spurRoads)
			dijr = avoidRoad(dijr, notWork)

			shortestPath := utils.Dijkstra(spurNode, dijp, dijr)[*dest]

			for _, sp := range shortestPath {
				result = append(result, append(spurRoads, sp...))
			}
			spurNode = nextPlace(spurNode, r)
			spurRoads = append(spurRoads)
		}
	}

	return
}
