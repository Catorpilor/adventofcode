package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	round int
)

func main() {
	flag.IntVar(&round, "round", 10, "define the number of moves")
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	dummy := &list{}
	prev := dummy
	for scanner.Scan() {
		txt := scanner.Text()
		for i := range txt {
			num, _ := strconv.Atoi(txt[i : i+1])
			prev.next = &list{val: num}
			prev = prev.next
		}
	}
	prev.next = dummy.next
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	head := dummy.next
	picks := make([]int, 10)
	play(head, picks, round)
	cur := dummy.next
	for cur.val != 1 {
		cur = cur.next
	}
	debug(cur.next, 8)
}

func play(head *list, picks []int, round int) {
	// n := len(cl)
	cur := head
	// pos := make(map[int]int, len(cl))
	for i := 0; i < round; i++ {
		// cv := cl[cur%n]
		cv := cur.val
		// pick 3 cups
		p1, p2, p3 := cur.next, cur.next.next, cur.next.next.next
		// fmt.Printf("cur: %d, picks: (%d, %d, %d)\n", cur.val, p1.val, p2.val, p3.val)
		// debug(cur)
		picks[p1.val]++
		picks[p2.val]++
		picks[p3.val]++
		cur.next = p3.next
		dest := choose(head, picks, cv)
		tmp := dest.next
		dest.next = p1
		p3.next = tmp
		picks[p1.val]--
		picks[p2.val]--
		picks[p3.val]--
		// debug(cur)
		cur = cur.next
	}
}

func debug(head *list, n int) {
	cur := head
	for i := 0; i < n; i++ {
		fmt.Print(cur.val)
		cur = cur.next
	}
	fmt.Println("\n-----")
}

type list struct {
	val  int
	next *list
}

func choose(head *list, picks []int, val int) *list {
	wanted := val - 1
	for wanted >= 1 {
		if picks[wanted] == 0 {
			break
		}
		wanted--
	}
	if wanted == 0 {
		for i := 9; i > 0; i-- {
			if picks[i] == 0 {
				wanted = i
				break
			}
		}
	}
	for cur := head; ; cur = cur.next {
		if cur.val == wanted {
			return cur
		}
	}
}
