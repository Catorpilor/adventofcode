package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var ans, pn int
	var qst [26]int
	for scanner.Scan() {
		txt := scanner.Text()
		if txt != "" {
			pn++
			// means the same group
			for i := range txt {
				qst[int(txt[i]-'a')]++
			}
		} else {
			ans += cal(qst, pn)
			for i := range qst {
				qst[i] = 0
			}
			pn = 0
		}
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	ans += cal(qst, pn)
	fmt.Println(ans)
}

func cal(g [26]int, pn int) int {
	// fmt.Printf("%d people in group\n", pn)
	var ans int
	for i := range g {
		if g[i] >= pn {
			ans++
		}
	}
	return ans
}
