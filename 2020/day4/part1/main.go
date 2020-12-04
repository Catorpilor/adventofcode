package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var sb strings.Builder
	// var rows [][]byte
	// var tmp string
	// var rows []string
	var ans int
	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			sb.WriteString(text)
			sb.WriteByte(' ')
			// tmp += text
			// tmp += " "
			// fmt.Printf("cur sb: %s\n", sb.String())
		} else {
			// fmt.Println("*******")
			// row := sb.String()
			// fmt.Printf("cur sb: %s\n", tmp)
			// local := make([]byte, len(tmp))
			// copy(local, []byte(tmp))
			// // fmt.Println("*******")
			// rows = append(rows, sb.String())
			// fmt.
			if isValid(sb.String()) {
				ans++
			}
			sb.Reset()
			// tmp = ""
		}
		// fmt.Printf("text: %s extra\n", scanner.Text()) // Println will add back the final '\n'
		// if scanner.Text() == "" {
		// 	fmt.Println("aha")
		// }
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	if isValid(sb.String()) {
		ans++
	}

	// rows = append(rows, []byte(tmp))
	// var ans int
	// for _, row := range rows {
	// 	if isValid(string(row)) {
	// 		ans++
	// 	}
	// }
	fmt.Println(ans)
}

func isValid(row string) bool {
	// fmt.Println(row)
	if !strings.Contains(row, "ecl") {
		return false
	}
	if !strings.Contains(row, "pid") {
		return false
	}
	if !strings.Contains(row, "eyr") {
		return false
	}
	if !strings.Contains(row, "hcl") {
		return false
	}
	if !strings.Contains(row, "byr") {
		return false
	}
	if !strings.Contains(row, "iyr") {
		return false
	}
	if !strings.Contains(row, "hgt") {
		return false
	}
	return true
}
