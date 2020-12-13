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
	var startTs int
	var lines int
	var buses []int
	for scanner.Scan() {
		txt := scanner.Text()
		lines++
		if lines == 1 {
			startTs, _ = strconv.Atoi(txt)
		} else {
			segs := strings.FieldsFunc(txt, func(c rune) bool {
				return c == ','
			})
			for _, seg := range segs {
				if seg != "x" {
					num, _ := strconv.Atoi(seg)
					buses = append(buses, num)
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	// fmt.Println(startTs)
	// fmt.Println(buses)
	dep, si := startTs, 0
	for _, id := range buses {
		did := startTs / id * id
		if did < startTs {
			did += id
		}
		diff := did - startTs
		if diff < dep {
			dep = diff
			si = id
		}
	}
	fmt.Println(dep * si)
}
