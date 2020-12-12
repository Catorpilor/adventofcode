package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type pos struct {
	x, y int
}

var (
	headings = map[int]pos{
		0:   pos{x: 0, y: -1},
		90:  pos{1, 0},
		180: pos{0, 1},
		270: pos{-1, 0},
	}
	dirs = map[byte]pos{
		'N': pos{0, -1},
		'E': pos{1, 0},
		'S': pos{0, 1},
		'W': pos{-1, 0},
	}
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var actions []byte
	var units []int
	for scanner.Scan() {
		txt := scanner.Text()
		actions = append(actions, txt[0])
		num, _ := strconv.Atoi(txt[1:])
		units = append(units, num)
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	wp := pos{x: 10, y: -1}
	ship := pos{}

	for i := range actions {
		act := actions[i]
		unit := units[i]
		switch act {
		case 'F':
			ship.x += wp.x * unit
			ship.y += wp.y * unit
		case 'N', 'S', 'W', 'E':
			wp.x += unit * dirs[act].x
			wp.y += unit * dirs[act].y
		case 'L', 'R':
			wp = rotate(wp, act, unit)
		}
	}
	ans := abs(ship.x) + abs(ship.y)
	fmt.Println(ans)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func rotate(wp pos, act byte, unit int) pos {
	for i := unit / 90; i > 0; i-- {
		if act == 'L' {
			wp = pos{wp.y, -wp.x}
		} else {
			wp = pos{-wp.y, wp.x}
		}
	}
	return wp
}
