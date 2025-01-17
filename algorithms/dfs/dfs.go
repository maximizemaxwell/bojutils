// dfs.go
package dfs

func DFS(grid [][]int, visited [][]bool, x, y int) {
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	n, m := len(grid), len(grid[0])

	visited[x][y] = true

	for _, dir := range directions {
		nx, ny := x+dir[0], y+dir[1]
		if nx >= 0 && nx < n && ny >= 0 && ny < m && grid[nx][ny] == 1 && !visited[nx][ny] {
			DFS(grid, visited, nx, ny)
		}
	}
}
