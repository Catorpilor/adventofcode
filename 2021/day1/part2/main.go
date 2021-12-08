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
	 prev := -1
	var res int
	for i:=0; i+2<len(nums); i++ {
		if prev == -1 {
			prev = nums[i]+nums[i+1]+nums[i+2]
		}else {
			cur := prev - nums[i-1]+nums[i+2]
			if cur > prev {
				res++
			}
			prev = cur
		}
	}
	fmt.Println(res)
}
