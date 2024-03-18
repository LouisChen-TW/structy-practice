// Write a function, connectedComponentsCount, that takes in the adjacency list of an undirected graph. The function should return the number of connected components within the graph.

package main

import "fmt"

type Graph struct {
	Nodes map[int][]int
}

func main() {
	graph := Graph{
		Nodes: map[int][]int{
			0: {8, 1, 5},
			1: {0},
			5: {0, 8},
			8: {0, 5},
			2: {3, 4},
			3: {2, 4},
			4: {3, 2},
		},
	}
	fmt.Println(connectedComponentsCount(graph))
}

// iterative
// func connectedComponentsCount(graph Graph) int {
// 	count := 0
// 	visited := make(map[int]bool)

// 	for i := range graph.Nodes {
// 		if _, ok := visited[i]; ok {
// 			continue
// 		}
// 		stack := []int{i}
// 		for len(stack) > 0 {
// 			cur := stack[len(stack)-1]
// 			stack = stack[:len(stack)-1]
// 			if _, ok := visited[cur]; ok {
// 				continue
// 			}
// 			visited[cur] = true
// 			for _, node := range graph.Nodes[cur] {
// 				stack = append(stack, node)
// 			}
// 		}
// 		count++
// 	}
// 	return count
// }

// recursive
func connectedComponentsCount(graph Graph) int {
	count := 0
	visited := make(map[int]bool)

	for i := range graph.Nodes {
		if explore(graph, i, &visited) == true {
			count++
		}
	}
	return count
}

func explore(graph Graph, current int, visited *map[int]bool) bool {
	if _, ok := (*visited)[current]; ok {
		return false
	}
	(*visited)[current] = true
	for _, neighbor := range graph.Nodes[current] {
		explore(graph, neighbor, visited)
	}
	return true
}
