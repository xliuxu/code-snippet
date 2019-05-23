/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	visit := make([][]bool, len(grid))
	for v := range visit {
		visit[v] = make([]bool, len(grid[0]))
	}
	var res int
	trow := len(grid)
	tcol := len(grid[0])
	for i := 0; i < trow; i++ {
		for j := 0; j < tcol; j++ {
			if !visit[i][j] && grid[i][j] == '1' {
				res++
				dfs(grid, visit, i, j, trow, tcol)
			} else {
				visit[i][j] = true
			}
		}
	}
	return res
}

func dfs(grid [][]byte, visit [][]bool, row, col, trow, tcol int) {
	if row >= trow || col >= tcol || row < 0 || col < 0 {
		return
	}
	if visit[row][col] {
		return
	}
	visit[row][col] = true
	if grid[row][col] == '0' {
		return
	}
	dfs(grid, visit, row+1, col, trow, tcol)
	dfs(grid, visit, row, col+1, trow, tcol)
	dfs(grid, visit, row, col-1, trow, tcol)
	dfs(grid, visit, row-1, col, trow, tcol)

}
