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
		minC, _ := strconv.Atoi(fields[0][:dashPos])
		maxC, _ := strconv.Atoi(fields[0][dashPos+1:])
		var count int
		for i := range fields[2] {
			if fields[2][i] == fields[1][0] {
				count++
			}
		}
		if count >= minC && count <= maxC {
			ans++
		}
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(ans)
}
