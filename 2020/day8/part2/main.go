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
	// var ans int
	var cmds []string
	var args, jmps, nops []int
	for scanner.Scan() {
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
		txt := scanner.Text()
		segs := strings.Fields(txt)
		arg, _ := strconv.Atoi(segs[1])
		cn := len(cmds)
		cmds = append(cmds, segs[0])
		if segs[0] == "jmp" {
			jmps = append(jmps, cn)
		} else if segs[0] == "nop" {
			nops = append(nops, cn)
		}
		args = append(args, arg)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	// n := len(cmds)
	// fmt.Println(len(jmps), len(nops))
	// var negj int
	for i := range jmps {
		// if args[jmps[i]] < 0 && jmps[i-1] >= jmps[i]+args[jmps[i]] {
		// 	fmt.Printf("cur op: %d, cur arg: %d, prev op: %d, prev arg: %d\n", jmps[i], args[jmps[i]],
		// 		jmps[i-1], args[jmps[i-1]])
		// 	negj++
		// }
		cop := jmps[i]
		// if args[cop] < 0 {
		cmds[cop] = "nop"
		res := solve(cmds, args)
		if res != -1 {
			fmt.Println(res)
			return
		}
		cmds[cop] = "jmp"
		// }
		// fmt.Printf("cur op: %d, cur arg: %d, next op: %d\n", jmps[i], args[jmps[i]], jmps[i]+args[jmps[i]])
	}
	for i := range nops {
		cop := nops[i]
		if args[cop] != 0 {
			cmds[cop] = "jmp"
			res := solve(cmds, args)
			if res != -1 {
				fmt.Println(res)
				return
			}
			cmds[cop] = "nop"
		}
	}
	// fmt.Printf("negj: %d\n", negj)

}

func solve(cmds []string, args []int) int {
	var ans int
	n := len(cmds)
	set := make(map[string]int, n)
	// for i := range jmps {
	// 	cur := args[jmps[i]]
	// 	if cur > 0 {
	// 		continue
	// 	}
	// 	// cur jmp back
	// 	prev := args[jmps[i-1]]

	// }
	var i int
	for i < n {
		key := fmt.Sprintf("%d-%s", i, cmds[i])
		// fmt.Printf("key: %s, arg: %d\n", key, args[i])
		set[key]++
		if set[key] > 1 {
			return -1
		}
		switch cmds[i] {
		case "acc":
			ans += args[i]
			i++
		case "jmp":
			i += args[i]
		default:
			i++
		}
	}
	return ans
}
