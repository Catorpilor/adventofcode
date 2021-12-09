package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	// "strconv"
	"strings"
)

type dot struct {
	boardID  int
	row, col int
	marked   bool
}

type boardStatus struct {
	rowS, colS [5]int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var inputs []string
	prev, cursor, idx := -1, -1, -1
	var boards [][]string
	globalSet := make(map[string][]*dot)
	var tmp []string
	for scanner.Scan() {
		cursor++
		// fmt.Printf("cursor: %d, prev: %d, idx: %d, text: (%s)\n", cursor, prev, idx, scanner.Text())
		if cursor == 0 {
			// means this is the first line
			inputs = strings.Split(scanner.Text(), ",")
			continue
		}
		if scanner.Text() == "" {
			if prev == -1 {
				continue
			}
			local := make([]string, 25)
			copy(local, tmp)
			// fmt.Printf("prepare to append %v to boards.\n", local)
			boards = append(boards, local)
			tmp = nil
			prev = -1
			continue
		}
		if scanner.Text() != "" && prev == -1 {
			prev = cursor
			tmp = make([]string, 0, 25)
			idx++
		}
		nums := strings.Split(scanner.Text(), " ")
		actual := make([]string, 0, len(nums))
		for i := range nums {
			if nums[i] != "" {
				actual = append(actual, nums[i])
				// fmt.Printf("nums[i]=%s at pos: (%d, %d)\n", nums[i], cursor-prev, len(actual)-1)
				globalSet[nums[i]] = append(globalSet[nums[i]], &dot{idx, cursor - prev, len(actual) - 1, false})

			}
		}
		tmp = append(tmp, actual...)
		// your routine here
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	if tmp != nil {
		boards = append(boards, tmp)
	}
	// fmt.Printf("boards: %v, inputs: %v, globalSet: %v\n", boards, inputs, globalSet)
	bs := make([]boardStatus, len(boards))
	wins := make(map[int]bool, len(boards))
	for _, num := range inputs {
		// update board status
		for _, point := range globalSet[num] {
			point.marked = true
			bs[point.boardID].rowS[point.row]++
			if bs[point.boardID].rowS[point.row] == 5 {
				if len(wins) == len(boards)-1 && wins[point.boardID] == false {
					fmt.Printf("bingo, board(%d) wins last\n", point.boardID)
					calculate(globalSet, boards[point.boardID], point.row, 0, num)
					return
				}
				fmt.Printf("bingo, board(%d) wins\n", point.boardID)
				wins[point.boardID] = true
			}
			bs[point.boardID].colS[point.col]++
			if bs[point.boardID].colS[point.col] == 5 {
				if len(wins) == len(boards)-1 && wins[point.boardID] == false{
					fmt.Printf("bingo, board(%d) wins last\n", point.boardID)
					calculate(globalSet, boards[point.boardID], point.col, 1, num)
					return
				}
				fmt.Printf("bingo, board(%d) wins\n", point.boardID)
				wins[point.boardID] = true
			}
		}
	}
}

func calculate(gs map[string][]*dot, board []string, idx, winType int, numS string) {
	fmt.Printf("board: %v wins, idx=%d, winType=%d, numS=%s\n", board, idx, winType, numS)
	var diff, sum int
	for _, b := range board {
		numB, _ := strconv.Atoi(b)
		if gs[b][0].marked == false {
			sum += numB
		}else{
			fmt.Printf("num=%d already marked\n", numB)
		}
		/*
			if winType == 0 {
				if i / 5 == idx {
					diff += numB
				}
			}else {
				if i%5 == idx {
					diff += numB
				}
			}
		*/
	}
	num, _ := strconv.Atoi(numS)
	fmt.Println(num * (sum - diff))
}
