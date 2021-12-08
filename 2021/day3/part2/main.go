package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var pSum []int
	var nums []string
	var total, n int
	for scanner.Scan() {
		n = len(scanner.Text())
		if pSum == nil {
			pSum = make([]int, n)
		}
		s := scanner.Text()
		for i := range s {
			pSum[i] += int(s[i] - '0')
		}
		total++
		nums = append(nums, scanner.Text())
		// your routine here
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	// fmt.Println(nums)

	var o2Rate, co2Rate int64
	if pSum[0] >= total>>1 {
		o2Rate = compute(nums, '1', 0, 1)
		co2Rate = compute(nums, '0', 0, 0)
	} else {
		o2Rate = compute(nums, '0', 0, 1)
		co2Rate = compute(nums, '1', 0, 0)
	}
	fmt.Println(o2Rate * co2Rate)
}

// compute time complexity O(MN), space complexity O(M)
func compute(nums []string, startByte byte, idx, rateType int) int64 {
	// fmt.Printf("nums: %v, startByte: %v, idx: %d, rateType: %d\n", nums, startByte, idx, rateType)
	if len(nums) == 1 {
		num, _ := strconv.ParseInt(nums[0], 2, 64)
		return num
	}
	local := make([]string, len(nums))
	copy(local, nums)
	tmp := make([]string, 0, len(nums))
	var res int
	for _, num := range local {
		n := len(num)
		// fmt.Printf("current num: %s, startByte: %v, idx: %d\n", num, startByte, idx)
		if num[idx] == startByte {
			tmp = append(tmp, num)
			// calculate the next idx
			if idx+1 < n {
				res += int(num[idx+1] - '0')
			}
		}
	}
	// fmt.Printf("tmp: %v, res: %d\n", tmp, res)
	got := len(tmp)
	wanted := (got+1) >> 1
	if res >= wanted {
		if rateType == 1 {
			return compute(tmp, '1', idx+1, rateType)
		}
		return compute(tmp, '0', idx+1, rateType)
	}
	if rateType == 1 {
		return compute(tmp, '0', idx+1, rateType)
	}
	return compute(tmp, '1', idx+1, rateType)
}
