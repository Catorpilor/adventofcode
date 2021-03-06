package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	sp1  = "bags contain"
	nsp1 = 12
	sp2  = "no other bags"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// var ans int
	// type bagSpec struct {
	// 	capacity int
	// 	// Type     string
	// }
	adj := make(map[string]map[string]int)
	for scanner.Scan() {
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
		// light red bags contain 1 bright white bag, 2 muted yellow bags.
		txt := scanner.Text()
		idx1 := strings.Index(txt, sp1)
		// idx1-1 here just trim the ' ' in the end.
		pb := txt[:idx1-1]
		// if adj[pb] == nil {
		// 	adj[pb] = make(map[string]int)
		// }
		remains := txt[idx1+nsp1:]
		if !strings.Contains(remains, sp2) {
			// no bags
			specs := strings.FieldsFunc(remains, func(c rune) bool {
				return c == ','
			})
			for _, spec := range specs {
				// fmt.Println(spec)
				spec = spec[1:]
				ci := strings.Index(spec, " ")
				// fmt.Println(ci)
				count, _ := strconv.Atoi(spec[:ci])
				li := strings.LastIndex(spec, " ")
				// fmt.Printf("spec[ci+1:]:%s, li:%d\n", spec[ci+1:], li)
				bagtype := spec[ci+1 : li]
				// fmt.Printf("bag type: %s$\n", bagtype)
				if adj[bagtype] == nil {
					adj[bagtype] = make(map[string]int)
				}
				adj[bagtype][pb] = count
			}
		}
	}
	// fmt.Println(adj)
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	set := make(map[string]bool)
	keys := []string{"shiny gold"}
	var i int
	for i < len(keys) {
		key := keys[i]
		// fmt.Printf("cur i:%d, checking key: %s, adj[key]:%v\n", i, key, adj[key])
		for k, v := range adj[key] {
			// fmt.Printf("adj[key] got k:%s, v:%d\n", k, v)
			if v >= 1 {
				set[k] = true
				keys = append(keys, k)
				// fmt.Printf("cur keys: %v\n", keys)
			}
		}
		i++
	}

	fmt.Println(len(set))
}
