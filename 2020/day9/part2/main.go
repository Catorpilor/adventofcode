package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var preamble, wanted int
	flag.IntVar(&preamble, "preamble", 25, "the length of preamble")
	// hard coded the invalid num got from part1
	// for the testcase it should be 127.
	flag.IntVar(&wanted, "wanted", 31161678, "target sum")
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	var store []int
	for scanner.Scan() {
		txt := scanner.Text()
		num, _ := strconv.Atoi(txt)
		store = append(store, num)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	helper(store, wanted)
}

func helper(store []int, target int) {
	curSum := store[0]
	l, n := 0, len(store)
	for i := 1; i < n; i++ {
		for curSum > target && l < i-1 {
			curSum -= store[l]
			l++
		}
		if curSum == target {
			calWeak(store, l, i)
			return
		}
		curSum += store[i]
	}
}

func isValid(store []int, set map[int]bool, cur int) bool {
	for _, num := range store {
		want := cur - num
		if set[want] {
			return true
		}
	}
	return false
}

func calWeak(store []int, l, r int) {
	cmin, cmax := store[l], store[l]
	for i := l + 1; i < r; i++ {
		if store[i] < cmin {
			cmin = store[i]
		}
		if store[i] > cmax {
			cmax = store[i]
		}
	}
	fmt.Println(cmin + cmax)
}
