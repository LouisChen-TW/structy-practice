// Write a function, largestComponent, that takes in the adjacency list of an undirected graph. The function should return the size of the largest connected component in the graph.

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
	visited := make(map[int]bool)
	max := 0

	for i := range graph.Nodes {
		componentLength := explore(graph, i, &visited)
		if componentLength > max {
			max = componentLength
		}
	}
	return max
}

func explore(graph Graph, current int, visited *map[int]bool) int {
	if _, ok := (*visited)[current]; ok {
		return 0
	}
	(*visited)[current] = true
	size := 1
	for _, neighbor := range graph.Nodes[current] {
		size += explore(graph, neighbor, visited)
	}
	return size
}
