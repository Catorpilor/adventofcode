package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type pick struct {
	next, prev *pick
	name       string
	value      int
}

var (
	rock, paper, scissors *pick
)

func main() {
	rock = &pick{value: 1, name: "rock"}
	paper = &pick{value: 2, name: "paper"}
	scissors = &pick{value: 3, name: "scissors"}
	rock.prev = paper
	rock.next = scissors
	paper.prev = scissors
	paper.next = rock
	scissors.prev = rock
	scissors.next = paper
	cat := map[string]*pick{
		"A": rock,
		"B": paper,
		"C": scissors,
		"X": rock,
		"Y": paper,
		"Z": scissors,
	}
	scanner := bufio.NewScanner(os.Stdin)
	var res int
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			continue
		}
		choices := strings.Split(txt, " ")
		a, b := cat[choices[0]], cat[choices[1]]
		// fmt.Printf("a choose: %s, b choose %s\n", a.name, b.name)
		if a == b {
			// fmt.Println("it's a draw")
			res += 3
		} else if a.prev == b {
			// fmt.Println("b win")
			res += 6
		}
		res += b.value
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(res)
}
