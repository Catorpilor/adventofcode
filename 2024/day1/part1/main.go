package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var res int
	nums1, nums2 := make([]int, 0, 1000), make([]int, 0, 1000)
	for scanner.Scan() {
		txt := scanner.Text()
		nums := strings.Split(txt, "   ")
		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi(nums[1])
		nums1 = append(nums1, a)
		nums2 = append(nums2, b)
	}
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i := range nums1 {
		res += diff(nums1[i], nums2[i])
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
