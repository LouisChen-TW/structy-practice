// Write a function, shortestPath, that takes in an array of edges for an undirected graph and two nodes (nodeA, nodeB). The function should return the length of the shortest path between A and B. Consider the length as the number of edges in the path, not the number of nodes. If there is no path between A and B, then return -1.

package main

import "fmt"

type Graph struct {
	Nodes map[string][]string
}

func buildGraph(edges [][]string) *Graph {
	graph := &Graph{
		Nodes: make(map[string][]string),
	}
	for _, edge := range edges {
		source := edge[0]
		destination := edge[1]
		graph.Nodes[source] = append(graph.Nodes[source], destination)
		graph.Nodes[destination] = append(graph.Nodes[destination], source)
	}
	return graph
}

func shortestPath(edges [][]string, nodeA, nodeB string) int {
	visited := make(map[string]int)
	graph := buildGraph(edges)
	queue := []string{nodeA}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur == nodeB {
			return visited[cur]
		}
		for _, node := range graph.Nodes[cur] {
			if _, ok := visited[node]; !ok {
				visited[node] = visited[cur] + 1
				queue = append(queue, node)
			}
		}
	}
	return -1
}

func main() {
	edges := [][]string{
		{"w", "c"},
		{"x", "y"},
		{"z", "y"},
		{"z", "v"},
		{"w", "v"},
	}
	fmt.Println(shortestPath(edges, "w", "z"))
}
