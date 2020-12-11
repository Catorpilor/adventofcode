package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var metrix [][]byte
	for scanner.Scan() {
		sb := scanner.Text()
		// for i := range sb {
		// 	if sb[i] == 'L' {
		// 		sb[i] = '#'
		// 	}
		// }
		metrix = append(metrix, []byte(sb))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	// for _, row := range metrix {
	// 	fmt.Println(string(row))
	// }
	display(metrix)
	rounds := 0
	// 2nd round
	m := len(metrix)
	n := len(metrix[0])

	prev := make([][]byte, m)
	for i := range prev {
		prev[i] = make([]byte, n)
		copy(prev[i], metrix[i])
	}

	for {
		rounds++
		// wanted := '#'
		// target := 'Z'
		// dst := 'L'
		// if rounds&1 != 0 {
		// 	wanted = 'L'
		// 	target = 'X'
		// 	dst = '#'
		// }
		// fmt.Printf("%d rounds, wanted: %s, target: %s, dst: %s\n", rounds,
		// 	string(wanted), string(target), string(dst))
		// fmt.Printf("%d round\n", rounds)
		visited := make([]bool, m*n)
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				// fmt.Printf("outer: visited[i*n+j]=%t, metrix[i][j]=%s\n", visited[i*n+j], string(metrix[i][j]))
				if !visited[i*n+j] && metrix[i][j] != '.' {
					// fmt.Printf("entring dfs at pos(i:%d, j:%d)\n", i, j)
					dfs(metrix, visited, i, j, m, n)
				}
			}
		}
		// display(metrix)
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if metrix[i][j] == 'Z' {
					metrix[i][j] = 'L'
				} else if metrix[i][j] == 'X' {
					metrix[i][j] = '#'
				}
			}
		}
		// display(prev)
		// display(metrix)
		if isEqual(prev, metrix) {
			break
		}
		for i := range prev {
			// prev[i] = make([]byte, n)
			copy(prev[i], metrix[i])
			// fmt.Println(string(metrix[i]))
		}
	}
	var ans int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if metrix[i][j] == '#' {
				ans++
			}
		}
	}
	fmt.Println(ans)

}

func dfs(metrix [][]byte, visited []bool, i, j, m, n int) {
	visited[i*n+j] = true
	occupied := count(metrix, i, j, m, n)
	if metrix[i][j] == '#' {
		if occupied >= 4 {
			metrix[i][j] = 'Z'
		}
	} else {
		// L
		if occupied == 0 {
			metrix[i][j] = 'X'
		}
	}
}

var (
	dirs = [][]int{[]int{1, 0}, []int{1, 1}, []int{1, -1}, []int{-1, -1},
		[]int{-1, 0}, []int{-1, 1}, []int{0, 1}, []int{0, -1}}
)

func count(metrix [][]byte, i, j, m, n int) int {
	var ans int
	for _, dir := range dirs {
		nx, ny := i+dir[0], j+dir[1]
		if nx < 0 || nx >= m || ny < 0 || ny >= n {
			continue
		}
		if metrix[nx][ny] == '#' || metrix[nx][ny] == 'Z' {
			ans++
		}
	}
	return ans
}

func isEqual(src, dst [][]byte) bool {
	for i := range src {
		if !bytes.Equal(src[i], dst[i]) {

			return false
		}
	}
	return true
}

func display(metrix [][]byte) {
	for _, row := range metrix {
		fmt.Println(string(row))
	}
	fmt.Println("---------")
}
