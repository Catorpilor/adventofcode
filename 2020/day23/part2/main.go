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
	for i := 10; i < 100_000_1; i++ {
		prev.next = &list{val: i}
		prev = prev.next
	}
	prev.next = dummy.next
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	// debug(dummy.next, 20)
	head := dummy.next
	cur := head
	pos := make(map[int]*list, 100_000_0)
	pos[cur.val] = cur
	tmp := cur.next
	for tmp != cur {
		pos[tmp.val] = tmp
		tmp = tmp.next
	}
	picks := make([]int, 100_000_1)
	play(pos, head, picks, round)
	cur = pos[1]
	ans := debug(cur.next, 2)
	fmt.Println(ans)
}

func play(store map[int]*list, head *list, picks []int, round int) {
	// n := len(cl)
	cur := head

	for i := 0; i < round; i++ {
		// fmt.Printf("round(%d) \n", i)
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
		dest := choose(store, head, picks, cv)
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

func debug(head *list, n int) int64 {
	cur := head
	ans := int64(1)
	for i := 0; i < n; i++ {
		fmt.Print(cur.val)
		ans *= int64(cur.val)
		cur = cur.next
		if i != n-1 {
			fmt.Print("->")
		}
	}
	fmt.Println("\n-----\n")
	return ans
}

type list struct {
	val  int
	next *list
}

func choose(store map[int]*list, head *list, picks []int, val int) *list {
	wanted := val - 1
	for wanted >= 1 {
		if picks[wanted] == 0 {
			break
		}
		wanted--
	}
	if wanted == 0 {
		for i := 100_000_0; i > 0; i-- {
			if picks[i] == 0 {
				wanted = i
				break
			}
		}
	}
	return store[wanted]
}
