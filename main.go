package main

import "fmt"

// SolveNQueens returns all placements of n queens on an n×n board with no
// two attacking. Boards are given as strings, 'Q' for a queen, '.' empty.
// One queen per row; columns and both diagonal families are O(1) lookups.
func SolveNQueens(n int) [][]string {
	out := [][]string{}
	cols := make([]bool, n)
	diag1 := make([]bool, 2*n) // r+c
	diag2 := make([]bool, 2*n) // r-c+n
	pos := make([]int, n)      // pos[r] = column of the queen in row r
	var place func(r int)
	place = func(r int) {
		if r == n {
			board := make([]string, n)
			for i, c := range pos {
				row := make([]byte, n)
				for j := range row {
					row[j] = '.'
				}
				row[c] = 'Q'
				board[i] = string(row)
			}
			out = append(out, board)
			return
		}
		for c := 0; c < n; c++ {
			if cols[c] || diag1[r+c] || diag2[r-c+n] {
				continue
			}
			cols[c], diag1[r+c], diag2[r-c+n] = true, true, true
			pos[r] = c
			place(r + 1)
			cols[c], diag1[r+c], diag2[r-c+n] = false, false, false
		}
	}
	place(0)
	return out
}

func main() {
	for _, b := range SolveNQueens(4) {
		for _, row := range b {
			fmt.Println(row)
		}
		fmt.Println()
	}
}
