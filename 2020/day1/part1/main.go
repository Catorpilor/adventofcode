package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	set := make(map[int]bool)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		if set[2020-num] {
			fmt.Println(num * (2020 - num))
			break
		}
		set[num] = true
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
