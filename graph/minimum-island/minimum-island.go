// Write a function, minimumIsland, that takes in a grid containing Ws and Ls. W represents water and L represents land. The function should return the size of the smallest island. An island is a vertically or horizontally connected region of land.

// You may assume that the grid contains at least one island.

package main

import (
	"fmt"
	"math"
)

type Graph struct {
	Nodes map[string][]string
}

func islandCount(grid [][]rune) int {
	minIsland := math.MaxInt
	visited := make(map[string]bool)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 'W' {
				continue
			}
			if visited[fmt.Sprintf("%d,%d", i, j)] {
				continue
			}
			islandCount := exploreIsland(grid, i, j, visited)
			if islandCount < minIsland {
				minIsland = islandCount
			}
		}
	}
	return minIsland
}

func exploreIsland(grid [][]rune, i, j int, visited map[string]bool) int {
	rowInbounds := 0 <= i && i < len(grid)
	colInbounds := 0 <= j && j < len(grid[0])
	if !rowInbounds || !colInbounds {
		return 0
	}
	if visited[fmt.Sprintf("%d,%d", i, j)] {
		return 0
	}
	if grid[i][j] == 'W' {
		return 0
	}
	if grid[i][j] == 'L' {
		visited[fmt.Sprintf("%d,%d", i, j)] = true
	}
	return 1 +
		exploreIsland(grid, i+1, j, visited) +
		exploreIsland(grid, i-1, j, visited) +
		exploreIsland(grid, i, j+1, visited) +
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
