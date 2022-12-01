package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var max1, max2, max3, cur int
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			if cur > max1 {
				max2, max3 = max1, max2
				max1 = cur
			} else if cur > max2 {
				max3 = max2
				max2 = cur
			} else if cur > max3 {
				max3 = cur
			}
			cur = 0
			continue
		}
		num, _ := strconv.Atoi(scanner.Text())
		cur += num
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if cur > max1 {
		max2, max3 = max1, max2
		max1 = cur
	} else if cur > max2 {
		max3 = max2
		max2 = cur
	} else if cur > max3 {
		max3 = cur
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	// fmt.Printf("top 3 is (%d, %d, %d)\n", max1, max2, max3)
	fmt.Println(max1 + max2 + max3)
}
