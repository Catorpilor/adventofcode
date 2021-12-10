package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type dot struct{ x, y int }

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var res int
	gSet := make(map[dot]int)
	for scanner.Scan() {
		processInput(gSet, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(gSet)
	for _, v := range gSet {
		if v >= 2 {
			res++
		}
	}
	fmt.Println(res)
}

func processInput(gset map[dot]int, str string) {
	idx1 := strings.Index(str, ",")
	x1 := str[:idx1]
	idx2 := strings.Index(str, " ")
	y1 := str[idx1+1 : idx2]
	idx1 = strings.LastIndex(str, ",")
	y2 := str[idx1+1:]
	idx2 = strings.LastIndex(str, " ")
	x2 := str[idx2+1 : idx1]
	if x1 != x2 && y1 != y2 {
		fmt.Printf("(%s) not a valid line, skip!\n", str)
		return
	}
	n1, _ := strconv.Atoi(x1)
	m1, _ := strconv.Atoi(y1)
	n2, _ := strconv.Atoi(x2)
	m2, _ := strconv.Atoi(y2)
	diff := 1
	if n1 != n2 {
		if n1 > n2 {
			diff = -1
		}
		tmp := n1
		for tmp != n2+diff {
			cur := dot{tmp, m1}
			gset[cur]++
			tmp += diff
		}
	}

	if m1 != m2 {
		if m1 > m2 {
			diff = -1
		}
		tmp := m1
		for tmp != m2+diff {
			cur := dot{n1, tmp}
			gset[cur]++
			tmp += diff
		}
	}
}
