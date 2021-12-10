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
	// set stores the count of the same state
	set := make(map[int]int64) 
	for scanner.Scan() {
                // your routine here
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
		segs := strings.Split(scanner.Text(), ",")
		for _, s := range segs {
			num, _ := strconv.Atoi(s)
			set[num]++
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	for i:=0; i<256;i++ {
		set = updateState(set)
	}
	var res int64
	for _, v := range set {
		res += v
	}
	fmt.Println(res)
}

func updateState(set map[int]int64) map[int]int64{
	ret := make(map[int]int64, len(set))
	ret[6] = set[0]+set[7] // reset + prev
	ret[8] = set[0] // spring
	for i:=1; i<9; i++ {
		if i == 7 {
			continue
		}
		ret[i-1] = set[i]
	}
	return ret
}
