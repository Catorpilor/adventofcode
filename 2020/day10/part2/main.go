package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var adps []int
	var cmax int
	for scanner.Scan() {
		txt := scanner.Text()
		num, _ := strconv.Atoi(txt)
		adps = append(adps, num)
		if num > cmax {
			cmax = num
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	// this is the hard coded version
	// no formular found so far.
	code := map[int]int{
		0: 1,
		1: 1,
		2: 2,
		3: 4,
		4: 7,
	}
	adps = append(adps, 0)
	sort.Ints(adps)
	adps = append(adps, cmax+3)
	n := len(adps)
	diffs := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		diff := adps[i+1] - adps[i]
		diffs[i] = diff
	}
	ans := 1
	var curSeg int
	for i := 0; i < n-1; i++ {
		if diffs[i] != 1 {
			// fmt.Printf("curSeg: %d\n", curSeg)
			ans *= code[curSeg]
			curSeg = 0
			continue
		}
		curSeg++
	}
	fmt.Println(ans)
}
