package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type op struct {
	cmd   string
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
		cycleSum = append(cycleSum, cycleSum[pre1-1]+cycle)
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(preSum)
	// wants := []int{20, 60, 100, 140, 180, 220}
	// for _, want := range wants {
	// 	idx := sort.SearchInts(cycleSum, want)
	// 	fmt.Printf("idx: %d, %d-th x=%d\n",idx, want, preSum[cycleSum[idx-1]])
	// 	res += want * preSum[cycleSum[idx-1]]
	// }
	pos := [3]int{0, 1, 2}
	var sb bytes.Buffer
	var idx int
	for i := 1; i < len(preSum); i++ {
		// fmt.Println(pos)
		if check(idx, pos) {
			sb.WriteByte('#')
		} else {
			sb.WriteByte('.')
		}
		pos[1] = preSum[i]
		pos[0], pos[2] = pos[1]-1, pos[1]+1
		idx++
		if idx%40 == 0 {
			idx = 0
			sb.WriteByte('\n')
		}
	}
	fmt.Println(sb.String())
}

func check(want int, a [3]int) bool {
	for i := 0; i < 3; i++ {
		if a[i] == want {
			return true
		}
	}
	return false
}
