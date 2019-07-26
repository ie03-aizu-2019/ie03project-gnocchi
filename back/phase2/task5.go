package phase2

import (
	"container/heap"
	"fmt"
	"log"
	"reflect"

	"github.com/uzimaru0000/ie03project-gnocchi/back/model"
	"github.com/uzimaru0000/ie03project-gnocchi/back/utils"
)

// An Item is something we manage in a priority queue.
type Item struct {
	roads    []*model.Road
	priority float64
	index    int
}

//A PriorityQueue implements heap.Interface and holds Items.
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

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
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

func avoidPlaces(base []*model.Place, avoids []*model.Road, ext *model.Place) (result []*model.Place) {
	for _, bs := range base {
		flg := true
		for _, av := range avoids {
			if bs.Id != ext.Id {
				if reflect.DeepEqual(bs, av.To) || reflect.DeepEqual(bs, av.From) {
					flg = false
					break
				}
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

func avoidSpurRoot(base []*model.Road, avoid []*model.Road, expt *model.Place) (result []*model.Road) {
	for _, bs := range base {
		flg := true
		if !reflect.DeepEqual(bs.To, expt) && !reflect.DeepEqual(bs.From, expt) {
			for _, av := range avoid {
				if bs.From.Id == av.From.Id || bs.From.Id == av.To.Id || bs.To.Id == av.From.Id || bs.To.Id == av.To.Id {
					flg = false
				}
			}
		}
		if flg {
			result = append(result, bs)
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
	}
	return r.To

}

func roadsLen(rs []*model.Road) (result float64) {
	for _, r := range rs {
		result += r.Length()
	}
	return
}

func isUniq(base [][]*model.Road, item []*model.Road) bool {
	for _, rs := range base {
		flg := true
		for i := range rs {
			if len(item) <= i || rs[i].Id != item[i].Id {
				flg = false
				break
			}
		}
		if flg {
			return false
		}
	}

	return true
}

func road2String(rs []*model.Road) (str string) {
	for _, r := range rs {
		str += fmt.Sprintf("%d, ", r.Id)
	}
	return
}

func setVisited(mp map[string]([]*model.Road), key []*model.Road, value *model.Road) map[string]([]*model.Road) {
	keyString := road2String(key)
	mp[keyString] = append(mp[keyString], value)
	return mp
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

	visited := map[string]([]*model.Road){}

	shortests := utils.Dijkstra(start, places, roads)[*dest]
	pq := make(PriorityQueue, len(shortests))
	for i, ss := range shortests {
		for i, sr := range ss {
			if i > 0 {
				visited = setVisited(visited, ss[:i], sr)
			}
		}
		pq[i] = &Item{
			roads:    ss,
			priority: roadsLen(ss),
			index:    i,
		}
	}
	heap.Init(&pq)

	for pq.Len() > 0 && len(result) < k {

		baseItem, ok := heap.Pop(&pq).(*Item)

		if !ok {
			log.Print("pop failed")
		}
		baseRoute := baseItem.roads
		if isUniq(result, baseRoute) {
			result = append(result, baseRoute)
			for i, br := range baseRoute {
				visited = setVisited(visited, baseRoute[:i], br)
			}
		} else {
			continue
		}

		spurRoot := []*model.Road{}
		spurNode := start
		for _, r := range baseRoute {
			notWork := r
			dijp := places
			dijr := avoidSpurRoot(roads, spurRoot, spurNode)
			dijr = avoidRoad(dijr, notWork)
			if v, ok := visited[road2String(spurRoot)]; ok {
				for _, visitedRoad := range v {
					dijr = avoidRoad(dijr, visitedRoad)
				}
			}

			shortestPath := utils.Dijkstra(spurNode, dijp, dijr)[*dest]
			if len(shortestPath) == 0 {
				spurNode = nextPlace(spurNode, r)
				spurRoot = append(spurRoot, r)
				continue
			}

			for _, sp := range shortestPath {
				sproad := append(spurRoot, sp...)
				item := &Item{
					roads:    sproad,
					priority: roadsLen(sproad),
				}
				heap.Push(&pq, item)
			}
			spurNode = nextPlace(spurNode, r)
			spurRoot = append(spurRoot, r)
		}
	}

	return
}
