package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var nums []int
	var total int
	for scanner.Scan() {
		segs := strings.Split(scanner.Text(), ",")
		for _, s := range segs {
			num, _ := strconv.Atoi(s)
			total += num
			nums = append(nums, num)
		}
                // your routine here
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	result := float64(total) / float64(len(nums))
	medNum := int(math.Floor(result))
	var res int
	fn := func(n int) int {
		return n*(n+1)/2
	}
	fmt.Printf("medNum:%d, total:%d, len(nums):%d, fn(11)=%d\n", medNum, total, len(nums), fn(11))
	for _, num := range nums {
		res += fn(abs(num - medNum))
	}
	fmt.Println(res)

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
