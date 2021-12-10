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
	var nums []int
	set := make(map[int]int)
	for scanner.Scan() {
		segs := strings.Split(scanner.Text(), ",")
		for _, s := range segs {
			num, _ := strconv.Atoi(s)
			nums = append(nums, num)
			set[num]++
		}
		// your routine here
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	// find the most common num
	var occ, mostComm int
	for k, v := range set {
		if v > occ {
			mostComm = k
			occ = v
		}
	}
	fmt.Println(mostComm)
	sort.Ints(nums)
	n := len(nums)
	med := nums[n/2]
	var res1, res2 int
	for _, num := range nums {
		res1 += abs(num - mostComm)
		res2 += abs(num - med)
	}
	fmt.Println(res1, res2)

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
