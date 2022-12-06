package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type op struct {
	from, to, amount int
}

var (
	// Define the regular expression pattern.
	pattern = `\d+`

	// Compile the regular expression.
	re = regexp.MustCompile(pattern)
)

type node struct {
	val       byte
	pre, next *node
}

type fifo struct {
	pos *node
}

func (f *fifo) MoveNTo(to *fifo, n int) {
	if n == 0 || f.pos == nil {
		return
	}
	// fmt.Printf("before move: from: %s, to: %s, amount: %d\n", f.String(), to.String(), n)
	for i := 0; i < n; i++ {
		f.pos = f.pos.pre
	}
	head := f.pos.next
	head.pre = to.pos
	to.pos.next = head
	for head.next != nil {
		head = head.next
	}
	to.pos = head
	f.pos.next = nil
	// fmt.Printf("after move: from: %s, to: %s, amount: %d\n", f.String(), to.String(), n)
}

func (f *fifo) Push(n *node) {
	f.pos.next = n
	n.pre = f.pos
	f.pos = n
}

func (f *fifo) String() string {
	var sb bytes.Buffer
	cur := f.pos
	for cur.pre != nil {
		sb.WriteByte(cur.val)
		cur = cur.pre
	}
	return sb.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var sb bytes.Buffer
	var flag bool
	var ops []op
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			flag = true
			continue
		}
		if !flag {
			//
			sb.WriteString(txt)
			sb.WriteByte('\n')
			continue
		}
		matches := re.FindAllString(txt, -1)
		if len(matches) == 3 {
			from, _ := strconv.Atoi(matches[1])
			to, _ := strconv.Atoi(matches[2])
			amount, _ := strconv.Atoi(matches[0])
			ops = append(ops, op{from: from, to: to, amount: amount})
		}
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	sbs := strings.Split(sb.String(), "\n")
	n := len(sbs)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		sbs[i], sbs[j] = sbs[j], sbs[i]
	}
	// get number of racks
	cols := strings.Split(sbs[1], " ")
	var count int
	for _, val := range cols {
		if val != "" {
			count++
		}
	}
	racks := make([]*fifo, count)
	for i := range racks {
		racks[i] = &fifo{pos: &node{}}
	}
	for i := 2; i < n; i++ {
		n2 := len(sbs[i])
		var idx int
		for j := 0; j < n2; j += 4 {
			if sbs[i][j] == '[' {
				racks[idx].Push(&node{val: sbs[i][j+1]})
			}
			idx++
		}
	}
	// fmt.Println(res)
	// fmt.Println(ops)

	// for debug
	// for i := range racks {
	//	fmt.Println(racks[i].String())
	// }
	for _, o := range ops {
		racks[o.from-1].MoveNTo(racks[o.to-1], o.amount)
	}

	//
	res := make([]byte, 0, len(racks))
	for i := range racks {
		str := racks[i].String()
		if len(str) > 0 {
			res = append(res, str[0])
		}
	}
	fmt.Println(string(res))
}
