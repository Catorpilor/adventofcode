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
	var res int
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			continue
		}
		strs := strings.Split(txt, "-")
		a_left, _ := strconv.ParseInt(strs[0], 10, 64)
		b_right, _ := strconv.ParseInt(strs[2], 10, 64)
		idx := strings.Index(strs[1], ",")
		a_right, _ := strconv.ParseInt(strs[1][:idx], 10, 64)
		b_left, _ := strconv.ParseInt(strs[1][idx+1:], 10, 64)
		if a_right < b_left || a_left > b_right {
			continue
		}
		res++
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(res)
}

