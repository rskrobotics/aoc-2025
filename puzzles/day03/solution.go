package main

import (
	"fmt"
	"os"
	"strings"
)

func ReadInput(path string) []string {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSpace(string(file)), "\n")
}

func MaxJolts(line string) int {
	discardsLeft := len(line) - 12
	stack := []int{}
	for i := 0; i < len(line); i++ {
		c := int(line[i] - '0')
		for len(stack) > 0 && discardsLeft > 0 && stack[len(stack)-1] < c {
			stack = stack[:len(stack)-1]
			discardsLeft--
		}
		stack = append(stack, c)
	}
	stack = stack[:12]
	maxJolts := 0
	for _, v := range stack {
		maxJolts = maxJolts*10 + v
	}

	return maxJolts
}

func main() {
	ret := 0
	lines := ReadInput("inputs/input-3.txt")
	for _, line := range lines {
		ret += MaxJolts(line)
	}
	fmt.Printf("Value is %v\n", ret)
}
