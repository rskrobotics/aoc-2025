package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Homework struct {
	values   []int
	operator rune
}

func splitAtIndexes(s string, indexes []int) []string {
	var parts []string
	prev := 0
	for _, i := range indexes {
		parts = append(parts, s[prev:i])
		prev = i
	}
	parts = append(parts, s[prev:])
	return parts
}

func TransformCephalophodMath(path string) []Homework {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	// find longestLine from those suckers
	lenLongestLine := 0
	for _, l := range lines {
		if len(l) > lenLongestLine {
			lenLongestLine = len(l)
		}
	}
	// padding to equal lentgh
	paddedLines := make([]string, len(lines))
	for i, s := range lines {
		paddedLines[i] = fmt.Sprintf("%-*s", lenLongestLine, s)
	}
	// find separators
	separators := []int{}
	for i := 0; i < lenLongestLine; i++ {
		allSpaces := true
		for _, line := range paddedLines {
			if line[i] != ' ' {
				allSpaces = false
				break
			}
		}
		if allSpaces {
			separators = append(separators, i)
		}
	}
	fmt.Printf("Separators %v\n", separators)
	// split lines
	parts := [][]string{}
	for _, l := range lines {
		p := splitAtIndexes(l, separators)
		parts = append(parts, p)
	}
	// operator finder
	vParts := parts[:len(parts)-1]
	oParts := parts[len(parts)-1]
	hw := make([]Homework, len(oParts))
	// oprator
	for i := range len(oParts) {
		hw[i].operator = rune(strings.TrimSpace(oParts[i])[0])
	}
	fmt.Printf("vParts %v\n", vParts)
	// values
	for i := range len(vParts[0]) {
		for j := range len(vParts[0][i]) {
			actualVals := ""
			for l := range len(vParts) {
				c := string(vParts[l][i][j])
				if c != " " {
					actualVals = actualVals + c
				}
			}
			if len(actualVals) >= 1 {
				c, _ := strconv.Atoi(actualVals)

				hw[i].values = append(hw[i].values, c)
			}
		}
	}
	fmt.Printf("hw %v\n", hw)
	return hw
}

// func ReadInputs(path string) []Homework {
// 	bytes, err := os.ReadFile(path)
// 	if err != nil {
// 		panic(err)
// 	}
// 	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")
//
// 	hw := make([]Homework, len(strings.Fields(lines[0])))
// 	fmt.Printf("Line 0 cols: %d\n", len(strings.Fields(lines[0])))
//
// 	for i := 0; i < len(lines)-1; i++ {
// 		vals := strings.Fields(lines[i])
// 		if len(vals) != len(hw) {
// 			fmt.Printf("Line %d has %d cols (expected %d)\n", i, len(vals), len(hw))
// 		}
// 		for j := range vals {
// 			c, _ := strconv.Atoi(vals[j])
//
// 			hw[j].values = append(hw[j].values, c)
// 		}
// 	}
// 	vals := strings.Fields(lines[len(lines)-1])
// 	for j := range vals {
// 		hw[j].operator = rune(vals[j][0])
// 	}
//
// 	return hw
// }

func main() {
	hw := TransformCephalophodMath("inputs/input-6.txt")
	fmt.Printf("Homework: %v\n", hw)
	result := 0
	for _, h := range hw {
		var temp int
		if h.operator == '+' {
			temp = 0
			for _, v := range h.values {
				temp += v
			}
		}
		if h.operator == '*' {
			temp = 1
			for _, v := range h.values {
				temp *= v
			}
		}

		result += temp
	}
	fmt.Printf("Result: %v\n", result)
}
