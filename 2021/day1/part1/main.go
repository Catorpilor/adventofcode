package main

import (
	"os"
	"bufio"
	"strconv"
	"fmt"
)


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	prev := -1
	var res int
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		if prev != -1 {
			if num > prev {
				res++
			}
		}
		prev = num
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(res)
}
