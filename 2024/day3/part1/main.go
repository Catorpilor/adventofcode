package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

const pattern = `mul\((\d+),(\d+)\)`

var re = regexp.MustCompile(pattern)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var res int
	for scanner.Scan() {
		txt := scanner.Text()
		matches := re.FindAllStringSubmatch(txt, -1)
		for _, m := range matches {
			var x, y int
			fmt.Sscanf(m[1], "%d", &x)
			fmt.Sscanf(m[2], "%d", &y)
			res += x * y
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(res)
}
