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
	var tickets [][]int
	var yourTicket []int
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			continue
		}
		por := strings.Index(txt, " or ")
		if por != -1 {
			// fmt.Println(txt)
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
			// fmt.Printf("posCo:%d, dashIdx:%d\n", posCo, dashIdx)
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
				ticket := make([]int, 0, len(segs))
				for _, seg := range segs {
					num, _ := strconv.Atoi(seg)
					ticket = append(ticket, num)
				}
				tickets = append(tickets, ticket)
			}
			if strings.Contains(txt, "your") {
				scanner.Scan()
				localtext := scanner.Text()
				fmt.Printf("your ticket: %s\n", localtext)
				segs := strings.FieldsFunc(localtext, func(c rune) bool {
					return c == ','
				})
				for _, seg := range segs {
					num, _ := strconv.Atoi(seg)
					yourTicket = append(yourTicket, num)
				}
				// ignore for now
			} else if strings.Contains(txt, "nearby") {
				flag = true
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	invalid := make(map[int]bool)
	for i, ticket := range tickets {
		for _, num := range ticket {
			if !isValid(regs, num) {
				invalid[i] = true
			}
		}
	}
	m, n := len(tickets)-len(invalid), len(tickets[0])
	reverted := make([][]int, n)
	for i := range reverted {
		reverted[i] = make([]int, 0, m)
		for k := 0; k < len(tickets); k++ {
			if !invalid[k] {
				reverted[i] = append(reverted[i], tickets[k][i])
			}
		}
	}
	fields := make([][]string, len(reverted))
	for i := range reverted {
		for k, v := range regs {
			if isValidv2(v, reverted[i]) {
				fields[i] = append(fields[i], k)
				// if strings.Contains(k, "departure") {
				// 	ans *= yourTicket[i]
				// 	fmt.Printf("key: %s match col:%d, val:%d\n", k, i, yourTicket[i])
				// }
			}
		}
		// delete(regs, fields[i])
	}
	// fileds can be sorted based on the length of protential field.
	coded := []int{10, 17, 0, 14, 5, 15}
	ans := 1
	for i := range coded {
		ans *= yourTicket[coded[i]]
	}
	fmt.Println(ans)
	// cols := make([]int, len(reverted))
	// for i := range cols {
	// 	cols[i] = i
	// }
	// fmt.Println(cols)
	// sort.Slice(cols, func(i, j int) bool {
	// 	return len(fields[i]) < len(fields[j])
	// })
	// fmt.Println(cols)
	// fixed := make([]string, len(reverted))
	// for i := range fields {
	// 	if len(fields[i]) == 1 {
	// 		fmt.Printf("col: %d is fixed with field:%s\n", i, fields[i][0])
	// 	} else {
	// 		fmt.Printf("col: %d has multiple (%d) fields: %v\n", i, len(fields[i]), fields[i])
	// 	}
	// }
}

func isValidv2(conds []cond, nums []int) bool {
	// fmt.Printf("cur conds: %v, nums:%v\n", conds, nums)
	for _, num := range nums {
		flag := false
		for _, c := range conds {
			if num >= c.l && num <= c.r {
				flag = true
				break
			}
		}
		if !flag {
			return false
		}
	}
	return true
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
