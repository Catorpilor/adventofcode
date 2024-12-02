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
	var res int
	nums1 := make([]int, 0, 1000)
	hset := make(map[int]int, 1000)
	for scanner.Scan() {
		txt := scanner.Text()
		nums := strings.Split(txt, "   ")
		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi(nums[1])
		nums1 = append(nums1, a)
		hset[b]++
	}
	for i := range nums1 {
		if v, found := hset[nums1[i]]; !found {
			res += 0
		} else {
			res += v * nums1[i]
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(res)
}

func diff(a, b int) int {
	if a >= b {
		return a - b
	}
	return b - a
}
