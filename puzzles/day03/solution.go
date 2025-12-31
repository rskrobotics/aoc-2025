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
	maxJolts := 0
	tens := int(line[0] - '0')
	for i := 1; i < len(line); i++ {
		c := int(line[i] - '0')
		if tens*10+c > maxJolts {
			maxJolts = tens*10 + c
		}
		if c > tens {
			tens = c
		}
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
