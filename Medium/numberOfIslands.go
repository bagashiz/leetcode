package main

// Link: https://leetcode.com/problems/number-of-islands/

func numIslands(grid [][]byte) int {
	// if the length of the grid is 0, return 0
	if len(grid) == 0 {
		return 0
	}

	// create a variable to store the number of islands
	var count int

	// loop through the grid
	for i := 0; i < len(grid); i++ {
		// loop through the second dimension of the grid
		for j := 0; j < len(grid[0]); j++ {
			// if the current element is 1, then we have an land
			if grid[i][j] == '1' {
				// increment the count
				count++

				// call the helper function
				helper(grid, i, j)
			}
		}
	}

	return count
}

// helper function to validate if the current element is an island (surrounded by water)
func helper(grid [][]byte, i int, j int) {
	// if the current element does not meet the criteria of being an island, return nothing
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '0' {
		return
	}

	// set the current element to 0
	grid[i][j] = '0'

	// call the helper function for the current element's neighbors
	helper(grid, i+1, j)
	helper(grid, i-1, j)
	helper(grid, i, j+1)
	helper(grid, i, j-1)
}

func main() {
	// test cases
	println(numIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'}})) // 1

	println(numIslands([][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'}})) // 3
}
