package main

import (
	"bufio"
	"fmt"
	"os"
)

type dir int

const (
	POINT dir = iota
	R
	L
	U
	D
	UL
	UR
	DL
	DR
)

var (
	dirs   = [5]int{-1, -1, 1, 1, -1}
	wanted = "XMAS"
)

// search searches the matrix from position(x,y) see if there is a match
// it returns the total matches from (x,y)
func search(b [][]byte, x, y int) int {
	var res int
	x1, y1 := x-1, y-1
	x2, y2 := x+1, y+1
	x3, y3 := x-1, y+1
	x4, y4 := x+1, y-1
	if b[x1][y1] == 'S' && b[x2][y2] == 'M' || b[x1][y1] == 'M' && b[x2][y2] == 'S' {
		res++
	}
	if b[x3][y3] == 'S' && b[x4][y4] == 'M' || b[x3][y3] == 'M' && b[x4][y4] == 'S' {
		res++
	}
	if res == 2 {
		return 1
	}

	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var res int
	matrix := make([][]byte, 0, 255)

	for scanner.Scan() {
		txt := scanner.Text()
		matrix = append(matrix, []byte(txt))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	m := len(matrix)
	n := len(matrix[0])
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if matrix[i][j] == 'A' {
				res += search(matrix, i, j)
			}
		}
	}
	fmt.Println(res)
}
