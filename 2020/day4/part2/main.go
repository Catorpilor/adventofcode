package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

var (
	ecls = map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}
	reg = regexp.MustCompile(`^#[0-9a-f]{6}$`)
)

func isValid(row string) bool {
	segs := strings.Fields(row)
	fields := make(map[string]string, len(segs))
	for _, f := range segs {
		fields[f[:3]] = f[4:]
	}
	if val, ok := fields["ecl"]; !ok {
		return false
	} else {
		if !ecls[val] {
			return false
		}
	}
	if val, ok := fields["pid"]; !ok {
		return false
	} else {
		if len(val) != 9 {
			return false
		}
		if _, err := strconv.Atoi(val); err != nil {
			return false
		}
	}

	if val, ok := fields["byr"]; !ok {
		return false
	} else {
		num, err := strconv.Atoi(val)
		if err != nil {
			return false
		}
		if num < 1920 || num > 2002 {
			return false
		}
	}
	if val, ok := fields["iyr"]; !ok {
		return false
	} else {
		num, err := strconv.Atoi(val)
		if err != nil {
			return false
		}
		if num < 2010 || num > 2020 {
			return false
		}
	}
	if val, ok := fields["eyr"]; !ok {
		return false
	} else {
		num, err := strconv.Atoi(val)
		if err != nil {
			return false
		}
		if num < 2020 || num > 2030 {
			return false
		}
	}

	if val, ok := fields["hgt"]; !ok {
		return false
	} else {
		n := len(val)
		num, err := strconv.Atoi(val[:n-2])
		if err != nil {
			return false
		}
		unit := val[n-2:]
		if unit == "cm" {
			if num < 150 || num > 193 {
				return false
			}
		} else if unit == "in" {
			if num < 59 || num > 76 {
				return false
			}
		} else {
			return false
		}
	}
	if val, ok := fields["hcl"]; !ok {
		return false
	} else {
		if !reg.MatchString(val) {
			return false
		}
	}

	// fmt.Println(row)
	// if !strings.Contains(row, "ecl") {
	// 	return false
	// }
	// if !strings.Contains(row, "pid") {
	// 	return false
	// }
	// if !strings.Contains(row, "eyr") {
	// 	return false
	// }
	// if !strings.Contains(row, "hcl") {
	// 	return false
	// }
	// if !strings.Contains(row, "byr") {
	// 	return false
	// }
	// if !strings.Contains(row, "iyr") {
	// 	return false
	// }
	// if !strings.Contains(row, "hgt") {
	// 	return false
	// }
	return true
}
