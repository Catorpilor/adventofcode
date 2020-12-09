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
	flag.IntVar(&wanted, "wanted", 31161678, "target sum")
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	// store := make([]int, preamble)
	var store []int
	// for i := range store {
	// 	store[i] = math.MinInt32
	// }
	// set := make(map[int]bool, preamble)
	for scanner.Scan() {
		txt := scanner.Text()
		num, _ := strconv.Atoi(txt)
		store = append(store, num)

		// pos := idx % preamble
		// // store[idx] = num
		// if idx > preamble && !isValid(store, set, num) {
		// 	fmt.Println(num)
		// 	return
		// }
		// prev := store[pos]
		// if prev != math.MinInt32 {
		// 	delete(set, prev)
		// }
		// set[num] = true
		// store[pos] = num
		// idx++
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
