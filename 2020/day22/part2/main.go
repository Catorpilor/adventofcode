package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Round struct {
	p1, p2 string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var group [2][]int
	var idx int
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			idx++
			continue
		}
		if strings.Contains(txt, ":") {
			continue
		} else {
			num, _ := strconv.Atoi(txt)
			group[idx] = append(group[idx], num)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	p1, p2 := group[0], group[1]
	_, ans := play(p1, p2)
	// _, ans := run(p1, p2, true)
	fmt.Println(ans)
	// fmt.Println(len(p1) + len(p2))
	// store := make(map[int]map[Round]bool)
	// // res := play(store, p1, p2, 1, 1)
	// round := 1
	// gr := 1
	// for len(p1) != 0 && len(p2) != 0 {
	// 	res, ok := play(store, p1, p2, gr, round)
	// 	if ok {
	// 		p2 = nil
	// 		continue
	// 	}
	// 	t1, t2 := p1[0], p2[0]
	// 	p1 = p1[1:]
	// 	p2 = p2[1:]
	// 	if res != 0 {
	// 		p2 = append(p2, t2, t1)
	// 	} else {
	// 		p1 = append(p1, t1, t2)
	// 	}
	// 	round++
	// }
	// res := p1
	// if len(p2) != 0 {
	// 	res = p2
	// }
	// var ans int
	// n := len(res)
	// fmt.Println(len(res))
	// for i := n - 1; i >= 0; i-- {
	// 	ans += res[i] * (n - i)
	// }
	// fmt.Println(ans)
}
func run(deck1, deck2 []int, rec bool) (int, int) {
	seen := map[string]struct{}{}

	var winner int
	for len(deck1) > 0 && len(deck2) > 0 {
		winner = 1
		if _, ok := seen[fmt.Sprint(deck1, deck2)]; rec && ok {
			break
		}
		seen[fmt.Sprint(deck1, deck2)] = struct{}{}

		card1, card2 := deck1[0], deck2[0]
		deck1, deck2 = deck1[1:], deck2[1:]

		if card2 > card1 {
			winner = 2
		}

		if rec && len(deck1) >= card1 && len(deck2) >= card2 {
			copy1, copy2 := append([]int(nil), deck1[:card1]...), append([]int(nil), deck2[:card2]...)
			winner, _ = run(copy1, copy2, rec)
		}

		if winner == 1 {
			deck1 = append(deck1, card1, card2)
		} else {
			deck2 = append(deck2, card2, card1)
		}
	}

	winDeck := deck1
	if winner == 2 {
		winDeck = deck2
	}

	score := 0
	for i, c := range winDeck {
		score += c * (len(winDeck) - i)
	}
	return winner, score
}

func toString(a []int) string {
	var sb strings.Builder
	n := len(a)
	for i, num := range a {
		sb.WriteString(strconv.Itoa(num))
		if i != n-1 {
			sb.WriteByte(',')
		}
	}
	return sb.String()
}

func play(p1, p2 []int) (int, int) {
	seen := make(map[string]bool)
	var winner int
	for len(p1) > 0 && len(p2) > 0 {
		winner = 1
		ck := fmt.Sprint(p1, p2)
		if seen[ck] {
			break
		}
		seen[ck] = true
		t1, t2 := p1[0], p2[0]
		p1 = p1[1:]
		p2 = p2[1:]
		if t2 > t1 {
			winner = 2
		}
		if len(p1) >= t1 && len(p2) >= t2 {
			np1, np2 := make([]int, t1), make([]int, t2)
			copy(np1, p1[:t1])
			copy(np2, p2[:t2])
			winner, _ = play(np1, np2)
		}
		if winner == 1 {
			p1 = append(p1, t1, t2)
		} else {
			p2 = append(p2, t2, t1)
		}
	}
	res := p1
	if winner == 2 {
		res = p2
	}
	n := len(res)
	var ans int
	for i := n - 1; i >= 0; i-- {
		ans += res[i] * (n - i)
	}
	return winner, ans
}
