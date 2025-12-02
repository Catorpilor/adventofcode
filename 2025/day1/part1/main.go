package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var res int
	start := 50
	for scanner.Scan() {
		txt := scanner.Text()
		dir := txt[0]
		pad, _ := strconv.Atoi(txt[1:])
		if dir == 'R' {
			start += pad
		} else {
			start -= pad
		}
		start = ((start % 100) + 100) % 100 // Always normalizes to 0-99
		if start == 0 {
			res++
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(res)
}
