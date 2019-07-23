package phase2

import (
	"testing"

	"github.com/uzimaru0000/ie03project-gnocchi/back/utils"
)

func TestCreateGraph(t *testing.T) {
	str := `6 6 4 0
0 0
2 5
4 7
5 5
7 1
9 5
1 4
1 6
2 5
3 5
4 6
4 2
5 1
11 5
5 4
3 6`

	data, err := utils.ParseData(str)
	if err != nil {
		t.Fatal("fomat error")
	}

	graph := createGraph(data.Roads)
	for _, node := range graph {
		t.Logf("%v ->", node.place.Id)
		for _, edge := range node.conn {
			t.Logf("\t%s", edge.dest.place.Id)
		}
	}
}

func TestDFS(t *testing.T) {
	str := `6 6 4 0
0 0
2 5
4 7
5 5
7 1
9 5
1 4
1 6
2 5
3 5
4 6
4 2
5 1
11 5
5 4
3 6`

	data, err := utils.ParseData(str)
	if err != nil {
		t.Fatal("fomat error")
	}

	graph := createGraph(data.Roads)

	dfs(graph[0], 0)
	for _, node := range graph {
		t.Logf("%v -> pre: %d, low: %d", node.place.Id, node.pre, node.low)
	}
}
