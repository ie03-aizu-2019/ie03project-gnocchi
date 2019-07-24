package phase2

import (
	"container/heap"
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

func avoidRoads(base []*model.Road, avoids []*model.Road, ext *model.Place) (result []*model.Road) {
	for _, bs := range base {
		flg := true
		for _, av := range avoids {
			if !(reflect.DeepEqual(bs.To, ext) || reflect.DeepEqual(bs.From, ext)) {
				if reflect.DeepEqual(bs, av) {
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
		if reflect.DeepEqual(rs, item) {
			return false
		}
	}
	return true
}

func road2String(rs []*model.Road) (str string) {
	for _, r := range rs {
		str += (r.From).Id + "-" + r.To.Id
	}
	return
}

func setVisited(mp map[model.Place]([]*model.Road), roads []*model.Road) map[model.Place]([]*model.Road) {
	for _, r := range roads {
		if v, ok := mp[*r.To]; ok {
			flg := true
			for _, vr := range v {
				if *vr == *r {
					flg = false
				}
			}
			if flg {
				mp[*r.To] = append(mp[*r.To], r)
			}
		} else {
			mp[*r.To] = []*model.Road{r}
		}

		if v, ok := mp[*r.From]; ok {
			flg := true
			for _, vr := range v {
				if *vr == *r {
					flg = false
				}
			}
			if flg {
				mp[*r.From] = append(mp[*r.From], r)
			}
		} else {
			mp[*r.From] = []*model.Road{r}
		}
	}
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

	visited := map[model.Place]([]*model.Road){}

	shortests := utils.Dijkstra(start, places, roads)[*dest]
	pq := make(PriorityQueue, len(shortests))
	for i, ss := range shortests {
		setVisited(visited, ss)
		pq[i] = &Item{
			roads:    ss,
			priority: roadsLen(ss),
			index:    i,
		}
	}
	heap.Init(&pq)

	count := 0

	for pq.Len() > 0 && len(result) < k {
		baseItem, ok := pq.Pop().(*Item)
		if !ok {
			log.Print("pop failed")
		}
		baseRoute := baseItem.roads
		if isUniq(result, baseRoute) {
			result = append(result, baseRoute)
			visited = setVisited(visited, baseRoute)
		}
		count++
		if count > 100 {
			break
		}

		spurRoot := []*model.Road{}
		spurNode := start
		for _, r := range baseRoute {
			notWork := r
			dijp := avoidPlaces(places, spurRoot, spurNode)
			dijr := avoidRoads(roads, spurRoot, spurNode)
			dijr = avoidRoad(dijr, notWork)
			if v, ok := visited[*spurNode]; ok {
				for _, visitedRoad := range v {
					dijr = avoidRoad(dijr, visitedRoad)
				}
			}

			shortestPath := utils.Dijkstra(spurNode, dijp, dijr)[*dest]
			if len(shortestPath) == 0 {
				continue
			}

			for _, sp := range shortestPath {
				sproad := append(spurRoot, sp...)
				item := &Item{
					roads:    sproad,
					priority: roadsLen(sproad),
					index:    pq.Len(),
				}
				heap.Push(&pq, item)
			}
			spurNode = nextPlace(spurNode, r)
			spurRoot = append(spurRoot, r)
		}
	}

	return
}
