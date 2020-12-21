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
	var ans int
	for scanner.Scan() {
		txt := scanner.Text()
		ans += helper(txt)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(ans)
}

func helper(s string) int {
	// var preOp byte // preOp stores the previous op of (
	var ans int
	var op byte
	// n := len(s)
	// prevSpc := -1
	segs := strings.Fields(s)
	var leftOpd int
	for _, seg := range segs {
		num, err := strconv.Atoi(seg)
		if err != nil {
			if len(seg) != 1 {
				// with ( or )

			} else {

			}
		}
	}
	// for i:=0; i<n; i++ {
	// 	if s[i] != ' ' {
	// 		continue
	// 	}
	// 	//
	// 	seg := s[prevSpc+1:i]
	// 	if len(seg) > 1 && seg[0] != '(' {

	// 	}

	// }
	return ans
}
