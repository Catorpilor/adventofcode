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
	var mask string
	mem := make(map[int]int)
	for scanner.Scan() {
		txt := scanner.Text()
		if strings.Contains(txt, "mask") {
			mask = txt[7:]
		} else {
			idx := strings.Index(txt, "=")
			val, _ := strconv.Atoi(txt[idx+2:])
			idx1, idx2 := strings.Index(txt, "["), strings.Index(txt, "]")
			pos, _ := strconv.Atoi(txt[idx1+1 : idx2])
			addrs := apply(mask, pos)
			for _, addr := range addrs {
				mem[addr] = val
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	// fmt.Println(mask)
	// fmt.Println(mem)
	var ans int
	for _, v := range mem {
		ans += v
	}
	fmt.Println(ans)
}

func apply(mask string, v int) []int {
	xpos := make([]int, 0, 36)
	for i := 35; i >= 0; i-- {
		if mask[i] != '0' {
			if mask[i] == '1' {
				v |= (1 << (35 - i))
			} else {
				xpos = append(xpos, (35 - i))
			}
		}
	}
	// fmt.Println(xpos)
	loop := int(math.Pow(2, float64(len(xpos))))
	res := make([]int, 0, loop)
	oldV := v
	for i := 0; i < loop; i++ {
		replaceMask := fmt.Sprintf("%0*b", len(xpos), i)
		// fmt.Printf("cur mask: %s, len(mask)=%d\n", replaceMask, len(replaceMask))
		for r := range replaceMask {
			// fmt.Printf("cur r:%d\n", r)
			if replaceMask[r] == '1' {
				v |= 1 << xpos[r]
			} else {
				v &= ^(1 << xpos[r])
			}
		}
		res = append(res, v)
		v = oldV
	}
	return res
}
