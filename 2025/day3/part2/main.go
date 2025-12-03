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

// processLineV2 finds the largest k-digit number that can be formed
// by selecting digits from s in order.
// using a stack-based greedy approach.
func processLineV2(s string, k int) int64 {
	n := len(s)
	if n < k {
		return -1
	}
	var result int64
	stack := make([]byte, 0, k)
	removals := n - k

	for i := 0; i < n; i++ {
		for len(stack) > 0 && removals > 0 && stack[len(stack)-1] < s[i] {
			stack = stack[:len(stack)-1]
			removals--
		}
		stack = append(stack, s[i])
	}
	result, _ = strconv.ParseInt(string(stack[:k]), 10, 64)
	return result
}
