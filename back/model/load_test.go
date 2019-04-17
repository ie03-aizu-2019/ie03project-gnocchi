package model

import (
	"os"
	"path/filepath"
	"testing"
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

	places := parcePlace(data)
	ans := []Place{
		{Id: 1, Coord: Point{X: 0, Y: 0}},
		{Id: 2, Coord: Point{X: 5, Y: 5}},
		{Id: 3, Coord: Point{X: 2, Y: 5}},
		{Id: 4, Coord: Point{X: 7, Y: 1}},
	}

	if len(data) != len(places) {
		t.Fatal("Not parseing")
	}

	for i, p := range places {
		if p != ans[i] {
			t.Fatalf("%d is not equal", i)
		}
	}
}

func TestParceRoad(t *testing.T) {
	data := []string{
		"1 2",
		"3 4",
	}

	places := []Place{
		{Id: 1, Coord: Point{X: 0, Y: 0}},
		{Id: 2, Coord: Point{X: 5, Y: 5}},
		{Id: 3, Coord: Point{X: 2, Y: 5}},
		{Id: 4, Coord: Point{X: 7, Y: 1}},
	}

	roads := parceRoad(data, places)
	ans := []Road{
		{
			Id:   1,
			From: Place{Id: 1, Coord: Point{X: 0, Y: 0}},
			To:   Place{Id: 2, Coord: Point{X: 5, Y: 5}},
		},
		{
			Id:   2,
			From: Place{Id: 3, Coord: Point{X: 2, Y: 5}},
			To:   Place{Id: 4, Coord: Point{X: 7, Y: 1}},
		},
	}

	if len(data) != len(roads) {
		t.Fatal("Not parseing")
	}

	for i, r := range roads {
		if r != ans[i] {
			t.Fatalf("%d is not equal", i)
		}
	}

}

func TestLoadFile(t *testing.T) {
	gopath := os.Getenv("GOPATH")
	path := filepath.Join(
		gopath,
		"src",
		"github.com",
		"uzimaru0000",
		"ie03project-gnocchi",
		"back",
		"test_data",
		"phase1",
		"case1.txt",
	)

	ans := datas{
		Places: []Place{
			{Id: 1, Coord: Point{X: 0, Y: 0}},
			{Id: 2, Coord: Point{X: 5, Y: 5}},
			{Id: 3, Coord: Point{X: 2, Y: 5}},
			{Id: 4, Coord: Point{X: 7, Y: 1}},
		},
		Roads: []Road{
			{
				Id:   1,
				From: Place{Id: 1, Coord: Point{X: 0, Y: 0}},
				To:   Place{Id: 2, Coord: Point{X: 5, Y: 5}},
			},
			{
				Id:   2,
				From: Place{Id: 3, Coord: Point{X: 2, Y: 5}},
				To:   Place{Id: 4, Coord: Point{X: 7, Y: 1}},
			},
		},
	}

	datas, err := Load(path)
	if err != nil {
		t.Fatal(err.Error())
	}

	for i, place := range datas.Places {
		if place != ans.Places[i] {
			t.Fatalf("%d is not equal.", i)
		}
	}

	for i, road := range datas.Roads {
		if road != ans.Roads[i] {
			t.Fatalf("%d is not equal.", i)
		}
	}

}
