package main

import (
	"os"
	"bufio"
	"strconv"
	"fmt"
)


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var res, cur int
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			if cur > res {
				res = cur
			}
			cur = 0
			continue
		}
		num, _ := strconv.Atoi(scanner.Text())
		cur += num
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(res)
}

