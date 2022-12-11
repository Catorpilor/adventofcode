package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	idx             int
	items           []int
	op              byte
	value           int
	testDivisor     int
	testTrueTarget  int
	testFalseTarget int
	targets         []*monkey
	count           int
}

func (m *monkey) onOp() {
	local := make([]int, len(m.items))
	if len(m.items) != 0 {
		copy(local, m.items)
	}
	for i := range m.items {
		v := m.value
		if v == 0 {
			v = m.items[i]
		}
		switch m.op {
		case '+':
			m.items[i] += v
		case '*':
			m.items[i] *= v
		}
		m.items[i] /= 3
		if m.items[i]%m.testDivisor != 0 {
			// fmt.Printf("throw %d to targets: %d\n", m.items[i],m.testFalseTarget)
			//
			m.targets[1].items = append(m.targets[1].items, m.items[i])
		} else {
			// fmt.Printf("throw %d to target: %d\n", m.items[i], m.testTrueTarget)
			m.targets[0].items = append(m.targets[0].items, m.items[i])
		}
		// fmt.Printf("i: %d, local: %v\n", i, local)
		local = local[1:]
		m.count++
	}
	m.items = local
}

func (m *monkey) String() string {
	return fmt.Sprintf("monkey(%d) with items: %v, divisor(%d), op(%s), value:(%d) targets(true: %d, false: %d), count: %d\n", m.idx, m.items, m.testDivisor, string(m.op), m.value, m.testTrueTarget, m.testFalseTarget, m.count)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var res int
	var monkeys []*monkey
	var blockIdx int
	var cur *monkey
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			blockIdx = 0
			if cur == nil {
				break
			}
			monkeys = append(monkeys, &monkey{
				idx:             cur.idx,
				items:           cur.items,
				op:              cur.op,
				value:           cur.value,
				testDivisor:     cur.testDivisor,
				testTrueTarget:  cur.testTrueTarget,
				testFalseTarget: cur.testFalseTarget,
				targets:         []*monkey{},
				count:           0,
			})
			continue
		}
		switch blockIdx {
		case 0:
			idx := strings.LastIndex(txt, " ")
			v, _ := strconv.Atoi(txt[idx+1 : len(txt)-1])
			cur = &monkey{idx: v}
		case 1:
			idx := strings.Index(txt, ":")
			if idx != -1 {
				strs := strings.Split(txt[idx+2:], ", ")
				for i := range strs {
					v, _ := strconv.Atoi(strs[i])
					cur.items = append(cur.items, v)
				}
			}
		case 2:

			idx := strings.Index(txt, "old ")
			cur.op = txt[idx+4]
			v, _ := strconv.Atoi(txt[idx+6:])
			cur.value = v
		case 3:
			idx := strings.Index(txt, "by ")
			v, _ := strconv.Atoi(txt[idx+3:])
			cur.testDivisor = v
		case 4:
			idx := strings.LastIndex(txt, " ")
			v, _ := strconv.Atoi(txt[idx+1:])
			cur.testTrueTarget = v
		case 5:
			idx := strings.LastIndex(txt, " ")
			v, _ := strconv.Atoi(txt[idx+1:])
			cur.testFalseTarget = v
		}
		blockIdx++

		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if cur != nil {
		monkeys = append(monkeys, cur)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	for i := range monkeys {
		monkeys[i].targets = append(monkeys[i].targets, monkeys[monkeys[i].testTrueTarget], monkeys[monkeys[i].testFalseTarget])
		// fmt.Println(monkeys[i].String())
	}
	for i:=1; i<21; i++ {
		for j := range monkeys {
			monkeys[j].onOp()
		}
	}
	sort.Slice(monkeys, func(i, j int) bool{
		return monkeys[i].count > monkeys[j].count
	})
	res = monkeys[0].count * monkeys[1].count
	fmt.Println(res)
}
