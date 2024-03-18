// Write a function, hasPath, that takes in an object representing the adjacency list of a directed acyclic graph and two nodes (src, dst). The function should return a boolean indicating whether or not there exists a directed path between the source and destination nodes.

// Hey. This is our first graph problem, so you should be liberal with watching the Approach and Walkthrough. Be productive, not stubborn. -AZ

package main

import "fmt"

type Graph struct {
	Nodes map[string][]string
}

func main() {
	graph := Graph{
		Nodes: map[string][]string{
			"f": {"g", "i"},
			"g": {"h"},
			"h": {},
			"i": {"g", "k"},
			"j": {"i"},
			"k": {},
		},
	}
	fmt.Println(hasPath(&graph, "f", "k"))
}

// iterative
func hasPath(graph *Graph, src, dst string) bool {
	stack := []string{src}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		if cur == dst {
			return true
		}
		stack = stack[:len(stack)-1]
		for _, neighbor := range graph.Nodes[cur] {
			stack = append(stack, neighbor)
		}
	}
	return false
}

// recursive
// func hasPath(graph *Graph, src, dst string) bool {
// 	if src == dst {
// 		return true
// 	}
// 	for _, neighbor := range graph.Nodes[src] {
// 		if hasPath(graph, neighbor, dst) == true {
// 			return true
// 		}
// 	}
// 	return false
// }
