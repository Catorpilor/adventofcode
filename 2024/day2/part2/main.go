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
	for scanner.Scan() {
		txt := scanner.Text()
		ss := strings.Fields(txt)
		nums := make([]int, 0, len(ss))
		for _, v := range ss {
			vv, _ := strconv.Atoi(v)
			nums = append(nums, vv)
		}
		if handleA(nums) {
			res++
		} else {
			if handleB(nums) {
				res++
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(res)
}

func handleA(nums []int) bool {
	if nums[0] >= nums[1] {
		for i := 1; i < len(nums); i++ {
			if nums[i-1]-nums[i] < 1 || nums[i-1]-nums[i] > 3 {
				return false
			}
		}
	}
	if nums[0] <= nums[1] {
		for i := 1; i < len(nums); i++ {
			if nums[i]-nums[i-1] < 1 || nums[i]-nums[i-1] > 3 {
				return false
			}
		}
	}
	return true
}

func handleB(nums []int) bool {
	for idx := range nums {
		narr := make([]int, 0, len(nums))
		for i, v := range nums {
			if i != idx {
				narr = append(narr, v)
			}
		}
		if handleA(narr) {
			return true
		}
	}
	return false
}
