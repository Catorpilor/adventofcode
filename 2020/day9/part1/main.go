package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var preamble int
	flag.IntVar(&preamble, "preamble", 25, "the length of preamble")
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	store := make([]int, preamble)
	for i := range store {
		store[i] = math.MinInt32
	}
	set := make(map[int]bool, preamble)
	var idx int
	for scanner.Scan() {
		txt := scanner.Text()
		num, _ := strconv.Atoi(txt)
		pos := idx % preamble
		// store[idx] = num
		if idx > preamble && !isValid(store, set, num) {
			fmt.Println(num)
			return
		}
		prev := store[pos]
		if prev != math.MinInt32 {
			delete(set, prev)
		}
		set[num] = true
		store[pos] = num
		idx++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
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
