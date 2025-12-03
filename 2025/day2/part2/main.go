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

// Multiplier for k-digit pattern repeated n times
// e.g., k=2, n=3: 10^4 + 10^2 + 1 = 10101
func getMultiplier(k, n int) int64 {
	var m int64 = 0
	for i := 0; i < n; i++ {
		m += pow10(i * k)
	}
	return m
}

func findInvalidIDs(L, R int64) []int64 {
	seen := make(map[int64]bool)
	var result []int64

	minDigits := countDigits(L)
	maxDigits := countDigits(R)

	for totalDigits := minDigits; totalDigits <= maxDigits; totalDigits++ {
		for k := 1; k*2 <= totalDigits; k++ {
			if totalDigits%k != 0 {
				continue
			}
			n := totalDigits / k
			multiplier := getMultiplier(k, n)
			lowerBound := int64(1)
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
				id := x * multiplier
				if !seen[id] {
					seen[id] = true
					result = append(result, id)
				}
			}
		}
	}
	return result
}
