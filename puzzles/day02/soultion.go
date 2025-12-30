package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IDRange struct {
	S, E int
}

func ReadInput(path string) []IDRange {
	file, err := os.ReadFile(path)
	r := []IDRange{}
	if err != nil {
		panic(err)
	}
	splits := strings.Split(strings.TrimSpace(string(file)), ",")

	for _, v := range splits {
		parts := strings.Split(v, "-")
		S, _ := strconv.Atoi(parts[0])
		E, _ := strconv.Atoi(parts[1])

		r = append(r, IDRange{S, E})
	}
	return r
}

func main() {
	ret := 0
	ranges := ReadInput("inputs/input-2.txt")
	fmt.Printf("Ranges %v\n", ranges)
	for _, v := range ranges {
		for i := v.S; i <= v.E; i++ {
			evaluated := strconv.Itoa(i)
			splitter := len(evaluated) / 2
			if len(evaluated)%2 == 1 {
				continue
			}
			if evaluated[:splitter] == evaluated[splitter:] {
				ret += i
			}
		}
	}
	fmt.Printf("Value is %v\n", ret)
}
