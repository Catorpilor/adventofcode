package main

import (
	"bufio"
	"fmt"
	"os"
)

type pos struct {
	x, y, count int
	flip        bool
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
		for i < n {
			var dir string
			if txt[i] != 'e' && txt[i] != 'w' {
				dir = txt[i : i+2]
				i += 2
			} else {
				dir = txt[i : i+1]
				i++
			}
			fmt.Printf("direction: %s\n, cur: %#v\n", dir, cur)
			np := pos{x: cur.x + dirs[dir].x, y: cur.y + dirs[dir].y, count: 1}
			// if np.x == 0 && np.y == 0 {
			// 	panic("go back to where we start")
			// }
			if op, exists := store[np.String()]; exists {
				op.x = np.x
				op.y = np.y
				op.flip = !op.flip
				op.count++
				cur = *op
			} else {
				np.flip = !np.flip
				store[np.String()] = &np
				cur = np
			}
		}
		// fmt.Printf("=====\n store: %v\n", store)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	var ans int
	for _, v := range store {
		if v.count&1 != 0 {
			ans++
			fmt.Printf("tile %s is black\n", v)
		}
	}
	fmt.Println(ans)
}
