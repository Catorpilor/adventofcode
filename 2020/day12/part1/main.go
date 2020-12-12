package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var actions []byte
	var units []int
	for scanner.Scan() {
		txt := scanner.Text()
		actions = append(actions, txt[0])
		num, _ := strconv.Atoi(txt[1:])
		units = append(units, num)
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	lead := "E"
	opp := map[string]string{
		"E": "W",
		"W": "E",
		"N": "S",
		"S": "N",
	}
	res := map[string]int{
		"E": 0,
		"W": 0,
		"S": 0,
		"N": 0,
	}
	dirs := map[string][]string{
		"N": []string{"W", "S", "E"},
		"W": []string{"S", "E", "N"},
		"S": []string{"E", "N", "W"},
		"E": []string{"N", "W", "S"},
	}
	for i := range actions {
		act := actions[i]
		unit := units[i]
		fmt.Printf("cur status: %v, lead:%s, act:%s, unit:%d\n", res, lead, string(act), unit)
		switch act {
		case 'F':
			oppd := res[opp[lead]]
			if oppd != 0 {
				if unit > oppd {
					unit -= oppd
					res[opp[lead]] = 0
				} else {
					oppd -= unit
					res[opp[lead]] = oppd
					unit = 0
				}
			}
			res[lead] += unit
		case 'N', 'S', 'W', 'E':
			d := string(act)
			oppd := res[opp[d]]
			if oppd != 0 {
				if unit > oppd {
					unit -= oppd
					res[opp[d]] = 0
				} else {
					oppd -= unit
					res[opp[d]] = oppd
					unit = 0
				}
			}
			res[d] += unit
		case 'L':
			//anti-clockwise
			npos := unit / 90
			lead = dirs[lead][npos-1]
		case 'R':
			npos := unit / 90
			lead = dirs[lead][3-npos]
		}
	}
	var ans int
	for _, v := range res {
		ans += v
	}
	fmt.Println(ans)
}
