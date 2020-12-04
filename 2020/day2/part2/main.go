package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var ans int
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		dashPos := strings.Index(fields[0], "-")
		p1, _ := strconv.Atoi(fields[0][:dashPos])
		p2, _ := strconv.Atoi(fields[0][dashPos+1:])
		var count int
		if fields[2][p1-1] == fields[1][0] {
			count++
		}
		if fields[2][p2-1] == fields[1][0] {
			count++
		}
		if count == 1 {
			ans++
		}

		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(ans)
}
