package phase2

import (
	"log"

	"github.com/uzimaru0000/ie03project-gnocchi/back/model"
)

type node struct {
	place *model.Place
	pre   int
	low   int
}

type connection struct {
	node    *node
	connect []*node
}

type graph map[string]*connection

func createGraph(roads []*model.Road) graph {
	graph := make(map[string]*connection)

	for _, r := range roads {
		_, ok := graph[r.From.Id]
		if !ok {
			n := &node{
				place: r.From,
				pre:   0,
				low:   0,
			}
			graph[r.From.Id] = &connection{
				node:    n,
				connect: []*node{},
			}
		}
		graph[r.From.Id].connect = append(graph[r.From.Id].connect,
			&node{
				place: r.To,
				pre:   0,
				low:   0,
			},
		)

		_, ok = graph[r.To.Id]
		if !ok {
			n := &node{
				place: r.To,
				pre:   0,
				low:   0,
			}
			graph[r.To.Id] = &connection{
				node:    n,
				connect: []*node{},
			}
		}
		graph[r.To.Id].connect = append(graph[r.To.Id].connect,
			&node{
				place: r.From,
				pre:   0,
				low:   0,
			},
		)
	}

	return graph
}

func dfs(current *node, pre *node, g graph) graph {
	if pre != nil {
		log.Printf("%v", *pre)
		current.pre = pre.pre + 1
		current.low = pre.pre + 1
	}

	newGraph := g
	for _, conn := range newGraph[current.place.Id].connect {
		if conn.pre == 0 {
			newGraph = dfs(conn, current, newGraph)
		}
	}

	return newGraph
}
