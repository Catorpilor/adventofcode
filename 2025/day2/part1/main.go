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
	var res int64
	for scanner.Scan() {
		txt := scanner.Text()
		segs := strings.Split(txt, ",")
		fmt.Printf("segs: %v\n", segs)
		for _, seg := range segs {
			orignals := strings.Split(seg, "-")
			if len(orignals) != 2 {
				continue
			}
			L, _ := strconv.ParseInt(orignals[0], 10, 64)
			R, _ := strconv.ParseInt(orignals[1], 10, 64)
			invalids := findInvalidIDs(L, R)
			for _, id := range invalids {
				res += id
			}
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(res)
}

func countDigits(n int64) int {
	return len(strconv.FormatInt(n, 10))
}

func pow10(k int) int64 {
	result := int64(1)
	for i := 0; i < k; i++ {
		result *= 10
	}
	return result
}

func findInvalidIDs(L, R int64) []int64 {
	var result []int64

	minK := (countDigits(L) + 1) / 2
	maxK := countDigits(R) / 2

	for k := minK; k <= maxK; k++ {
		multiplier := pow10(k) + 1

		// Bounds for k-digit X (no leading zeros)
		var lowerBound int64 = 1
		if k > 1 {
			lowerBound = pow10(k - 1)
		}
		upperBound := pow10(k) - 1

		// X must satisfy: L <= X * multiplier <= R
		minX := (L + multiplier - 1) / multiplier // ceil(L / multiplier)
		maxX := R / multiplier                    // floor(R / multiplier)

		// Clamp to valid k-digit range
		if minX < lowerBound {
			minX = lowerBound
		}
		if maxX > upperBound {
			maxX = upperBound
		}

		for x := minX; x <= maxX; x++ {
			result = append(result, x*multiplier)
		}
	}

	return result
}
