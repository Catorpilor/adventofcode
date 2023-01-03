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

func (c *coord) Move(steps int, dir, tail *coord, visited map[string]bool) {
	for steps > 0 {
		cdx, cdy := c.x-tail.x, c.y-tail.y
		c.x += dir.x
		c.y += dir.y
		steps--
		moved := dir.String()
		if cdx == 1 && cdy == 1 {
			if moved == "1-0" || moved == "0-1" {
				tail.x += cdx
				tail.y += cdy
				visited[tail.String()] = true
			}
		}else if cdx == 1 && cdy == -1 {
			if moved == "1-0" || moved == "0--1" {
				tail.x += cdx
				tail.y += cdy
				visited[tail.String()] = true
			}
		}else if cdx == -1 && cdy == 1 {
			if moved == "-1-0" || moved == "0-1" {
				tail.x += cdx
				tail.y += cdy
				visited[tail.String()] = true
			}
		}else if cdx == -1 && cdy == -1 {
			if moved == "-1-0" || moved == "0--1" {
				tail.x += cdx
				tail.y += cdy
				visited[tail.String()] = true
			}
		}else if cdx == 1 && cdy == 0 {
			if moved == "1-0" {
				tail.x += cdx
				tail.y += cdy
				visited[tail.String()] = true
			}
		}else if cdx == -1 && cdy == 0 {
			if moved == "-1-0" {
				tail.x += cdx
				tail.y += cdy
				visited[tail.String()] = true
			}
		}else if cdy == 1 && cdx == 0 {
			if moved == "0-1" {
				tail.x += cdx
				tail.y += cdy
				visited[tail.String()] = true
			}
		}else if cdy == -1 && cdx == 0 {
			if moved == "0--1" {
				tail.x += cdx
				tail.y += cdy
				visited[tail.String()] = true
			}
		}
	}
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	head, tail := &coord{}, &coord{}
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
		head.Move(steps, dirs[strs[0]], tail, visited)
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(len(visited))
}
