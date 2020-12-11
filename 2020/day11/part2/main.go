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
		// thing could mess up i dont know why.
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
		display(metrix)
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
	// fmt.Printf("pos(i:%d, j:%d) occupied: %d\n", i, j, occupied)
	if metrix[i][j] == '#' {
		if occupied >= 5 {
			metrix[i][j] = 'Z'
		}
	} else {
		// L
		if occupied == 0 {
			metrix[i][j] = 'X'
		}
	}
}

func count(metrix [][]byte, i, j, m, n int) int {
	var ans int
	tmpj := j
	for tmpj > 0 {
		tmpj--
		if metrix[i][tmpj] == '.' {
			continue
		} else {
			if metrix[i][tmpj] == '#' || metrix[i][tmpj] == 'Z' {
				ans++
			}
			break
		}

	}
	tmpj = j
	for tmpj < n-1 {
		tmpj++
		if metrix[i][tmpj] == '.' {
			continue
		} else {
			if metrix[i][tmpj] == '#' || metrix[i][tmpj] == 'Z' {
				ans++
			}
			break
		}
	}
	tmpi := i
	for tmpi > 0 {
		tmpi--
		if metrix[tmpi][j] == '.' {
			continue
		} else {
			if metrix[tmpi][j] == '#' || metrix[tmpi][j] == 'Z' {
				ans++
			}
			break
		}
	}
	tmpi = i
	for tmpi < m-1 {
		tmpi++
		if metrix[tmpi][j] == '.' {
			continue
		} else {
			if metrix[tmpi][j] == '#' || metrix[tmpi][j] == 'Z' {
				ans++
			}
			break
		}
	}
	tmpi, tmpj = i, j
	for tmpi > 0 && tmpj < n-1 {
		// up right
		tmpi, tmpj = tmpi-1, tmpj+1
		if metrix[tmpi][tmpj] == '.' {
			continue
		} else {
			if metrix[tmpi][tmpj] == '#' || metrix[tmpi][tmpj] == 'Z' {
				ans++
			}
			break
		}
	}
	tmpi, tmpj = i, j
	for tmpj > 0 && tmpi > 0 {
		// up left
		tmpi, tmpj = tmpi-1, tmpj-1
		if metrix[tmpi][tmpj] == '.' {
			continue
		} else {
			if metrix[tmpi][tmpj] == '#' || metrix[tmpi][tmpj] == 'Z' {
				ans++
			}
			break
		}
	}
	tmpi, tmpj = i, j
	for tmpi < m-1 && tmpj < n-1 {
		// down right
		tmpi, tmpj = tmpi+1, tmpj+1
		if metrix[tmpi][tmpj] == '.' {
			continue
		} else {
			if metrix[tmpi][tmpj] == '#' || metrix[tmpi][tmpj] == 'Z' {
				ans++
			}
			break
		}
	}
	tmpi, tmpj = i, j
	for tmpi < m-1 && tmpj > 0 {
		// down left
		tmpi, tmpj = tmpi+1, tmpj-1
		if metrix[tmpi][tmpj] == '.' {
			continue
		} else {
			if metrix[tmpi][tmpj] == '#' || metrix[tmpi][tmpj] == 'Z' {
				ans++
			}
			break
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
