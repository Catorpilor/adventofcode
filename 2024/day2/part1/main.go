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
		if handle(txt) {
			res++
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(res)
}

func handle(txt string) bool {
	t := strings.Split(txt, " ")
	n := len(t)
	l, _ := strconv.Atoi(t[0])
	var f int // 1 means decreasing, -1 means increasing, default means not set
	for i := 1; i < n; i++ {
		r, _ := strconv.Atoi(t[i])
		if r == l {
			// diff at least 1
			return false
		} else if r > l {
			if f > 0 {
				// alreay decreasing
				return false
			}
			if r-l > 3 {
				return false
			}
			f = -1
		} else {
			if f < 0 {
				return false
			}
			if l-r > 3 {
				return false
			}
			f = 1
		}
		l = r
	}
	return true
}
