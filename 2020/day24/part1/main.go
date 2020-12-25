package main

import (
	"bufio"
	"fmt"
	"os"
)

type pos struct {
	x, y int
	flip bool
}

func (p pos) String() string {
	return fmt.Sprint(p.x, ",", p.y)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	store := make(map[string]*pos)
	st := pos{x: 0, y: 0, flip: false}
	store[st.String()] = &st
	dirs := map[string]pos{
		"e":  pos{x: 2, y: 0},
		"se": pos{x: 1, y: -1},
		"sw": pos{x: -1, y: -1},
		"w":  pos{x: -2, y: 0},
		"nw": pos{x: -1, y: 1},
		"ne": pos{x: 1, y: 1},
	}
	for scanner.Scan() {
		txt := scanner.Text()
		fmt.Printf("processing txt:%s\n", txt)
		n := len(txt)
		var i int
		cur := *(store["0,0"])
		var dir string
		for i < n {
			if txt[i] != 'e' && txt[i] != 'w' {
				dir = txt[i : i+2]
				i += 2
			} else {
				dir = txt[i : i+1]
				i++
			}
			np := pos{x: cur.x + dirs[dir].x, y: cur.y + dirs[dir].y}
			fmt.Printf("direction: %s\n, prev: %#v\n", dir, cur)
			// if i >= n {
			// 	// flip this
			// }
			if op, exists := store[np.String()]; exists {
				cur = *op
			} else {
				store[np.String()] = &np
				cur = np
			}
			if i >= n {
				// flip np
				store[cur.String()].flip = !store[cur.String()].flip
			}
			fmt.Printf("cur: %#v\n", cur)
		}
		// fmt.Printf("=====\n store: %v\n", store)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	var ans int
	for _, v := range store {
		if v.flip {
			ans++
		}
	}
	fmt.Println(ans)
}
