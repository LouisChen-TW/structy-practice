// Write a function, undirectedPath, that takes in an array of edges for an undirected graph and two nodes (nodeA, nodeB). The function should return a boolean indicating whether or not there exists a path between nodeA and nodeB.

package main

import "fmt"

type Graph struct {
	Nodes map[string][]string
}

func undirectedPath(edges [][]string, nodeA, nodeB string) bool {
	graph := edgesToGraph(edges)
	newSet := make(map[string]struct{})
	stack := []string{nodeA}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		if _, ok := newSet[cur]; ok {
			stack = stack[:len(stack)-1]
			continue
		}
		if cur == nodeB {
			return true
		}
		newSet[cur] = struct{}{}
		stack = stack[:len(stack)-1]
		for _, node := range (*graph).Nodes[cur] {
			stack = append(stack, node)
		}
	}
	return false
}

func edgesToGraph(edges [][]string) *Graph {
	graph := &Graph{
		Nodes: make(map[string][]string),
	}
	for _, edge := range edges {
		for i, node := range edge {
			nodes := graph.Nodes[node]
			if i == 0 {
				nodes = append(nodes, edge[1])
			} else {
				nodes = append(nodes, edge[0])
			}
			graph.Nodes[node] = nodes
		}
	}
	return graph
}

func main() {
	edges := [][]string{
		{"i", "j"},
		{"k", "i"},
		{"m", "k"},
		{"k", "l"},
		{"o", "n"},
	}

	fmt.Println(undirectedPath(edges, "j", "m"))
}
