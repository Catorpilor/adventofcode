package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	row = 128
	col = 8
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var ans, cur int
	for scanner.Scan() {
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
		cur = process(scanner.Text())
		ans = max(ans, cur)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(ans)
}

func process(s string) int {
	n := len(s)
	// process row
	l, r := 0, row-1
	for i := 0; i < n-3; i++ {
		mid := l + (r-l)>>1
		if s[i] == 'F' {
			// mid := l + (r-l)>>1
			r = mid
		} else {
			if l < r {
				l = mid + 1
			} else {
				l = mid
			}
		}
	}

	// process seat
	sl, sr := 0, col-1
	for i := n - 3; i < n; i++ {
		mid := sl + (sr-sl)>>1
		if s[i] == 'L' {
			sr = mid
		} else {
			if sl < sr {
				sl = mid + 1
			} else {
				sl = mid
			}
		}
	}
	// fmt.Printf("row: %d, seat: %d\n", l, sl)
	return l*8 + sl
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
