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
	_, ans := run(p1, p2, true)
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

func play(store map[int]map[Round]bool, p1, p2 []int, gr, round int) (int, bool) {
	if store[gr] == nil {
		store[gr] = make(map[Round]bool)
	}
	fmt.Printf("-- Round %d (Game %d) --\n", round, gr)
	fmt.Printf("Player 1' deck: %v\n", p1)
	fmt.Printf("Player 2' deck: %v\n", p2)
	// fmt.Printf("--debug store[%d]: %v---\n", gr, store[gr])
	k1, k2 := toString(p1), toString(p2)
	cur := Round{k1, k2}
	// rule 1
	if round > 1 {
		if store[gr][cur] {
			fmt.Println("rule 1 matched, play1 wins")
			fmt.Printf("Play1 wins round %d of game %d!\n", round, gr)
			return 0, true
		}
	}
	store[gr][cur] = true
	t1, t2 := p1[0], p2[0]
	p1 = p1[1:]
	p2 = p2[1:]
	fmt.Printf("Player 1' plays: %d\n", t1)
	fmt.Printf("Player 2' plays: %d\n", t2)
	if len(p1) >= t1 && len(p2) >= t2 {
		np1, np2 := make([]int, t1), make([]int, t2)
		copy(np1, p1[:t1])
		copy(np2, p2[:t2])
		// sub-game
		localRound := 1
		for len(np1) != 0 && len(np2) != 0 {
			res, ok := play(store, np1, np2, gr+1, localRound)
			if ok {
				np2 = nil
				continue
			}
			t1, t2 := np1[0], np2[0]
			np1 = np1[1:]
			np2 = np2[1:]
			if res != 0 {
				np2 = append(np2, t2, t1)
				fmt.Printf("Play2 wins round %d of game %d!\n", round, gr+1)
			} else {
				np1 = append(np1, t1, t2)
				fmt.Printf("Play1 wins round %d of game %d!\n", round, gr+1)
			}
			localRound++
		}
		if len(np1) == 0 {
			fmt.Printf("...anyway, back to game %d.\nPlay2 wins round %d of game %d!\n", gr, round, gr+1)
			return 1, false
		}
		fmt.Printf("...anyway, back to game %d.\nPlay1 wins round %d of game %d!\n", gr, round, gr+1)
		return 0, false
	} else {
		if t1 < t2 {
			fmt.Printf("Play2 wins round %d of game %d!\n", round, gr)
			return 1, false
		}
		fmt.Printf("Play1 wins round %d of game %d!\n", round, gr)
		return 0, false
	}
}
