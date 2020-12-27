package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var pubs []int
	for scanner.Scan() {
		txt := scanner.Text()
		num, _ := strconv.Atoi(txt)
		pubs = append(pubs, num)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(pubs)
	sl1 := solve(pubs[0])
	sl2 := solve(pubs[1])
	fmt.Println(sl1, sl2)
	ek1 := encryptionKey(pubs[0], sl2)
	ek2 := encryptionKey(pubs[1], sl1)
	fmt.Println(ek1, ek2)
}

func solve(pub int) int {
	val := 1
	var ans int
	for val != pub {
		val *= 7
		val %= 20201227
		ans++
	}
	fmt.Println(val)
	return ans
}

func encryptionKey(pub, loop int) int {
	var i int
	val := 1
	for i < loop {
		val *= pub
		val %= 20201227
		i++
	}
	return val
}
