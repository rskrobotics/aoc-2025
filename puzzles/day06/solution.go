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

func ReadInputs(path string) []Homework {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	hw := make([]Homework, len(strings.Fields(lines[0])))
	fmt.Printf("Line 0 cols: %d\n", len(strings.Fields(lines[0])))

	for i := 0; i < len(lines)-1; i++ {
		vals := strings.Fields(lines[i])
		if len(vals) != len(hw) {
			fmt.Printf("Line %d has %d cols (expected %d)\n", i, len(vals), len(hw))
		}
		for j := range vals {
			c, _ := strconv.Atoi(vals[j])

			hw[j].values = append(hw[j].values, c)
		}
	}
	vals := strings.Fields(lines[len(lines)-1])
	for j := range vals {
		hw[j].operator = rune(vals[j][0])
	}

	return hw
}

func main() {
	hw := ReadInputs("inputs/input-6.txt")
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
	// fmt.Printf("Ranges: %v\n", ranges)
	// fmt.Printf("Ids: %v\n", IDs)
}
