package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput(path string) []string {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSpace(string(file)), "\n")
}

func CircularBuffer(current int, max int) int {
	return ((current % max) + max) % max
}

func main() {
	lines := ReadInput("inputs/input-1.txt")
	result := 0
	curr := 50

	for _, v := range lines {
		fmt.Println(v)
		sign := v[0]
		cranks, _ := strconv.Atoi(v[1:])
		if sign == 'L' {
			cranks = cranks * -1
		}
		curr = CircularBuffer(curr+cranks, 100)
		if curr == 0 {
			result += 1
		}

	}

	fmt.Printf("Result is %v\n", result)
}
