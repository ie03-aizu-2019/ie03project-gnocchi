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
	N, M, _, _ := getNums(strs[0])

	data.Places = parcePlace(strs[1 : N+1])
	data.Roads = parceRoad(strs[N+1:M+N+1], data.Places)

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
func parcePlace(plane []string) []Place {
	places := make([]Place, len(plane))

	for i, str := range plane {
		places[i] = Place{
			Id:    i + 1,
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
