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
	var ans int
	var cmds []string
	var args []int
	for scanner.Scan() {
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
		txt := scanner.Text()
		segs := strings.Fields(txt)
		arg, _ := strconv.Atoi(segs[1])
		cmds = append(cmds, segs[0])
		args = append(args, arg)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	n := len(cmds)
	set := make(map[string]int, n)
	var i int
	for i < n {
		key := fmt.Sprintf("%d-%s", i, cmds[i])
		// fmt.Printf("key: %s, arg: %d\n", key, args[i])
		set[key]++
		if set[key] > 1 {
			fmt.Println(ans)
			return
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
}
