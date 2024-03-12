// basic graph traverse, use depth and breadth.

package main

import "fmt"

type Graph struct {
	Nodes map[string][]string
}

func main() {
	graph := Graph{
		Nodes: map[string][]string{
			"a": {"b", "c"},
			"b": {"d"},
			"c": {"e"},
			"d": {"f"},
			"e": {},
			"f": {},
		},
	}
	// depthFirstPrint(&graph, "a")
	breadthFirstPrint(&graph, "a")
}

// iterative depth
// func depthFirstPrint(graph *Graph, source string) {
// 	stack := []string{source}
// 	for len(stack) > 0 {
// 		cur := stack[len(stack)-1]
// 		fmt.Print(cur)
// 		stack = stack[:len(stack)-1]

// 		if _, ok := graph.Nodes[cur]; ok {
// 			for _, neighbor := range graph.Nodes[cur] {
// 				stack = append(stack, neighbor)
// 			}
// 		}
// 	}
// }

// recursive depth
func depthFirstPrint(graph *Graph, source string) {
	fmt.Print(source)
	if _, ok := graph.Nodes[source]; ok {
		for _, neighbor := range graph.Nodes[source] {
			depthFirstPrint(graph, neighbor)
		}
	}
}

func breadthFirstPrint(graph *Graph, source string) {
	queue := []string{source}
	for len(queue) > 0 {
		cur := queue[0]
		fmt.Print(cur)
		queue = queue[1:]

		if _, ok := graph.Nodes[cur]; ok {
			for _, neighbor := range graph.Nodes[cur] {
				queue = append(queue, neighbor)
			}
		}
	}
}
