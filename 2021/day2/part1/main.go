package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	pos := struct{x, y int}{0, 0}
	for scanner.Scan() {
		segs := strings.Split(scanner.Text(), " ")
		num, _ := strconv.Atoi(segs[1])
		apply(segs[0], num, &pos.x, &pos.y)
                // your routine here
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(pos.x * pos.y)
}

func apply(op string, val int, x, y *int) {
	switch op {
	case "forward":
		*x += val
	case "up":
		*y -= val
	case "down":
		*y += val
	default:
		fmt.Println("unknown op")
	}
}
