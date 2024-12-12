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
	dirs   = [9]int{-1, 0, 1, 1, 0, -1, -1, 1, -1}
	wanted = "XMAS"
)

// search searches the matrix from position(x,y) see if there is a match
// it returns the total matches from (x,y)
func search(b [][]byte, x, y int) int {
	var res int
	m := len(b)
	n := len(b[0])
	for i := 0; i < len(dirs)-1; i++ {
		matched := 1
		for j := 1; j < 4; j++ {
			nx, ny := x+j*dirs[i], y+j*dirs[i+1]
			if nx < 0 || nx >= m || ny < 0 || ny >= n {
				break
			}
			if b[nx][ny] == wanted[j] {
				matched++
				continue
			}
		}
		if matched == len(wanted) {
			res++
		}
	}
	return res
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
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 'X' {
				res += search(matrix, i, j)
			}
		}
	}
	fmt.Println(res)
}
