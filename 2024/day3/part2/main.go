package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	pattern     = `mul\((\d+),(\d+)\)`
	doPattern   = `do\(\)`
	dontPattern = `don't\(\)'`
	cmdPattern  = `(?:do\(\)|don't\(\)|mul\(\d+,\d+\))`
)

var (
	re     = regexp.MustCompile(pattern)
	doRe   = regexp.MustCompile(doPattern)
	dontRe = regexp.MustCompile(dontPattern)
	cmdsRe = regexp.MustCompile(cmdPattern)
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var res int
	for scanner.Scan() {
		txt := scanner.Text()
		dontSplits := strings.SplitN(txt, "don't()", -1)
		res += handle(dontSplits[0])
		dontSplits = dontSplits[1:]
		for _, s := range dontSplits {
			doSplits := strings.Split(s, "do()")
			if len(doSplits) > 1 {
				for i := range doSplits {
					if i == 0 {
						continue
					}
					res += handle(doSplits[i])
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(res)
}

func handle(txt string) int {
	var res int
	fmt.Printf("handle str: %s\n", txt)
	matches := re.FindAllStringSubmatch(txt, -1)
	for _, m := range matches {
		var x, y int
		fmt.Sscanf(m[1], "%d", &x)
		fmt.Sscanf(m[2], "%d", &y)
		res += x * y
	}
	return res
}
