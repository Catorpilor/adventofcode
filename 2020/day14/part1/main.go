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
			mem[pos] = apply(mask, val)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(mask)
	fmt.Println(mem)
	var ans int
	for _, v := range mem {
		ans += v
	}
	fmt.Println(ans)
}

func apply(mask string, v int) int {
	for i := 35; i >= 0; i-- {
		if mask[i] != 'X' {
			if mask[i] == '1' {
				v |= (1 << (35 - i))
			} else {
				v &= ^(1 << (35 - i))
			}
		}
	}
	return v
}
