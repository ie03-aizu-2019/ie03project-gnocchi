package utils

import (
	"testing"

	"github.com/uzimaru0000/ie03project-gnocchi/back/model"
)

func TestGetNums(t *testing.T) {
	data := "4 2 0 0"
	N, M, P, Q := getNums(data)

	if N != 4 || M != 2 || P != 0 || Q != 0 {
		t.Logf("N = %d", N)
		t.Logf("M = %d", M)
		t.Logf("P = %d", P)
		t.Logf("Q = %d", Q)
		t.Fatal("not parseing")
	}
}

func TestParcePlace(t *testing.T) {
	data := []string{
		"0 0",
		"5 5",
		"2 5",
		"7 1",
	}

	places := parcePlace(data, 0)
	ans := []*model.Place{
		{Id: "1", Coord: model.Point{X: 0, Y: 0}},
		{Id: "2", Coord: model.Point{X: 5, Y: 5}},
		{Id: "3", Coord: model.Point{X: 2, Y: 5}},
		{Id: "4", Coord: model.Point{X: 7, Y: 1}},
	}

	if len(data) != len(places) {
		t.Fatal("Not parseing")
	}

	for i, p := range places {
		if p.Coord != ans[i].Coord {
			t.Fatalf("%d is not equal", i)
		}
	}
}

func TestParceRoad(t *testing.T) {
	data := []string{
		"1 2",
		"3 4",
	}

	places := []*model.Place{
		{Id: "1", Coord: model.Point{X: 0, Y: 0}},
		{Id: "2", Coord: model.Point{X: 5, Y: 5}},
		{Id: "3", Coord: model.Point{X: 2, Y: 5}},
		{Id: "4", Coord: model.Point{X: 7, Y: 1}},
	}

	roads := parceRoad(data, places)
	ans := []*model.Road{
		{
			Id:   1,
			From: &model.Place{Id: "1", Coord: model.Point{X: 0, Y: 0}},
			To:   &model.Place{Id: "2", Coord: model.Point{X: 5, Y: 5}},
		},
		{
			Id:   2,
			From: &model.Place{Id: "3", Coord: model.Point{X: 2, Y: 5}},
			To:   &model.Place{Id: "4", Coord: model.Point{X: 7, Y: 1}},
		},
	}

	if len(data) != len(roads) {
		t.Fatal("Not parseing")
	}

	for i, r := range roads {
		if r.From.Id != ans[i].From.Id || r.To.Id != ans[i].To.Id {
			t.Fatalf("%d is not equal", i)
		}
	}

}

func TestQuery(t *testing.T) {
	data := []string{
		"1 4 1",
		"C1 6 1",
		"C1000 2 4",
	}

	ans := []*model.Query{
		{Start: "1", Dest: "4", Num: 1},
		{Start: "C1", Dest: "6", Num: 1},
		{Start: "C1000", Dest: "2", Num: 4},
	}

	queries := parceQuery(data)

	if len(data) != len(queries) {
		t.Fatal("Not parseing")
	}

	for i, q := range queries {
		if q.Start != ans[i].Start || q.Dest != ans[i].Dest || q.Num != ans[i].Num {
			t.Fatalf("%d is not equal", i)
		}
	}
}

func TestLoadFile(t *testing.T) {
	ans := Datas{
		Places: []*model.Place{
			{Id: "1", Coord: model.Point{X: 0, Y: 0}},
			{Id: "2", Coord: model.Point{X: 5, Y: 5}},
			{Id: "3", Coord: model.Point{X: 2, Y: 5}},
			{Id: "4", Coord: model.Point{X: 7, Y: 1}},
			{Id: "5", Coord: model.Point{X: 3, Y: 2}},
			{Id: "6", Coord: model.Point{X: 0, Y: 5}},
		},
		Roads: []*model.Road{
			{
				Id:   1,
				From: &model.Place{Id: "1", Coord: model.Point{X: 0, Y: 0}},
				To:   &model.Place{Id: "2", Coord: model.Point{X: 5, Y: 5}},
			},
			{
				Id:   2,
				From: &model.Place{Id: "3", Coord: model.Point{X: 2, Y: 5}},
				To:   &model.Place{Id: "4", Coord: model.Point{X: 7, Y: 1}},
			},
		},
		Queries: []*model.Query{
			{Start: "C1", Dest: "4", Num: 1},
			{Start: "2", Dest: "3", Num: 1},
		},
	}

	datas, err := ParseData("4 2 2 2\n0 0\n5 5\n2 5\n7 1\n1 2\n3 4\n3 2\n0 5\nC1 4 1\n2 3 1")
	if err != nil {
		t.Fatal(err.Error())
	}

	for i, place := range datas.Places {
		if place.Coord != ans.Places[i].Coord {
			t.Fatalf("Place %d is not equal.", i)
			return
		}
	}

	for i, road := range datas.Roads {
		if road.From.Coord != ans.Roads[i].From.Coord || road.To.Coord != ans.Roads[i].To.Coord {
			t.Fatalf("Road %d is not equal.", i)
			return
		}
	}

	for i, query := range datas.Queries {
		if query.Start != ans.Queries[i].Start || query.Dest != ans.Queries[i].Dest || query.Num != ans.Queries[i].Num {
			t.Fatalf("Query %d is not equal.", i)
			return
		}
	}
}
