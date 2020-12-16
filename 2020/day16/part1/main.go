package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cond struct {
	l, r int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var flag bool
	regs := make(map[string][]cond)
	var tickts []int
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			continue
		}
		por := strings.Index(txt, " or ")
		if por != -1 {
			fmt.Println(txt)
			posCo := strings.Index(txt, ":")
			n := len(txt)
			// two segs posCo+2: por-1, por+1:
			var dashIdx int
			for i := posCo + 2; i < por; i++ {
				if txt[i] == '-' {
					dashIdx = i
					break
				}
			}
			fmt.Printf("posCo:%d, dashIdx:%d\n", posCo, dashIdx)
			l, _ := strconv.Atoi(txt[posCo+2 : dashIdx])
			r, _ := strconv.Atoi(txt[dashIdx+1 : por])
			regs[txt[:posCo]] = append(regs[txt[:posCo]], cond{l: l, r: r})
			for i := por + 4; i < n; i++ {
				if txt[i] == '-' {
					dashIdx = i
					break
				}
			}
			l, _ = strconv.Atoi(txt[por+4 : dashIdx])
			r, _ = strconv.Atoi(txt[dashIdx+1 : n])
			regs[txt[:posCo]] = append(regs[txt[:posCo]], cond{l: l, r: r})
		} else {
			// fmt.Println(txt)
			if flag {
				segs := strings.FieldsFunc(txt, func(c rune) bool {
					return c == ','
				})
				for _, seg := range segs {
					num, _ := strconv.Atoi(seg)
					tickts = append(tickts, num)
				}
			}
			if strings.Contains(txt, "your") {
				scanner.Scan()
				localtext := scanner.Text()
				fmt.Printf("your ticket: %s\n", localtext)
				// ignore for now
			} else if strings.Contains(txt, "nearby") {
				flag = true
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(regs)
	fmt.Println(tickts)
	var ans int
	for _, num := range tickts {
		if !isValid(regs, num) {
			ans += num
		}
	}
	fmt.Println(ans)
}

func isValid(segs map[string][]cond, num int) bool {
	for _, v := range segs {
		// fmt.Printf("conds: %v, num:%d\n", v, num)
		for _, c := range v {
			if c.l <= num && c.r >= num {
				return true
			}
		}
	}
	return false
}
