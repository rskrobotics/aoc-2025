package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func ReadInputs(path string) ([]Range, []int) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	splits := strings.Split(strings.TrimSpace(string(bytes)), "\n\n")
	rangesLines := strings.Split(splits[0], "\n")
	idLines := strings.Split(splits[1], "\n")

	ranges := []Range{}
	for _, line := range rangesLines {
		parts := strings.Split(line, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		ranges = append(ranges, Range{start, end})

	}

	IDs := []int{}
	for _, line := range idLines {
		id, _ := strconv.Atoi(line)
		IDs = append(IDs, id)
	}

	return ranges, IDs
}

func mergeRanges(ranges []Range) []Range {
	if len(ranges) == 0 {
		return ranges
	}
	writeIdx := 0
	for i := 1; i < len(ranges); i++ {
		if ranges[i].start <= ranges[writeIdx].end {
			if ranges[i].end > ranges[writeIdx].end {
				ranges[writeIdx].end = ranges[i].end
			}
		} else {
			writeIdx++
			ranges[writeIdx] = ranges[i]
		}
	}
	return ranges[:writeIdx+1]
}

func rangesSummary(ranges []Range) int {
	var allIDs int
	for _, r := range ranges {
		allIDs += r.end - r.start + 1
	}

	return allIDs
}

func main() {
	ranges, IDs := ReadInputs("inputs/input-5.txt")
	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i].start != ranges[j].start {
			return ranges[i].start < ranges[j].start
		}
		return ranges[i].end < ranges[j].end
	})
	ranges = mergeRanges(ranges)
	allIDs := rangesSummary(ranges)
	freshIngredients := 0
	for _, r := range ranges {
		for _, id := range IDs {
			if id >= r.start && id <= r.end {
				freshIngredients++
				continue
			}
		}
	}

	fmt.Printf("allIDs: %v\n", allIDs)
	fmt.Printf("Fresh: %v\n", freshIngredients)
	fmt.Printf("Ranges: %v\n", ranges)
	fmt.Printf("Ids: %v\n", IDs)
}
