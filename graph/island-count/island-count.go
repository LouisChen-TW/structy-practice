// Write a function, islandCount, that takes in a grid containing Ws and Ls. W represents water and L represents land. The function should return the number of islands on the grid. An island is a vertically or horizontally connected region of land.
package main

import "fmt"

type Graph struct {
	Nodes map[string][]string
}

func islandCount(grid [][]rune) int {
	count := 0
	visited := make(map[string]bool)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 'W' {
				continue
			}
			if visited[fmt.Sprintf("%d,%d", i, j)] {
				continue
			}
			exploreIsland(grid, i, j, visited)
			count++
		}
	}
	return count
}

func exploreIsland(grid [][]rune, i, j int, visited map[string]bool) {
	rowInbounds := 0 <= i && i < len(grid)
	colInbounds := 0 <= j && j < len(grid[0])
	if !rowInbounds || !colInbounds {
		return
	}
	if visited[fmt.Sprintf("%d,%d", i, j)] {
		return
	}
	if grid[i][j] == 'W' {
		return
	}
	if grid[i][j] == 'L' {
		visited[fmt.Sprintf("%d,%d", i, j)] = true
	}
	exploreIsland(grid, i+1, j, visited)
	exploreIsland(grid, i-1, j, visited)
	exploreIsland(grid, i, j+1, visited)
	exploreIsland(grid, i, j-1, visited)
}

func main() {
	grid := [][]rune{
		{'W', 'L', 'W', 'W', 'W'},
		{'W', 'L', 'W', 'W', 'W'},
		{'W', 'W', 'W', 'L', 'W'},
		{'W', 'W', 'L', 'L', 'W'},
		{'L', 'W', 'W', 'L', 'L'},
		{'L', 'L', 'W', 'W', 'W'},
	}

	fmt.Println(islandCount(grid))
}

// In Go, a rune is an alias for int32 and represents a Unicode code point. It's commonly used to represent individual characters in strings. When dealing with single characters or small strings, using rune is efficient and idiomatic in Go.
