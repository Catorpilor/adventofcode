package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var res int
	for scanner.Scan() {
		txt := scanner.Text()
		res += processLine(txt)

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(res)
}

func processLine(s string) int {
	n := len(s)
	if n < 2 {
		return -1
	}

	var result int
	maxSoFar := int(s[n-1] - '0') // Start with last digit

	for i := n - 2; i >= 0; i-- {
		digit := int(s[i] - '0')
		twoDigit := digit*10 + maxSoFar
		if twoDigit > result {
			result = twoDigit
		}
		if digit > maxSoFar {
			maxSoFar = digit
		}
	}
	return result
}
