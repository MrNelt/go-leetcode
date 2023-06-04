package max_area_of_island

func maxAreaOfIsland(grid [][]int) int {
	ans := 0
	visited := make([][]bool, len(grid))
	for i := range grid {
		visited[i] = make([]bool, len(grid[i]))
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 && !visited[i][j] {
				ans = max(ans, dfs(i, j, grid, visited))
			}
		}
	}
	return ans
}

func dfs(i, j int, grid [][]int, visited [][]bool) int {
	if i > len(grid)-1 || i < 0 || j < 0 || j > len(grid[i])-1 || visited[i][j] || grid[i][j] == 0 {
		return 0
	}
	visited[i][j] = true
	return 1 + dfs(i+1, j, grid, visited) + dfs(i, j+1, grid, visited) + dfs(i-1, j, grid, visited) + dfs(i, j-1, grid, visited)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
