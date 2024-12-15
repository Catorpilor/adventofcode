package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type dir int

const (
	POINT dir = iota
	R
	L
	U
	D
	UL
	UR
	DL
	DR
)

var (
	dirs   = [9]int{-1, 0, 1, 1, 0, -1, -1, 1, -1}
	wanted = "XMAS"
)

// search searches the matrix from position(x,y) see if there is a match
// it returns the total matches from (x,y)
func search(b [][]byte, x, y int) int {
	var res int
	m := len(b)
	n := len(b[0])
	for i := 0; i < len(dirs)-1; i++ {
		matched := 1
		for j := 1; j < 4; j++ {
			nx, ny := x+j*dirs[i], y+j*dirs[i+1]
			if nx < 0 || nx >= m || ny < 0 || ny >= n {
				break
			}
			if b[nx][ny] == wanted[j] {
				matched++
				continue
			}
		}
		if matched == len(wanted) {
			res++
		}
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var res int
	matrix := make([][]string, 0, 255)
	var flag bool
	link := make(map[string]map[string]bool)
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			flag = true
			continue
		}
		if !flag {
			//fmt.Println(txt)
			ls := strings.Split(txt, "|")
			if _, exists := link[ls[0]]; !exists {
				link[ls[0]] = make(map[string]bool)
			}
			link[ls[0]][ls[1]] = true
		} else {
			update := strings.Split(txt, ",")
			matrix = append(matrix, update)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	for _, st := range matrix {
		if solve(link, st) {
			n := len(st)
			a, _ := strconv.Atoi(st[n/2])
			res += a
		}
	}
	fmt.Println(res)
}

func solve(link map[string]map[string]bool, input []string) bool {
	pageIndex := make(map[string]int)
	for i := range input {
		pageIndex[input[i]] = i
	}
	for k, v := range link {
		if _, exists := pageIndex[k]; !exists {
			continue
		}
		for kk := range v {
			if pk, exists := pageIndex[kk]; exists {
				if pageIndex[k] > pk {
					return false
				}
			}
		}
	}
	return true
}
