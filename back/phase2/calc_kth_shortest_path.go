package phase2

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/uzimaru0000/ie03project-gnocchi/back/model"
)

type Item struct {
	roads    []*model.Road
	priority float64
	index    int
}

//-----------------------------------------------to debug------------
func (i Item) ToString() (result string) {
	for _, r := range i.roads {
		result += fmt.Sprintln(r.ToString())
	}
	result += fmt.Sprintf("len:%f index:%d\n", i.priority, i.index)
	return
}

func roadsToString(rs []*model.Road) (result string) {
	for _, r := range rs {
		result += fmt.Sprintf("to:%s from:%s -> \n", r.To.Id, r.From.Id)
	}
	result += fmt.Sprintln()
	return
}

func nextPlace(p *model.Place, r *model.Road) *model.Place {
	if *p == *r.From {
		return r.To
	}
	return r.From
}

// --------------------------------------------------------

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

func (pq *PriorityQueue) find(roads []*model.Road) (*Item, error) {
	for _, i := range *pq {
		if reflect.DeepEqual(i.roads, roads) {
			return (*Item)(i), nil
		}
	}
	return nil, errors.New("Not found")
}

func (item *Item) update(dist float64) {
	item.priority = dist
}

func roadFilter(list []*model.Road, remove *model.Road) (result []*model.Road) {
	for _, l := range list {
		if *remove != *l {
			result = append(result, l)
		}
	}
	return
}

func placeFilter(list []*model.Place, removeList []*model.Road) (result []*model.Place) {
	for _, p := range list {
		flg := true
		for _, r := range removeList {
			if p == r.To || p == r.From {
				flg = false
				break
			}
		}
		if flg {
			result = append(result, p)
		}
	}
	return
}

func rootLength(rs []*model.Road) (result float64) {
	for _, r := range rs {
		result += r.Length()
	}
	return
}

func calcKthShortestPath(q model.Query, places []*model.Place, roads []*model.Road) (result [][]*model.Road) {
	var start *model.Place
	var dest *model.Place
	k := q.Num

	for _, p := range places {
		if p.Id == q.Start {
			start = p
		}
		if p.Id == q.Dest {
			dest = p
		}
	}

	return
}
