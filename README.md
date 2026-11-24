# N-Queens

> Striver SDE Sheet 056 · Day 10 Backtracking — one queen per row; columns and the two diagonal families as boolean arrays make safety O(1). Go, stdlib only, tested.

## Problem

Place n queens on an n×n chessboard so that none attacks another (no shared row, column or diagonal). Return every solution as a board drawing.

## Approach

Rows are free: placing exactly one queen per row eliminates row conflicts by construction, leaving columns and diagonals. The indexing insight that removes all scanning: every ↗ diagonal shares a constant r+c and every ↘ diagonal a constant r−c, so three boolean arrays (`cols`, `diag1[r+c]`, `diag2[r−c+n]`) answer "is this square attacked?" in O(1). Backtracking marks the three lines, recurses to the next row, and unmarks. Verified against the known solution counts (92 for n=8).

**Complexity:** exponential in n (it must be — solutions are exponential); O(n) state.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
