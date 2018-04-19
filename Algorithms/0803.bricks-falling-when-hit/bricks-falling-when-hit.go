package problem0803

func hitBricks(grid [][]int, hits [][]int) []int {
	res := make([]int, len(hits))

	for idx, hit := range hits {
		i, j := hit[0], hit[1]
		grid[i][j] = 0

		hanging := idx + 2
		for k := 0; k < len(grid[0]); k++ {
			if grid[0][k] == hanging-1 {
				update(hanging, 0, k, grid)
			}
		}

		for k := 0; k < 4; k++ {
			x := i + dx[k]
			y := j + dy[k]
			res[idx] += drop(hanging, x, y, grid)
		}
	}

	return res
}

func update(hanging, i, j int, grid [][]int) {
	grid[i][j] = hanging
	m, n := len(grid), len(grid[0])
	for k := 0; k < 4; k++ {
		x := i + dx[k]
		y := j + dy[k]
		if 0 <= x && x < m && 0 <= y && y < n && grid[x][y] == hanging-1 {
			update(hanging, x, y, grid)
		}
	}
}

func drop(hanging, i, j int, grid [][]int) int {
	m, n := len(grid), len(grid[0])

	if i < 0 || m <= i ||
		j < 0 || n <= j ||
		grid[i][j] == hanging ||
		grid[i][j] == 0 {
		return 0
	}

	grid[i][j] = 0
	res := 1

	for k := 0; k < 4; k++ {
		x := i + dx[k]
		y := j + dy[k]
		res += drop(hanging, x, y, grid)
	}

	return res
}

var dx = []int{-1, 0, 0, 1}
var dy = []int{0, -1, 1, 0}