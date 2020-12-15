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
	var nums []int
	for scanner.Scan() {
		txt := scanner.Text()
		segs := strings.FieldsFunc(txt, func(c rune) bool {
			return c == ','
		})
		for _, seg := range segs {
			num, _ := strconv.Atoi(seg)
			nums = append(nums, num)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	type lastTwo struct {
		first, second int
	}
	store := make(map[int]*lastTwo)
	var last int
	for i, num := range nums {
		if store[num] == nil {
			store[num] = &lastTwo{
				first:  i + 1,
				second: -1,
			}
			last = num
		}
	}
	i := len(nums) + 1
	for i < 2021 {
		var target int
		if store[last].second == -1 {
			// so this round is 0
			target = 0
		} else {
			target = store[last].second - store[last].first
		}
		// fmt.Printf("turn:%d, last:%d, target: %d, store[last]: %v, store[target]:%v\n", i, last, target,
		// store[last], store[target])
		if store[target] == nil {
			store[target] = &lastTwo{
				first:  i,
				second: -1,
			}
		} else {
			if store[target].second != -1 {
				store[target].first = store[target].second
			}
			store[target].second = i
		}
		last = target
		i++
	}
	fmt.Println(last)
}
