package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var res int
	var matrix [][]int
	var n, m int
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			continue
		}
		n = len(txt)
		row := make([]int, 0, n)
		for _, c := range txt {
			row = append(row, int(c-'0'))
		}
		matrix = append(matrix, row)
		m++
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	// calulate inner
	res = calc(matrix, m, n)
	fmt.Println(res)
}

type pos struct{
	left, top, right, bottom int
}

func (p pos) Val() int {
	return p.left * p.right * p.top * p.bottom
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func calc(matrix [][]int, m, n int) int {
	var res int
	for i:=0; i<m; i++ {
		for j:=0; j<n; j++ {
			if i== 0 || j == 0 || i == m-1 || j== m-1 {
				continue
			}
			var left, top, right, bottom  int
			for d := i-1; d >= 0; d-- {
				top++
				if matrix[d][j] >= matrix[i][j] {break}
			}
			for d := j+1; d<n; d++ {
				right++
				if matrix[i][d] >= matrix[i][j] {break}
			}
			for d := i+1; d < m; d++ {
				bottom++
				if matrix[d][j] >= matrix[i][j] {break}
			}
			for d := j-1; d>=0; d-- {
				left++
				if matrix[i][d] >= matrix[i][j] {break}
			}
			res = max(res, left*top*right*bottom)
		}
	}
	return res
}
