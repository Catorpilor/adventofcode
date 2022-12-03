package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var res, idx int
	dict := make(map[byte]int, 52)
	for scanner.Scan() {
		txt := scanner.Text()
		// fmt.Printf("scan txt(%s) and curr idx=%d\n", txt, idx%3)
		if txt == "" {
			continue
		}
		n := len(txt)
		for i := 0; i < n; i++ {
			if v, exists := dict[txt[i]]; !exists {
				if idx == 0 {
					dict[txt[i]] = 1
				}
			} else {
				if v < idx+1 {
					dict[txt[i]] += 1
				}
			}
		}
		for k, v := range dict {
			if v < idx+1 {
				delete(dict, k)
			}
		}
		// fmt.Printf("cur idx:%d, map:%v\n", idx, dict)
		idx++
		if idx == 3 {
			// calculate res
			idx = 0
			// fmt.Println(dict)
			for k, v := range dict {
				if v == 3 {
					// fmt.Printf("cur block we got %s\n", string(k))
					if k >= 'a' && k <= 'z' {
						res += 1 + int(k-'a')
					} else {
						val := 26 + int(k-'A')
						if val > 26 {
							val += 1
						}
						res += val
					}
				}
			}
			// clear map
			dict = make(map[byte]int, 52)
		}
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(res)
}
