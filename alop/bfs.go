package main

import (
	"fmt"
)

// Graph represents an undirected graph using an adjacency list representation.
type Graph struct {
	adjList map[int][]int
}

// NewGraph creates a new instance of a Graph.
func NewGraph() *Graph {
	return &Graph{adjList: make(map[int][]int)}
}

// AddEdge adds an edge to the graph.
func (g *Graph) AddEdge(u, v int) {
	g.adjList[u] = append(g.adjList[u], v)
	g.adjList[v] = append(g.adjList[v], u)
}

// BFS performs a breadth-first search on the graph starting from the given start vertex.
func (g *Graph) BFS(start int) {
	visited := make(map[int]bool)
	queue := []int{start}

	for len(queue) > 0 {
		// Dequeue a vertex from queue
		vertex := queue[0]
		queue = queue[1:]

		// If not visited, visit the vertex
		if !visited[vertex] {
			visited[vertex] = true
			fmt.Println(vertex)

			// Add all unvisited neighbors to the queue
			for _, neighbor := range g.adjList[vertex] {
				if !visited[neighbor] {
					queue = append(queue, neighbor)
				}
			}
		}
	}
}

func main2() {
	// Create a new graph instance
	graph := NewGraph()

	// Add some edges to the graph
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 0)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 3)

	// Perform BFS starting from vertex 2
	fmt.Println("Starting BFS from vertex 2:")
	graph.BFS(1)
}
