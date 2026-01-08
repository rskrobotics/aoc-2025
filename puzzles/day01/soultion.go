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
		fullRev := cranks / 100
		cranks = cranks % 100
		result += fullRev
		if sign == 'L' {
			cranks = cranks * -1
		}
		if (cranks > 0 && curr+cranks >= 100) || (cranks < 0 && curr > 0 && curr+cranks <= 0) {
			result += 1
		}
		curr = CircularBuffer(curr+cranks, 100)

	}

	fmt.Printf("Result is %v\n", result)
}
