package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type op struct{
	cmd string
	value int
}




func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cycleSum := []int{0}
	preSum := []int{1}
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			continue
		}
		pre1 := len(cycleSum)
		pre2 := len(preSum)
		strs := strings.Split(txt, " ")
		cycle := 1
		var toAdd int
		if strs[0] == "addx" {
			cycle = 2
			toAdd, _ = strconv.Atoi(strs[1])
			preSum = append(preSum, preSum[pre2-1])
			pre2 = len(preSum)
		}
		preSum = append(preSum, preSum[pre2-1]+toAdd)
		cycleSum = append(cycleSum, cycleSum[pre1-1] + cycle)
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	var res int
	fmt.Println(res)
	fmt.Println(cycleSum)
	fmt.Println(preSum)
	wants := []int{20, 60, 100, 140, 180, 220}
	for _, want := range wants {
		idx := sort.SearchInts(cycleSum, want)
		fmt.Printf("idx: %d, %d-th x=%d\n",idx, want, preSum[cycleSum[idx-1]])
		res += want * preSum[cycleSum[idx-1]]
	}
	fmt.Println(res)
}

