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
	var lines int
	var buses []int
	var offsets []int
	for scanner.Scan() {
		txt := scanner.Text()
		lines++
		if lines != 1 {
			// startTs, _ = strconv.Atoi(txt)
			segs := strings.FieldsFunc(txt, func(c rune) bool {
				return c == ','
			})
			offset := 0
			for _, seg := range segs {
				if seg != "x" {
					num, _ := strconv.Atoi(seg)
					buses = append(buses, num)
					offsets = append(offsets, offset)
					// if i == 0 {
					// 	offsets = append(offsets, 0)
					// } else {
					// 	offsets = append(offsets, offset)
					// 	offset = 1
					// }
				}
				offset++
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	// fmt.Println(buses)
	// fmt.Println(offsets)
	// st := 100_000_000_000_000
	// // st := 100_179_6
	// startTs := st / buses[0] * buses[0]
	// if startTs < st {
	// 	startTs += buses[0]
	// }
	// fmt.Println(startTs)
	n := len(buses)
	n1 := int64(buses[0])
	incr := int64(n1)

	for i := 1; i < n; i++ {
		n1, incr = lcp(n1, int64(buses[i]), int64(offsets[i]), incr)
	}
	fmt.Println(n1)
	// for {
	// 	count := 1
	// 	for i := n - 1; i > 0; i-- {
	// 		if (startTs+offsets[i])%buses[i] != 0 {
	// 			break
	// 		} else {
	// 			count++
	// 			fmt.Printf("startTs:%d, offset: %d, match bus: %d, count:%d\n", startTs,
	// 				offsets[i], buses[i], count)
	// 		}
	// 	}
	// 	if count == n {
	// 		break
	// 	}
	// 	startTs += buses[0]
	// }
	// dep, si := startTs, 0
	// for _, id := range buses {
	// 	did := startTs / id * id
	// 	if did < startTs {
	// 		did += id
	// 	}
	// 	diff := did - startTs
	// 	if diff < dep {
	// 		dep = diff
	// 		si = id
	// 	}
	// }
	// fmt.Println(dep * si)
	// fmt.Println(startTs)
}

func lcp(n1, n2, offset, incr int64) (int64, int64) {
	i := n1
	for (i+offset)%n2 != 0 {
		i += incr
	}
	return i, lcm(incr, n2)
}

// greatest common divisor (GCD) via Euclidean algorithm
func gcd(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func lcm(a, b int64, integers ...int64) int64 {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}
