package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var res int
	for scanner.Scan() {
		// segs := strings.Split(scanner.Text(), " | ")

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
