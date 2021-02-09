package igraphs

import (
	"container/heap"
	"fmt"
)

type edge struct {
	dstNode int
	edgeWeight int
}

type igraph struct {
	nodes map[int][]edge
}

func (g *igraph) addEdge(src, dst int, weight int) {
	g.nodes[src] = append(g.nodes[src], edge{dst, weight})
	g.nodes[dst] = append(g.nodes[dst], edge{src, weight})
}

func (g *igraph) getEdges(node int) []edge {
	return g.nodes[node]
}

// Dijkstra description
func (g *igraph) Dijkstra(src, dst int) int {
	h := &pathHeap{}
	heap.Init(h)
	heap.Push(h, path{0, []int{src}})

	isVisited := make(map[int]bool)

	for h.Len() > 0 {
		getMinPath := heap.Pop(h).(path)
		lastNode := getMinPath.nodes[len(getMinPath.nodes)-1]

		if isVisited[lastNode] {
			continue
		}

		if lastNode == dst {
			return getMinPath.val
		}

		for _, e :=range g.getEdges(lastNode) {
			if !isVisited[e.dstNode] {
				heap.Push(h, path{getMinPath.val + e.edgeWeight, append(getMinPath.nodes, e.dstNode)})
			}
		}
		isVisited[lastNode] = true
	}

	return 0
}


// IGraphs description
func IGraphs() {
	fmt.Println("Граф из wiki: https://ru.wikipedia.org/wiki/Алгоритм_Дейкстры")
	fmt.Println("Результат работы алгоритма виден на последнем рисунке: \n" +
		"кратчайший путь от вершины 1\n" +
		"\tдо 2-й - 7, \n" +
		"\tдо 3-й — 9, \n" +
		"\tдо 4-й — 20, \n" +
		"\tдо 5-й — 20, \n" +
		"\tдо 6-й — 11.")
	g := &igraph{make(map[int][]edge)}
	g.addEdge(1,2,7);g.addEdge(1,3,9);g.addEdge(1,6,14)
	g.addEdge(2,3,10);g.addEdge(2,4,15);g.addEdge(3,4,11)
	g.addEdge(3,6,2);g.addEdge(4,5,6);g.addEdge(5,6,9)
	fmt.Println("Результат программы:")
	fmt.Println(g.Dijkstra(1,2));fmt.Println(g.Dijkstra(1,3));fmt.Println(g.Dijkstra(1,4))
	fmt.Println(g.Dijkstra(1,5));
	fmt.Println(g.Dijkstra(1,6))
}


