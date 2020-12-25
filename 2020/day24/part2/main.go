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

var (
	dirs = map[string]pos{
		"e":  pos{x: 2, y: 0},
		"se": pos{x: 1, y: -1},
		"sw": pos{x: -1, y: -1},
		"w":  pos{x: -2, y: 0},
		"nw": pos{x: -1, y: 1},
		"ne": pos{x: 1, y: 1},
	}
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var cmds [][]string
	for scanner.Scan() {
		txt := scanner.Text()
		n := len(txt)
		var i int
		var tmp []string
		var dir string
		for i < n {
			if txt[i] != 'e' && txt[i] != 'w' {
				dir = txt[i : i+2]
				i += 2
			} else {
				dir = txt[i : i+1]
				i++
			}
			tmp = append(tmp, dir)
		}
		cmds = append(cmds, tmp)
		// fmt.Printf("=====\n store: %v\n", store)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	_ = process(cmds, 100)
}

func process(cmds [][]string, n int) int {
	store := make(map[string]*pos)

	st := pos{x: 0, y: 0, flip: false}
	store[st.String()] = &st
	deltas := make([]pos, 0, len(dirs))
	for _, v := range dirs {
		deltas = append(deltas, v)
	}
	for i := 0; i < n; i++ {
		cur := *(store["0,0"])
		var needsFlip []string
		nbrs := make(map[string]int)
		for _, dd := range cmds {
			n := len(dd)
			for i, d := range dd {
				np := pos{x: cur.x + dirs[d].x, y: cur.y + dirs[d].y}
				// if i >= n {
				// 	// flip this
				// }
				if op, exists := store[np.String()]; exists {
					cur = *op
				} else {
					store[np.String()] = &np
					cur = np
				}
				if i == n-1 {
					// flip np
					needsFlip = append(needsFlip, cur.String())
					// store[cur.String()].flip = !store[cur.String()].flip
					// get neighbor's colors
					for _, delta := range deltas {
						np := pos{x: cur.x + delta.x, y: cur.y + delta.y}
						if old, exists := store[np.String()]; exists {
							if old.flip {
								nbrs[cur.String()]++
							}
						}
					}
				}
			}
		}
		fmt.Printf("day(%d) needFlip: %v, nbrs: %v\n", i+1, needsFlip, nbrs)
		for _, p := range needsFlip {
			if store[p].flip {
				if nbrs[p] == 0 || nbrs[p] > 2 {
					store[p].flip = false
				}
			} else {
				if nbrs[p] == 2 {
					store[p].flip = true
				}
			}
		}
		var ans int
		for _, v := range store {
			if v.flip {
				ans++
			}
		}
		fmt.Printf("day(%d): black tiles: %d\n", i+1, ans)
	}
	return 0
}
