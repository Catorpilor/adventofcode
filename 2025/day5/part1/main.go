package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	start, end int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var ranges []Range

	// Parse ranges
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		ranges = append(ranges, Range{start, end})
	}

	merged := mergeRanges(ranges)

	var res int
	// Query numbers
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		if inRange(merged, num) {
			res++
		}
	}
	fmt.Println(res)
}

func mergeRanges(ranges []Range) []Range {
	if len(ranges) == 0 {
		return nil
	}

	// Sort by start
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	// Merge overlapping
	merged := make([]Range, 0, len(ranges))
	merged = append(merged, ranges[0])
	for i := 1; i < len(ranges); i++ {
		last := &merged[len(merged)-1]
		curr := ranges[i]

		if curr.start <= last.end { // Overlapping
			if curr.end > last.end {
				last.end = curr.end
			}
		} else { // Non-overlapping
			merged = append(merged, curr)
		}
	}

	return merged
}

func inRange(ranges []Range, num int) bool {
	// Binary search: find rightmost range where start <= num
	left, right := 0, len(ranges)-1
	result := -1

	for left <= right {
		mid := (left + right) / 2
		if ranges[mid].start <= num {
			// find the right most range where start <= num
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if result == -1 {
		return false
	}
	return num <= ranges[result].end
}
