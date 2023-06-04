package number_of_islands

func NumIslands(grid [][]byte) int {
	ans := 0
	visited := make([][]bool, len(grid))
	for i := range grid {
		visited[i] = make([]bool, len(grid[i]))
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' && !visited[i][j] {
				ans++
				dfs(i, j, grid, visited)
			}
		}
	}
	return ans
}

func dfs(i, j int, grid [][]byte, visited [][]bool) {
	if i > len(grid)-1 || i < 0 || j < 0 || j > len(grid[i])-1 || visited[i][j] || grid[i][j] == '0' {
		return
	}
	visited[i][j] = true
	dfs(i+1, j, grid, visited)
	dfs(i, j+1, grid, visited)
	dfs(i-1, j, grid, visited)
	dfs(i, j-1, grid, visited)
}
