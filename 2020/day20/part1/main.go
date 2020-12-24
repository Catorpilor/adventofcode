package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type tile struct {
	id                    int
	matix                 [][]byte
	left, right, up, down stats
}

type stats struct {
	dotCount, pondCount int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	store := make(map[int]*tile)
	var tmp [][]byte
	var id int
	var l0, r0, up0, down0, n int
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			// copy
			local := make([][]byte, len(tmp))
			for i := range tmp {
				local[i] = make([]byte, len(tmp[i]))
				copy(local[i], tmp[i])
			}
			store[id].matix = local
			store[id].left = stats{dotCount: l0, pondCount: n - l0}
			store[id].right = stats{dotCount: r0, pondCount: n - r0}
			store[id].up = stats{dotCount: up0, pondCount: n - up0}
			store[id].down = stats{dotCount: down0, pondCount: n - down0}
			l0, r0, up0, down0 = 0, 0, 0, 0
			continue
		}
		n = len(txt)
		if txt[n-1] == ':' {
			// tile head
			// extract no.
			for i := n - 2; i >= 0; i-- {
				if txt[i] == ' ' {
					id, _ = strconv.Atoi(txt[i+1 : n-1])
					store[id] = &tile{id: id}
					break
				}
			}
			tmp = nil
		} else {
			if len(tmp) == 0 {
				for i := range txt {
					if txt[i] == '.' {
						up0++
					}
				}
			}
			if len(tmp) == n-1 {
				for i := range txt {
					if txt[i] == '.' {
						down0++
					}
				}
			}
			tmp = append(tmp, []byte(txt))
			if txt[0] == '.' {
				l0++
			}
			if txt[n-1] == '.' {
				r0++
			}

		}

	}
	store[id].matix = tmp
	store[id].left = stats{dotCount: l0, pondCount: n - l0}
	store[id].right = stats{dotCount: r0, pondCount: n - r0}
	store[id].up = stats{dotCount: up0, pondCount: n - up0}
	store[id].down = stats{dotCount: down0, pondCount: n - down0}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	for k := range store {
		fmt.Printf("key: %d, left: %v, right: %v, up: %v, bottom: %v\n", k, store[k].left, store[k].right,
			store[k].up, store[k].down)
	}
}
