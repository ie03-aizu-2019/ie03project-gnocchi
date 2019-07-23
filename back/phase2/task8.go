package phase2

import (
	"math"

	"github.com/uzimaru0000/ie03project-gnocchi/back/model"
)

type node struct {
	place *model.Place
	pre   int
	low   int
	conn  []*edge
}

type edge struct {
	length float64
	dest   *node
}

type graph []*node

func createGraph(roads []*model.Road) graph {
	places, table := createTable(createMap(roads))

	nodeTable := make(map[string]int)
	graph := make(graph, len(places))
	i := 0
	for id, place := range places {
		graph[i] = createNode(place)
		nodeTable[id] = i
		i++
	}

	for id, i := range nodeTable {
		for _, conn := range table[id] {
			node := graph[nodeTable[conn]]
			graph[i].conn = append(graph[i].conn, createEdge(places[id], places[conn], node))
		}
	}

	return graph
}

func createMap(roads []*model.Road) map[*model.Place][]*model.Place {
	matrix := make(map[*model.Place][]*model.Place)

	for _, road := range roads {
		conn, ok := matrix[road.From]
		if !ok {
			matrix[road.From] = []*model.Place{road.To}
		} else {
			matrix[road.From] = append(conn, road.To)
		}

		conn, ok = matrix[road.To]
		if !ok {
			matrix[road.To] = []*model.Place{road.From}
		} else {
			matrix[road.To] = append(conn, road.From)
		}
	}

	return matrix
}

func createTable(placeMap map[*model.Place][]*model.Place) (map[string]*model.Place, map[string][]string) {
	places := make(map[string]*model.Place)
	table := make(map[string][]string)

	for key, val := range placeMap {
		places[key.Id] = key
		table[key.Id] = make([]string, len(val))

		for i, p := range val {
			table[key.Id][i] = p.Id
		}
	}

	return places, table
}

func createNode(place *model.Place) *node {
	return &node{
		place: place,
		pre:   -1,
		low:   -1,
		conn:  []*edge{},
	}
}

func createEdge(from, to *model.Place, dest *node) *edge {
	return &edge{
		length: (&model.Road{
			Id:   0,
			From: from,
			To:   to,
		}).Length(),
		dest: dest,
	}
}

func dfs(current *node, preOrder int) int {
	current.pre = preOrder
	current.low = preOrder

	for _, edge := range current.conn {
		if edge.dest.pre == -1 {
			preOrder = dfs(edge.dest, preOrder+1)
		} else {
			current.low = int(math.Min(float64(edge.dest.low), float64(current.low)))
		}
	}

	return preOrder
}
