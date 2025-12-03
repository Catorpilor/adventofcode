package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var res int64
	for scanner.Scan() {
		txt := scanner.Text()
		res += processLine(txt, 12)

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(res)
}

// processLine finds the largest k-digit number that can be formed
// by selecting digits from s in order.
// using a greedy approach.
func processLine(s string, k int) int64 {
	n := len(s)
	if n < k {
		return -1
	}
	buf := make([]byte, k)

	var result int64
	start := 0
	for i := 0; i < k; i++ {
		end := n - k + i
		// find max digit in range [start, end]
		maxIdx := start
		for j := start + 1; j <= end; j++ {
			if s[j] > s[maxIdx] {
				maxIdx = j
			}
		}
		buf[i] = s[maxIdx]
		start = maxIdx + 1
	}

	result, _ = strconv.ParseInt(string(buf), 10, 64)
	return result
}
