package model

import (
	"container/heap"
	"testing"
)

func TestRoadQueue(t *testing.T) {
	roads := &Roads{
		&Road{
			Id: 1,
			From: &Place{
				Id:    "1",
				Coord: Point{0, 0},
			},
			To: &Place{
				Id:    "1",
				Coord: Point{20, 10},
			},
		},
		&Road{
			Id: 2,
			From: &Place{
				Id:    "1",
				Coord: Point{0, 0},
			},
			To: &Place{
				Id:    "2",
				Coord: Point{0, 10},
			},
		},
		&Road{
			Id: 3,
			From: &Place{
				Id:    "1",
				Coord: Point{0, 0},
			},
			To: &Place{
				Id:    "3",
				Coord: Point{20, 10},
			},
		},
	}

	heap.Init(roads)

	min := heap.Pop(roads).(*Road)

	if min.Id != 2 {
		t.Fatal("Not Equal")
	}
}
