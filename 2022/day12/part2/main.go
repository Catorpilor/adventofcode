package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var matrix [][]byte
	var row int
	var starts []coord
	var end coord
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			continue
		}
		curRow := make([]byte, 0, len(txt))
		for i := range txt {
			b := txt[i]
			if txt[i] == 'a' || txt[i] == 'S' {
				starts = append(starts, coord{row, i})
				b = 'a'
			} else if txt[i] == 'E' {
				end = coord{row, i}
				b = 'z'
			}
			curRow = append(curRow, b)
		}
		matrix = append(matrix, curRow)
		row++
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	// fmt.Printf("start(%d, %d), end(%d, %d), m:%d, n:%d\n", sx, sy, ex, ey, len(matrix), len(matrix[0]))
	ret := math.MaxInt
	for _, s := range starts {
		q := &queue{}
		q.Push(item{0, s})
		res := bfs(matrix, s, end, q)
		// fmt.Printf("calculate start: (%d,%d), end(%d, %d) got %d\n", s.x, s.y, end.x, end.y, res)
		if res < ret {
			ret = res
		}
	}
	fmt.Println(ret)
}

type coord struct{ x, y int }

func (c coord) String() string {
	return fmt.Sprintf("%d-%d", c.x, c.y)
}

type item struct {
	depth int
	c     coord
}

type queue struct {
	items []item
}

func (q *queue) Push(o item) {
	q.items = append(q.items, o)
}

func (q *queue) Front() item {
	return q.items[0]
}

func (q *queue) Pop() {
	if len(q.items) > 0 {
		q.items = q.items[1:]
	}
}

func (q *queue) IsEmpty() bool {
	return len(q.items) == 0
}

func bfs(matrix [][]byte, start, end coord, q *queue) int {
	dirs := []int{0, 1, 0, -1, 0}
	m := len(matrix)
	n := len(matrix[0])
	ret := math.MaxInt
	seen := map[coord]bool{}
	for !q.IsEmpty() {
		it := q.Front()
		q.Pop()
		ori := matrix[it.c.x][it.c.y]
		if it.c.x == end.x && it.c.y == end.y {
			// fmt.Println(it.depth)
			return it.depth
		}
		if seen[it.c]{
			continue
		}
		// fmt.Printf("cur (%s) depth: %d, coord: (%d, %d),visited: %t\n", string(ori), it.depth, it.c.x, it.c.y, visited[it.c.String()])
		// if visited[it.c.String()] {
		// 	continue
		// }
		// fmt.Printf("cur (%s) depth: %d, coord: (%d, %d)\n", string(ori), it.depth,  it.c.x, it.c.y)
		// visited[it.c.String()] = true
		seen[it.c] = true
		for i := 0; i < 4; i++ {
			nx, ny := it.c.x+dirs[i], it.c.y+dirs[i+1]
			if nx < 0 || ny < 0 || nx >= m || ny >= n {
				continue
			}
			nc := coord{nx, ny}
			if int(matrix[nx][ny]) <= int(ori)+1 {
				q.Push(item{it.depth + 1, nc})
			}
		}
	}
	return ret
}
