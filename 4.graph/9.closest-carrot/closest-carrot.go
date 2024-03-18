// Write a function, closestCarrot, that takes in a grid, a starting row, and a starting column. In the grid, 'X's are walls, 'O's are open spaces, and 'C's are carrots. The function should return a number representing the length of the shortest path from the starting position to a carrot. You may move up, down, left, or right, but cannot pass through walls (X). If there is no possible path to a carrot, then return -1.

package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	row      int
	col      int
	distance int
}

func closestCarrot(grid [][]rune, startRow, startCol int) int {
	queue := list.New()
	queue.PushBack(&Queue{startRow, startCol, 0})
	visited := make(map[string]bool)
	visited[fmt.Sprintf("%d,%d", startRow, startCol)] = true

	for queue.Len() > 0 {
		cur := queue.Remove(queue.Front()).(*Queue)
		row := cur.row
		col := cur.col
		distance := cur.distance
		if grid[row][col] == 'C' {
			return distance
		}

		deltas := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		for _, delta := range deltas {
			deltaRow := delta[0]
			deltaCol := delta[1]
			neighborRow := row + deltaRow
			neighborCol := col + deltaCol
			if isValidMove(grid, neighborRow, neighborCol, visited) {
				visited[fmt.Sprintf("%d,%d", neighborRow, neighborCol)] = true
				queue.PushBack(&Queue{neighborRow, neighborCol, distance + 1})
			}

		}
	}

	return -1
}

func isValidMove(grid [][]rune, row, col int, visited map[string]bool) bool {
	rowInbounds := 0 <= row && row < len(grid)
	colInbounds := 0 <= col && col < len(grid[0])
	_, ok := visited[fmt.Sprintf("%d,%d", row, col)]
	return rowInbounds && colInbounds && grid[row][col] != 'X' && !ok
}

func main() {
	grid := [][]rune{
		{'O', 'O', 'O', 'O', 'O'},
		{'O', 'X', 'O', 'O', 'O'},
		{'O', 'X', 'X', 'O', 'O'},
		{'O', 'X', 'C', 'O', 'O'},
		{'O', 'X', 'X', 'O', 'O'},
		{'C', 'O', 'O', 'O', 'O'},
	}

	fmt.Println(closestCarrot(grid, 1, 2))
}
