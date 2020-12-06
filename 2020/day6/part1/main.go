package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var ans int
	var qst [26]int
	for scanner.Scan() {
		txt := scanner.Text()
		if txt != "" {
			// means the same group
			for i := range txt {
				qst[int(txt[i]-'a')]++
			}
		} else {
			ans += cal(qst)
			for i := range qst {
				qst[i] = 0
			}
		}
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	ans += cal(qst)
	fmt.Println(ans)
}

func cal(g [26]int) int {
	var ans int
	for i := range g {
		if g[i] != 0 {
			ans++
		}
	}
	return ans
}
