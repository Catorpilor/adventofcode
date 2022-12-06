package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			continue
		}
		// fmt.Printf("for txt(%s) the first marker after %d character arrived\n", txt, calculate(txt))
		fmt.Println(calculate(txt))
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func calculate(txt string) int {
	n := len(txt)
	if n < 4 {
		return -1
	}
	seen := make([]int, 26)
	for i := range seen {
		seen[i] = -1
	}
	var res, begin_pos int
	for i := 0; i < n; i++ {
		pos := int(txt[i] - 'a')
		// fmt.Printf("cur char(%s) at pos %d\n", string(txt[i]), i)
		if seen[pos] != -1 {
			// fmt.Printf("char(%s) seen before at pos(%d),update begin_pos from (%d) to (%d)\n", string(txt[i]), seen[pos], begin_pos, seen[pos]+1)
			if seen[pos] >= begin_pos {
				// only upate begin_pos if necessary
				begin_pos = seen[pos] + 1
			}
		}
		seen[pos] = i
		if i-begin_pos == 13 {
			res = i + 1
			break
		}
	}
	return res
}
