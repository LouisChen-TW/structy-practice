// Write a function, closestCarrot, that takes in a grid, a starting row, and a starting column. In the grid, 'X's are walls, 'O's are open spaces, and 'C's are carrots. The function should return a number representing the length of the shortest path from the starting position to a carrot. You may move up, down, left, or right, but cannot pass through walls (X). If there is no possible path to a carrot, then return -1.

const closestCarrot = (grid, startRow, startCol) => {
    const queue = [{ row: startRow, col: startCol, distance: 0 }];
    const visited = new Set([`${startRow},${startCol}`]);
    while (queue.length > 0) {
        const { row, col, distance } = queue.shift();
        if (grid[row][col] === 'C') return distance;

        const deltas = [
            [1, 0],
            [-1, 0],
            [0, 1],
            [0, -1],
        ];

        for (let delta of deltas) {
            const [deltaRow, deltaCol] = delta;
            const neighborRow = row + deltaRow;
            const neighborCol = col + deltaCol;
            if (isValidMove(grid, neighborRow, neighborCol, visited)) {
                visited.add(`${neighborRow},${neighborCol}`);
                queue.push({ row: neighborRow, col: neighborCol, distance: distance + 1 });
            }
        }
    }
    return -1;
};

const isValidMove = (grid, row, col, visited) => {
    const rowInbounds = 0 <= row && row < grid.length;
    const colInbounds = 0 <= col && col < grid[0].length;
    return rowInbounds && colInbounds && grid[row][col] !== 'X' && !visited.has(`${row},${col}`);
};

const grid = [
    ['O', 'O', 'O', 'O', 'O'],
    ['O', 'X', 'O', 'O', 'O'],
    ['O', 'X', 'X', 'O', 'O'],
    ['O', 'X', 'C', 'O', 'O'],
    ['O', 'X', 'X', 'O', 'O'],
    ['C', 'O', 'O', 'O', 'O'],
];

console.log(closestCarrot(grid, 1, 2)); // -> 4
