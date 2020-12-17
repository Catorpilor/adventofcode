package main

import (
	"bufio"
	"fmt"
	"os"
)

type pos struct {
	x, y, z, w int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	actives := make(map[pos]bool)
	var row int
	for scanner.Scan() {
		txt := scanner.Text()
		for i, c := range txt {
			if c == '#' {
				actives[pos{row, i, 0, 0}] = true
			}
		}
		row++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	ans := helper(actives, 6)
	fmt.Println(ans)
}

func helper(activies map[pos]bool, round int) int {
	for i := 0; i < round; i++ {
		activies = spin(activies)
	}
	return len(activies)
}

func spin(activies map[pos]bool) map[pos]bool {
	// 26
	affcted := make(map[pos]int, 80)
	for p := range activies {
		for x := p.x - 1; x <= p.x+1; x++ {
			for y := p.y - 1; y <= p.y+1; y++ {
				for z := p.z - 1; z <= p.z+1; z++ {
					for w := p.w - 1; w <= p.w+1; w++ {
						np := pos{x, y, z, w}
						if np == p {
							continue
						}
						affcted[np]++
					}
				}
			}
		}
	}
	res := make(map[pos]bool)
	for p, c := range affcted {
		if activies[p] {
			if c == 2 || c == 3 {
				res[p] = true
			}
		} else {
			if c == 3 {
				res[p] = true
			}
		}
	}
	return res
}
