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

func calc(matrix [][]int, m, n int) int {
	res := 2*(m+n) - 4
	if m == 2 || n == 2 {
		return res
	}
	flags := make([][]bool, m-2)
	for i := range flags {
		flags[i] = make([]bool, n-2)
	}
	for i := 1; i < m-1; i++ {
		prev := matrix[i][0]
		for l := 1; l < n-1; l++ {
			if matrix[i][l] > prev {
				flags[i-1][l-1] = true
				prev = matrix[i][l]
			}
		}
		prev = matrix[i][n-1]
		for r := n - 2; r > 0; r-- {
			if matrix[i][r] > prev {
				flags[i-1][r-1] = true
				prev = matrix[i][r]
			}
		}
	}
	for i := 1; i < n-1; i++ {
		prev := matrix[0][i]
		for l := 1; l < m-1; l++ {
			if matrix[l][i] > prev {
				prev = matrix[l][i]
				flags[l-1][i-1] = true
			}
		}
		prev = matrix[m-1][i]
		for r := m - 2; r > 0; r-- {
			if matrix[r][i] > prev {
				prev = matrix[r][i]
				flags[r-1][i-1] = true
			}
		}
	}
	for i:=0; i<m-2; i++ {
		for j:=0; j<n-2; j++ {
			if flags[i][j] {
				res++
			}
		}
	}
	return res
}
