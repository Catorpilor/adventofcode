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

		// Count how many times we pass 0 during this rotation
		var crosses int
		if start == 0 {
			// Already at 0, next crossing after 100 steps
			crosses = pad / 100
		} else if dir == 'L' {
			// Going left (decreasing), first hit 0 after `start` steps
			if pad >= start {
				crosses = (pad-start)/100 + 1
			}
		} else {
			// Going right (increasing), first hit 0 after `100-start` steps
			gap := 100 - start
			if pad >= gap {
				crosses = (pad-gap)/100 + 1
			}
		}
		res += crosses

		// Update position
		if dir == 'L' {
			start -= pad
		} else {
			start += pad
		}
		start = ((start % 100) + 100) % 100
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(res)
}
