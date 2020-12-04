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
	var nums []int
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	sort.Ints(nums)
	n := len(nums)
	// use two pointers time complexity O(N^2), space complexity O(1)
	for i := 0; i < n-2; i++ {
		j, k := i+1, n-1
		for j < k {
			cur := nums[i] + nums[j] + nums[k]
			if cur > 2020 {
				k--
			} else if cur < 2020 {
				j++
			} else {
				fmt.Println(nums[i] * nums[j] * nums[k])
				return
			}
		}
	}
}
