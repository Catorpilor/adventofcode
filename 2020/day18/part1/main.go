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
<<<<<<< HEAD
	var ans int
	for scanner.Scan() {
		txt := scanner.Text()
		ans += helper(txt)
=======
	var exprs []string
	var ans int64
	for scanner.Scan() {
		txt := scanner.Text()
		// exprs = append(exprs, strings.ReplaceAll(txt, " ", ""))
		exprs = append(exprs, txt)
>>>>>>> aa500e13cd5126fda5cca77df1b36d8d4a553f33
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
<<<<<<< HEAD
	fmt.Println(ans)
}

func helper(s string) int {
	// var preOp byte // preOp stores the previous op of (
	var ans int
	var op byte
	// n := len(s)
	// prevSpc := -1
	segs := strings.Fields(s)
	var leftOpd int
	for _, seg := range segs {
		num, err := strconv.Atoi(seg)
		if err != nil {
			if len(seg) != 1 {
				// with ( or )

			} else {

			}
		}
	}
	// for i:=0; i<n; i++ {
	// 	if s[i] != ' ' {
	// 		continue
	// 	}
	// 	//
	// 	seg := s[prevSpc+1:i]
	// 	if len(seg) > 1 && seg[0] != '(' {

	// 	}

	// }
=======
	for _, expr := range exprs {
		ans += helper(expr)
	}
	fmt.Println(ans)
	// fmt.Println(sumResults(exprs, solve))
}

// func sumResults(equations []string, f func([]rune, int) (int, int)) int {
// 	var sum int
// 	for _, e := range equations {
// 		v, _ := f([]rune(e), 0)
// 		fmt.Printf("expr: %s ==>> %d\n", e, v)
// 		sum += v
// 	}
// 	return sum
// }

// func solve(s []rune, i int) (int, int) {
// 	var value int

// 	var op rune

// 	for ; i < len(s); i++ {
// 		c := s[i]
// 		var newV int
// 		if unicode.IsDigit(c) {
// 			newV = int(c - '0')
// 		} else if c == ')' {
// 			break
// 		} else if c == '(' {
// 			newV, i = solve(s, i+1)
// 		} else {
// 			op = c
// 			continue
// 		}

// 		switch op {
// 		case 0:
// 			value = newV
// 		case '+':
// 			value += newV
// 		case '*':
// 			value *= newV
// 		}

// 	}
// 	return value, i
// }

func helper(str string) int64 {
	segs := strings.Fields(str)
	var ans int64
	op := byte('+')
	var st []int64
	var ops []byte
	for _, seg := range segs {
		num, err := strconv.ParseInt(seg, 10, 64)
		if err != nil {
			// fmt.Printf("cur seg: %s\n", seg)
			if len(seg) == 1 {
				op = seg[0]
			} else {
				// with ()
				var i int
				ln := len(seg)
				if seg[0] == '(' {
					for i < ln {
						if seg[i] == '(' {
							ops = append(ops, op)
							st = append(st, ans)
							ans = 0
							// reset op
							op = byte('+')
							i++
							continue
						}
						break
					}
					num, _ := strconv.ParseInt(seg[i:], 10, 64)
					ans += num
				} else {
					for i < ln {
						if seg[i] != ')' {
							i++
							continue
						}
						break
					}
					num, _ := strconv.ParseInt(seg[:i], 10, 64)
					nst := len(st)
					switch op {
					case '+':
						ans += num
					case '*':
						ans *= num
					}
					for i < ln {
						// start to pop
						preOp := ops[nst-1]
						preVal := st[nst-1]
						ops = ops[:nst-1]
						st = st[:nst-1]
						nst--
						switch preOp {
						case '+':
							ans += preVal
						case '*':
							ans *= preVal
						}
						i++
					}
				}

			}
		} else {
			switch op {
			case '+':
				ans += num
			case '*':
				ans *= num
			}
		}
	}
	fmt.Printf("expr: %s ===> %d\n", str, ans)
>>>>>>> aa500e13cd5126fda5cca77df1b36d8d4a553f33
	return ans
}
