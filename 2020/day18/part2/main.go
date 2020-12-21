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
	var exprs []string
	var ans int64
	for scanner.Scan() {
		txt := scanner.Text()
		exprs = append(exprs, txt)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	for _, expr := range exprs {
		ans += helper(expr)
	}
	fmt.Println(ans)
}

func helper(str string) int64 {
	segs := strings.Fields(str)
	var ans int64
	op := byte('+')
	var st []int64
	var ops []byte
	for _, seg := range segs {
		num, err := strconv.ParseInt(seg, 10, 64)
		if err != nil {
			// fmt.Printf("cur seg: %s, ans: %d\n", seg, ans)
			if len(seg) == 1 {
				op = seg[0]
				if op == '*' {
					// push ans and op in the stack.
					st = append(st, ans)
					ops = append(ops, op)
					ans = 0
					op = byte('+')
				}
			} else {
				// with ()
				var i int
				ln := len(seg)
				if seg[0] == '(' {
					for i < ln {
						if seg[i] == '(' {
							ops = append(ops, op)
							ops = append(ops, '(')
							st = append(st, ans)
							st = append(st, 0)
							// fmt.Printf("cur st: %v, ops: %v\n", st, ops)

							ans = 0
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
					// fmt.Printf("right ) cur num:%d, cur ans: %d, op: %s\n", num, ans, string(op))
					nst := len(st)
					switch op {
					case '+':
						ans += num
					case '*':
						ans *= num
					}
					for i < ln {
						// start to pop
						// in the () they have same priority
						for nst > 0 && ops[nst-1] != '(' {
							preVal := st[nst-1]
							preOp := ops[nst-1]
							ops = ops[:nst-1]
							st = st[:nst-1]
							nst--
							switch preOp {
							case '+':
								ans += preVal
							case '*':
								ans *= preVal
							}
						}
						// pop ( and 0
						ops = ops[:nst-1]
						st = st[:nst-1]
						nst--
						// `+` has higher priority, so update ans
						for nst > 0 && ops[nst-1] == '+' {
							ops = ops[:nst-1]
							preVal := st[nst-1]
							st = st[:nst-1]
							nst--
							ans += preVal
						}
						i++
					}

					// fmt.Printf("after ) we got ans:%d, ops:%v, st:%v\n", ans, ops, st)
				}

			}
		} else {
			switch op {
			case '+':
				ans += num
			}
		}
	}
	nst := len(st)
	// if nst > 0 {
	// 	fmt.Printf("for str: %s, remaings st: %v, and ops:%v\n", str, st, ops)
	// }
	for nst > 0 {
		// multiplication part
		preVal := st[nst-1]
		ops = ops[:nst-1]
		st = st[:nst-1]
		nst--
		ans *= preVal
	}
	// fmt.Printf("expr: %s ===> %d\n", str, ans)
	return ans
}
