package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var res int
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			continue
		}
		n := len(txt)
		half := n >> 1
		dict := make(map[byte]bool, 52)
		for i := 0; i < half; i++ {
			dict[txt[i]] = true
		}
		for j := half; j < n; j++ {
			if dict[txt[j]] {
				// fmt.Printf("line(%s) identified by %s\n", txt, string(txt[j]))
				if txt[j] >= 'a' && txt[j] <= 'z' {
					// val :=  1 + int(txt[j]-'a')
					// fmt.Printf("string(%s)'s value = %d\n", string(txt[j]), val)
					res += 1 + int(txt[j] - 'a')
				} else {
					val :=  26 + int(txt[j]-'A')
					if val > 26 {
						val += 1
					}
					// fmt.Printf("string(%s)'s value = %d\n", string(txt[j]), val)
					res += val
				}
				break
			}
		}
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(res)
}
