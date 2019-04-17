package model

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func Load(path string) (*datas, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	data := &datas{}
	strs := strings.Split(string(b), "\n")
	N, M, P, Q := getNums(strs[0])

	data.Places = parcePlace(strs[1:N+1], 0)
	data.Roads = parceRoad(strs[N+1:M+N+1], data.Places)

	if P != 0 {
		data.Places = append(data.Places, parcePlace(strs[M+N+1:P+M+N+1], N)...)
	}

	if Q != 0 {
		data.Queries = parceQuery(strs[P+M+N+1:])
	}

	return data, nil
}

// 空白区切りで文字列をパースする
func getNums(plane string) (int, int, int, int) {
	nums := make([]int, 4)

	splited := strings.Split(plane, " ")
	for i, str := range splited {
		if n, err := strconv.Atoi(str); err != nil {
			nums[i] = 0
		} else {
			nums[i] = n
		}
	}

	return nums[0], nums[1], nums[2], nums[3]
}

// 文字列から座標をパースする
func parcePoint(str string) Point {
	x, y, _, _ := getNums(str)
	return Point{X: float64(x), Y: float64(y)}
}

// 文字列配列から地点をパースする
// plane : 入力された文字列
// n		 : idの始まり
func parcePlace(plane []string, n int) []Place {
	places := make([]Place, len(plane))

	for i, str := range plane {
		places[i] = Place{
			Id:    n + i + 1,
			Coord: parcePoint(str),
		}
	}

	return places
}

// 文字列配列から道をパースする
func parceRoad(plane []string, places []Place) []Road {
	roads := make([]Road, len(plane))

	for i, str := range plane {
		from, to, _, _ := getNums(str)
		roads[i] = Road{
			Id:   i + 1,
			From: places[from-1],
			To:   places[to-1],
		}
	}

	return roads
}

// 文字列配列からクエリをパースする
func parceQuery(plane []string) []Query {
	queries := make([]Query, len(plane))

	for i, str := range plane {
		query := strings.Split(str, " ")
		num, err := strconv.Atoi(query[2])
		if err != nil {
			num = 0
		}
		queries[i] = Query{
			Start: query[0],
			Dest:  query[1],
			Num:   num,
		}
	}

	return queries

}
