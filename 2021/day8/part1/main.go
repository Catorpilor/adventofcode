package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var res int
	for scanner.Scan() {
		segs := strings.Split(scanner.Text(), " | ")
		digs := strings.Split(segs[1], " ")
		for _, s := range digs {
			n := len(s)
			if n == 2 || n == 4 || n == 7 || n == 3 {
				fmt.Printf("%s count\n", s)
				res++
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

}
