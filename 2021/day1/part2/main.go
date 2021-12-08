package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
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
	cals := make([]int, 0, len(nums))
	preSum := -1
	for i:=0; i+2<len(nums); i++ {
		if preSum == -1 {
			preSum = nums[i]+nums[i+1]+nums[i+2]
		}else {
			preSum = preSum - nums[i-1]+nums[i+2]
		}
		cals = append(cals, preSum)
	}
	prev := -1
	var res int
	for i := range cals {
		if prev != -1 {
			if cals[i] > prev {
				res++
			}
		}
		prev = cals[i]
	}
	fmt.Println(res)
}
