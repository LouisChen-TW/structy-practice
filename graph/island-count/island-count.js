// Write a function, islandCount, that takes in a grid containing Ws and Ls. W represents water and L represents land. The function should return the number of islands on the grid. An island is a vertically or horizontally connected region of land.

const islandCount = grid => {
    let count = 0
    const visited = new Set()

    for (let i = 0; i < grid.length; i++) {
        for (let j = 0; j < grid[0].length; j++) {
            if (grid[i][j] === 'W') continue;
            if (visited.has(`${i},${j}`)) continue;
            exploreIsland(grid, i, j, visited);
            count++
        }
    }
    return count
};

const exploreIsland = (grid, i, j, visited) => {
    const rowInbounds = 0 <= i && i < grid.length;
    const colInbounds = 0 <= j && j < grid[0].length;
    if (!rowInbounds || !colInbounds) return;
    if (visited.has(`${i},${j}`)) return;
    if (grid[i][j] === 'W') return;
    if (grid[i][j] === 'L') visited.add(`${i},${j}`);
    exploreIsland(grid, i + 1, j, visited);
    exploreIsland(grid, i - 1, j, visited);
    exploreIsland(grid, i, j + 1, visited);
    exploreIsland(grid, i, j - 1, visited);
}

const grid = [
  ['W', 'L', 'W', 'W', 'W'],
  ['W', 'L', 'W', 'W', 'W'],
  ['W', 'W', 'W', 'L', 'W'],
  ['W', 'W', 'L', 'L', 'W'],
  ['L', 'W', 'W', 'L', 'L'],
  ['L', 'L', 'W', 'W', 'W'],
];

console.log(islandCount(grid)); // -> 3