// Write a function, minimumIsland, that takes in a grid containing Ws and Ls. W represents water and L represents land. The function should return the size of the smallest island. An island is a vertically or horizontally connected region of land.

// You may assume that the grid contains at least one island.

const minimumIsland = grid => {
    let minIsland = Infinity;
    const visited = new Set();

    for (let i = 0; i < grid.length; i++) {
        for (let j = 0; j < grid[0].length; j++) {
            if (grid[i][j] === 'W') continue;
            if (visited.has(`${i},${j}`)) continue;
            const isLandCount = exploreIsland(grid, i, j, visited);
            if (isLandCount < minIsland) minIsland = isLandCount;
        }
    }
    return minIsland;
};

const exploreIsland = (grid, i, j, visited) => {
    const rowInbounds = 0 <= i && i < grid.length;
    const colInbounds = 0 <= j && j < grid[0].length;
    if (!rowInbounds || !colInbounds) return 0;
    if (visited.has(`${i},${j}`)) return 0;
    if (grid[i][j] === 'W') return 0;
    if (grid[i][j] === 'L') visited.add(`${i},${j}`);
    const right = exploreIsland(grid, i + 1, j, visited);
    const left = exploreIsland(grid, i - 1, j, visited);
    const down = exploreIsland(grid, i, j + 1, visited);
    const up = exploreIsland(grid, i, j - 1, visited);
    return 1 + right + left + down + up;
};

const grid = [
    ['W', 'L', 'W', 'W', 'W'],
    ['W', 'L', 'W', 'W', 'W'],
    ['W', 'W', 'W', 'L', 'W'],
    ['W', 'W', 'L', 'L', 'W'],
    ['L', 'W', 'W', 'L', 'L'],
    ['L', 'L', 'W', 'W', 'W'],
];

console.log(minimumIsland(grid)); // -> 2
