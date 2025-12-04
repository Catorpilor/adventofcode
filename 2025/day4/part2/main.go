package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var res int
	matrix := [][]byte{}
	for scanner.Scan() {
		txt := scanner.Text()
		matrix = append(matrix, []byte(txt))

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	for {
		ret := cal(matrix)
		if ret == 0 {
			break
		}
		res += ret
	}
	fmt.Println(res)
}

func cal(matrix [][]byte) int {
	rows := len(matrix)
	if rows == 0 {
		return 0
	}
	cols := len(matrix[0])
	pos := []int{0, 1, 0, -1, 1, 1, -1, -1, 0} // 8 directions
	res := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '@' {
				if probe(matrix, i, j, pos) {
					res++
					matrix[i][j] = '#' // mark as visited
				}
			}
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '#' {
				matrix[i][j] = '.' // mark as deleted
			}
		}
	}
	return res
}

func probe(matrix [][]byte, x, y int, pos []int) bool {
	rows := len(matrix)
	cols := len(matrix[0])
	var count int
	for i := 0; i < 8; i++ {
		dx, dy := pos[i], pos[i+1]
		nx, ny := x+dx, y+dy
		if nx >= 0 && nx < rows && ny >= 0 && ny < cols {
			if matrix[nx][ny] == '@' || matrix[nx][ny] == '#' {
				count++
			}
		}
	}
	return count < 4
}
