package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var grid []string
	for scanner.Scan() {
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
		grid = append(grid, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	m := len(grid)
	n := len(grid[0])
	// fmt.Println(m, n)
	// right 3 down 1
	px, py := 0, 0
	var ans int
	for py < m-1 {
		py++
		px += 3
		tpx := px
		if tpx >= n {
			tpx = px % n
		}
		// fmt.Printf("next stop: px:%d, py:%d, tpx: %d, grid[%d][%d]=%c\n", px, py, tpx, py, tpx, grid[py][tpx])
		if grid[py][tpx] == '#' {
			ans++
		}
	}
	fmt.Println(ans)
}
