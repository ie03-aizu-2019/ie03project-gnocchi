package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/uzimaru0000/ie03project-gnocchi/back/model"
)

type Datas struct {
	Places    []*model.Place
	Roads     []*model.Road
	AddPlaces []*model.Place
	Queries   []*model.Query
}

// Load file to return content
func Load(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// ParseData is string to datas
func ParseData(str string) (*Datas, error) {
	data := &Datas{}
	strs := strings.Split(str, "\n")
	fN, fM, fP, fQ := getNums(strs[0])

	N := int(fN)
	M := int(fM)
	P := int(fP)
	Q := int(fQ)

	data.Places = parcePlace(strs[1:N+1], 0)
	data.Roads = parceRoad(strs[N+1:M+N+1], data.Places)

	if P != 0 {
		data.AddPlaces = parcePlace(strs[M+N+1:P+M+N+1], N)
	}

	if Q != 0 {
		data.Queries = parceQuery(strs[P+M+N+1:])
	}

	return data, nil
}

// 空白区切りで文字列をパースする
func getNums(plane string) (float64, float64, float64, float64) {
	nums := make([]float64, 4)

	splited := strings.Split(plane, " ")
	for i, str := range splited {
		if n, err := strconv.ParseFloat(str, 64); err != nil {
			nums[i] = 0
		} else {
			nums[i] = n
		}
	}

	return nums[0], nums[1], nums[2], nums[3]
}

// 文字列から座標をパースする
func parcePoint(str string) model.Point {
	x, y, _, _ := getNums(str)
	return model.Point{X: float64(x), Y: float64(y)}
}

// 文字列配列から地点をパースする
// plane : 入力された文字列
// n		 : idの始まり
func parcePlace(plane []string, n int) []*model.Place {
	places := make([]*model.Place, len(plane))

	for i, str := range plane {
		places[i] = &model.Place{
			Id:    strconv.Itoa(n + i + 1),
			Coord: parcePoint(str),
		}
	}

	return places
}

// 文字列配列から道をパースする
func parceRoad(plane []string, places []*model.Place) []*model.Road {
	roads := make([]*model.Road, len(plane))

	for i, str := range plane {
		f, t, _, _ := getNums(str)

		from := int(f)
		to := int(t)

		roads[i] = &model.Road{
			Id:   i + 1,
			From: places[from-1],
			To:   places[to-1],
		}
	}

	return roads
}

// 文字列配列からクエリをパースする
func parceQuery(plane []string) []*model.Query {
	queries := make([]*model.Query, len(plane))

	for i, str := range plane {
		query := strings.Split(str, " ")
		num, err := strconv.Atoi(query[2])
		if err != nil {
			num = 0
		}
		queries[i] = &model.Query{
			Start: query[0],
			Dest:  query[1],
			Num:   num,
		}
	}

	return queries

}

func DatasToQuerys(datas Datas) string {
	placeMap := createPlaceMap(datas.Places)

	placeNum := len(datas.Places)
	roadNum := len(datas.Roads)
	addPlaceNum := len(datas.AddPlaces)
	queryNum := len(datas.Queries)

	places := placesToQuery(datas.Places)
	roads := roadsToQuery(datas.Roads, placeMap)
	addPlaces := placesToQuery(datas.AddPlaces)
	queries := queriesToQuery(datas.Queries)

	template := `%d %d %d %d
%s%s%s%s`

	return fmt.Sprintf(
		template,
		placeNum,
		roadNum,
		addPlaceNum,
		queryNum,
		places,
		roads,
		addPlaces,
		queries,
	)
}

func createPlaceMap(places []*model.Place) map[string]int {
	placeMap := make(map[string]int)

	for i, p := range places {
		placeMap[p.Id] = i + 1
	}

	return placeMap
}

func placesToQuery(places []*model.Place) string {
	var result string

	for _, p := range places {
		result += fmt.Sprintln(p.Coord.ToString())
	}

	return result
}

func roadsToQuery(roads []*model.Road, places map[string]int) string {
	var result string

	for _, r := range roads {
		result += fmt.Sprintf("%d %d\n", places[r.From.Id], places[r.To.Id])
	}

	return result
}

func queriesToQuery(queries []*model.Query) string {
	var result string

	for _, q := range queries {
		result += fmt.Sprintln(q.ToString())
	}

	return result
}
