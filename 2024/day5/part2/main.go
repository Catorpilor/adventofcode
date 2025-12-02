package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
