package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rule struct {
	id       int
	isEnd    bool
	val      string
	children [][]int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	flag := false
	var strs []string
	store := make(map[int]*rule)
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			flag = true
			continue
		}
		if flag {
			strs = append(strs, txt)
		} else {
			comIdx := strings.Index(txt, ":")
			id, _ := strconv.Atoi(txt[:comIdx])
			store[id] = &rule{id: id}
			if idx := strings.Index(txt, `"`); idx != -1 {
				for i := idx + 1; i < len(txt); i++ {
					if txt[i] != '"' {
						continue
					}
					store[id].isEnd = true
					store[id].val = txt[idx+1 : i]
					break
				}
			} else {
				segs := strings.Fields(txt[comIdx+2:])
				// fmt.Printf("id: %d, segs: %v\n", id, segs)
				var tmp []int
				for _, seg := range segs {
					num, err := strconv.Atoi(seg)
					if err != nil {
						local := make([]int, len(tmp))
						// fmt.Printf("|, num: %d, cur tmp: %v\n", num, tmp)
						copy(local, tmp)
						store[id].children = append(store[id].children, local)
						tmp = nil
						continue
					}
					tmp = append(tmp, num)
				}
				local := make([]int, len(tmp))
				copy(local, tmp)
				store[id].children = append(store[id].children, local)
			}
		}
	}
	store[8] = &rule{id: 8, children: [][]int{[]int{42}, []int{42, 8}}}
	store[11] = &rule{id: 11, children: [][]int{[]int{42, 31}, []int{42, 11, 31}}}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(helper(strs, store))
}

func helper(strs []string, rules map[int]*rule) int {
	var ans int
	for _, s := range strs {
		if validate([]byte(s), 0, 0, nil, rules) {
			ans++
		}
	}
	return ans
}

func validate(sb []byte, rid, idx int, next []int, rules map[int]*rule) bool {
	r := rules[rid]
	fmt.Printf("s: %s, rid: %d, rule: %v, next: %v, idx: %d\n", string(sb), rid, *r, next, idx)
	if r.isEnd {
		if r.val != string(sb[idx]) {
			return false
		}
		if len(next) == 0 {
			return idx == len(sb)-1
		} else if idx+1 >= len(sb) {
			return false
		} else {
			return validate(sb, next[0], idx+1, next[1:], rules)
		}
	}
	toAdd := r.children[0][1:]
	nn := make([]int, len(next)+len(toAdd))
	copy(nn, toAdd)
	copy(nn[len(toAdd):], next)
	if validate(sb, r.children[0][0], idx, nn, rules) {
		return true
	}
	if len(r.children) > 1 {
		toAdd := r.children[1][1:]
		nn := make([]int, len(next)+len(toAdd))
		copy(nn, toAdd)
		copy(nn[len(toAdd):], next)
		if validate(sb, r.children[1][0], idx, nn, rules) {
			return true
		}
	}
	return false
}
