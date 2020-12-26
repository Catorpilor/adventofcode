package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pos struct {
	x, y int
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
	store := make(map[string]bool)
	// store["0,0"] = &pos{x: 0, y: 0}
	st := pos{x: 0, y: 0}
	for scanner.Scan() {
		txt := scanner.Text()
		fmt.Printf("processing txt:%s\n", txt)
		n := len(txt)
		var i int
		cur := st
		var dir string
		for i < n {
			if txt[i] != 'e' && txt[i] != 'w' {
				dir = txt[i : i+2]
				i += 2
			} else {
				dir = txt[i : i+1]
				i++
			}
			// if np.x == 0 && np.y == 0 {
			// 	panic("go back to where we start")
			cur.x += dirs[dir].x
			cur.y += dirs[dir].y
			if i >= n {
				// flip np
				store[cur.String()] = !store[cur.String()]
			}
		}
		// fmt.Printf("=====\n store: %v\n", store)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	_ = process(store, 100)
}

func toPos(s string) pos {
	segs := strings.FieldsFunc(s, func(c rune) bool {
		return c == ','
	})
	var ret pos
	ret.x, _ = strconv.Atoi(segs[0])
	ret.y, _ = strconv.Atoi(segs[1])
	return ret

}

func dailyFlip(blackTiles map[string]bool) map[string]bool {
	blackNeighbours := make(map[string]int)

	deltas := make([]pos, 0, len(dirs))
	for _, v := range dirs {
		deltas = append(deltas, v)
	}
	for p, b := range blackTiles {
		if !b {
			continue
		}
		cur := toPos(p)
		for _, d := range deltas {
			vp := pos{cur.x + d.x, cur.y + d.y}
			blackNeighbours[vp.String()]++
		}
	}

	newBlackTiles := make(map[string]bool)

	for p, count := range blackNeighbours {
		previousColour := blackTiles[p]

		if previousColour {
			if !(count == 0 || count > 2) {
				newBlackTiles[p] = true
			}
		} else if count == 2 {
			newBlackTiles[p] = true
		}
	}

	return newBlackTiles
}

func process(store map[string]bool, n int) int {
	var ans int
	for i := 0; i < n; i++ {
		ans = 0
		store = dailyFlip(store)
		for _, v := range store {
			if v {
				ans++
			}
		}
		fmt.Printf("day(%d): number of black: %d\n", i+1, ans)
	}
	return ans
}
