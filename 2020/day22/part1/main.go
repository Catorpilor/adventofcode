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
	var group [2][]int
	var idx int
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			idx++
			continue
		}
		if strings.Contains(txt, ":") {
			continue
		} else {
			num, _ := strconv.Atoi(txt)
			group[idx] = append(group[idx], num)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	p1, p2 := group[0], group[1]
	for len(p1) != 0 && len(p2) != 0 {
		// pick a card
		t1, t2 := p1[0], p2[0]
		p1 = p1[1:]
		p2 = p2[1:]
		if t1 < t2 {
			p2 = append(p2, t2, t1)
		} else {
			p1 = append(p1, t1, t2)
		}
	}
	res := p1
	if len(p2) != 0 {
		res = p2
	}
	var ans int
	n := len(res)
	for i := n - 1; i >= 0; i-- {
		ans += res[i] * (n - i)
	}
	fmt.Println(ans)
}
