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
	pos := struct{x, y, aim int}{0, 0, 0}
	for scanner.Scan() {
		segs := strings.Split(scanner.Text(), " ")
		num, _ := strconv.Atoi(segs[1])
		apply(segs[0], num, &pos.x, &pos.y, &pos.aim)
                // your routine here
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(pos.x * pos.y)
}

func apply(op string, val int, x, y, aim *int) {
	switch op {
	case "forward":
		*x += val
		*y += (*aim)*val
	case "up":
		*aim -= val
	case "down":
		*aim += val
	default:
		fmt.Println("unknown op")
	}
}
