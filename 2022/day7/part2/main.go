package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type node struct {
	name     string
	size     int
	nType    string
	children []*node
	// parent   *node
}

func LevelWalk(root *node, res *[]int) ([]string, int) {
	if root == nil {
		return nil, 0
	}
	cur := root
	if len(root.children) == 0 {
		// it's a file
		return []string{cur.name}, root.size
	}
	var tmp []string
	for _, child := range root.children {
		names, size := LevelWalk(child, res)
		// fmt.Printf("node: %s(%s) with size: %d, names: %v\n", child.name, child.nType,  size, names)
		for _, name := range names {
			tmp = append(tmp,name)
		}
		if child.nType == "dir" {
			*res = append(*res, size)
		}
		cur.size += size
	}
	// ret := make([]string, 0, len(tmp))
	// for _, name := range tmp {
	// 	ret = append(ret, cur.name + "/" + name)
	// }
	return tmp, cur.size
}

type stack struct{
	paths []string
}

func (s *stack) String() string {
	// fmt.Printf("++debug++, path: %v\n", s.paths)
	pp := strings.Join(s.paths, "/")
	if len(pp) > 2 && pp[:2] == "//" {
		pp = pp[1:]
	}
	return pp
}

func (s *stack) push(path string) {
	s.paths = append(s.paths, path)
}

func (s *stack) pop() {
	n := len(s.paths)
	s.paths = s.paths[:n-1]
}

func (s *stack) top() string {
	n := len(s.paths)
	return s.paths[n-1]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var parent *node
	cache := make(map[string]*node)
	st := &stack{}
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			continue
		}
		args := strings.Split(txt, " ")
		if args[0] == "$" {
			// command
			if len(args) == 3 {
				if args[2] == ".." {
					st.pop()
					continue
				}
				st.push(args[2])
				path := st.String()
				// fmt.Printf("push %s , current path: %s\n", args[2], path)
				if _, exists := cache[path]; !exists {
					cache[path] = &node{name: path, nType: "dir"}
				}
				parent = cache[path]
			}
		} else {
			path := st.String()
			path += "/" + args[1]
			if len(path) > 2 && path[:2] == "//" {
				path = path[1:]
			}
			if _, exists := cache[path]; !exists {
				cur := &node{name: path, nType: "dir"}
				if args[0] != "dir" {
					size, _ := strconv.Atoi(args[0])
					cur.size = size
					cur.nType = "file"
				}
				parent.children = append(parent.children, cur)
				cache[path] = cur
			}
		}
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	root := cache["/"]
	var res []int
	_, totalSize := LevelWalk(root, &res)
	res = append(res, totalSize)
	remaining := 70000000 - totalSize
	wants := 30000000 - remaining
	sort.Ints(res)
	i := sort.SearchInts(res, wants)
	fmt.Println(res[i])
}
