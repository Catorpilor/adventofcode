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
	total := make([]int, 0, 200000)
	for scanner.Scan() {
                // your routine here
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
		segs := strings.Split(scanner.Text(), ",")
		for _, s := range segs {
			num, _ := strconv.Atoi(s)
			total = append(total, num)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(total)
	for i:=0; i<80; i++ {
		for idx := range total {
			if total[idx] == 0 {
				total = append(total, 8)
				total[idx] = 6
				continue
			}
			total[idx]--
		}
	}
	fmt.Println(len(total))
} 
