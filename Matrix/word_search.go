package main

import "fmt"

func dfs(board [][]byte, i, j int, word string) bool {
	if len(word) == 0 {
		return true
	}

	m := len(board)
	n := len(board[0])

	x := []int{-1, 0, 1, 0}
	y := []int{0, -1, 0, 1}

	if i < 0 || i > m || j < 0 || j > n || board[i][j] != word[0] {
		return false
	}

	c := board[i][j]
	board[i][j] = '*'
	word = word[1:]
	res := false
	for k := 0; k < 4; k++ {
		res = res || dfs(board, i+x[k], j+y[k], word)
	}

	board[i][j] = c
	return res
}

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(board, i, j, word) {
				return true
			}
		}
	}
	return false
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"

	fmt.Println(exist(board, word)) // Output: true
}
