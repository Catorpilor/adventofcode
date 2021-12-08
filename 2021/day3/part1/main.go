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
	var total, n int
	for scanner.Scan() {
		n = len(scanner.Text())
		if pSum == nil {
			pSum = make([]int, n)
		}
		s := scanner.Text()
		for i := range s {
			pSum[i] += int(s[i]-'0')
		}
		total++
                // your routine here
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	gamma, epsilon := make([]byte, 0, n), make([]byte, 0, n)
	mc := total>>1
	for i:=0; i<n; i++ {
		if pSum[i] > mc {
			gamma = append(gamma,'1')
			epsilon = append(epsilon, '0')
			continue
		}
		gamma = append(gamma, '0')
		epsilon = append(epsilon, '1')
	}
	// fmt.Printf("gamma: %v, epsilon: %v\n", gamma, epsilon)
	gammaNum, _ := strconv.ParseInt(string(gamma), 2, 64)
	epsilonNum, _ := strconv.ParseInt(string(epsilon), 2, 64)
	fmt.Println(gammaNum * epsilonNum)
}
