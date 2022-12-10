package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x, y int
}

func (c *coord) String() string {
	return fmt.Sprintf("%d-%d", c.x, c.y)
}

type knot struct {
	c     *coord
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (k *knot) Move(steps int, dir *coord, visited map[string]bool, restKnots ...*knot) {
	moved := dir.String()
	for steps > 0 {
		prevKnot := k
		for i := 0; i < 9; i++ {
			cur := restKnots[i]
			cdx, cdy := k.c.x-prevKnot.c.x, k.c.y-prevKnot.c.y
			k.c.x += dir.x
			k.c.y += dir.y
			if cdx == 0 && cdy == 0 {
				break
			}
			if cdx == 1 && cdy == 1 {
				if moved == "1-0" || moved == "0-1" {
					cur.c.x += 1
					cur.c.y += 1
				}
			}
		}
		steps--
	}
}

func nextP(hp, tp int) int {
	d := hp - tp
	diff := d
	if abs(d) == 2 {
		diff /= 2
	}
	return tp + diff
}

func move(a, b *coord) {
	d := max(abs(b.x - a.x), abs(b.y - a.y))
	if d > 1 {
		b.x = nextP(a.x, b.x)
		b.y = nextP(a.y, b.y)
	}
}

func onOperate(steps int, dirs *coord, visited map[string]bool, knots ...*knot) {
	dx, dy := dirs.x, dirs.y
	for i:=0; i<steps; i++ {
		knots[0].c.x += dx
		knots[0].c.y += dy
		for j := 1; j<=9; j++ {
			move(knots[j-1].c, knots[j].c)
		}
		visited[knots[9].c.String()] = true
	}
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	rope := make([]*knot, 10)
	for i := range rope {
		rope[i] = &knot{c: &coord{}}
	}
	dirs := map[string]*coord{
		"R": &coord{1, 0},
		"U": &coord{0, 1},
		"L": &coord{-1, 0},
		"D": &coord{0, -1},
	}
	visited := map[string]bool{
		"0-0": true,
	}
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			continue
		}
		strs := strings.Split(txt, " ")
		steps, _ := strconv.Atoi(strs[1])
		//  head.Move(steps, dirs[strs[0]], visited, rope[1:]...)
		onOperate(steps, dirs[strs[0]], visited, rope...)

		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(len(visited))
}
