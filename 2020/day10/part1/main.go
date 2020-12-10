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
	var adps []int
	var cmax int
	for scanner.Scan() {
		txt := scanner.Text()
		num, _ := strconv.Atoi(txt)
		adps = append(adps, num)
		if num > cmax {
			cmax = num
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	adps = append(adps, 0)
	sort.Ints(adps)
	var ones, threes int
	n := len(adps)
	for i := 0; i < n-1; i++ {
		diff := adps[i+1] - adps[i]
		if diff == 1 {
			ones++
		} else if diff == 3 {
			threes++
		}
	}
	threes++
	fmt.Printf("ones:%d, threes:%d\n", ones, threes)
	fmt.Println(ones * threes)
}
