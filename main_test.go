package main

import (
	"reflect"
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	got := SolveNQueens(4)
	want := [][]string{
		{".Q..", "...Q", "Q...", "..Q."},
		{"..Q.", "Q...", "...Q", ".Q.."},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SolveNQueens(4) = %v, want %v", got, want)
	}

	// Known solution counts.
	counts := map[int]int{1: 1, 2: 0, 3: 0, 5: 10, 6: 4, 8: 92}
	for n, want := range counts {
		if got := len(SolveNQueens(n)); got != want {
			t.Errorf("n=%d: %d solutions, want %d", n, got, want)
		}
	}
}
