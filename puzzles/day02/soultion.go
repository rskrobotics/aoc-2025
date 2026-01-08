package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IDRange struct {
	s, e int
}

func ReadInput(path string) []IDRange {
	file, err := os.ReadFile(path)
	var r []IDRange
	if err != nil {
		panic(err)
	}
	splits := strings.SplitSeq(strings.TrimSpace(string(file)), ",")

	for v := range splits {
		parts := strings.Split(v, "-")
		S, _ := strconv.Atoi(parts[0])
		E, _ := strconv.Atoi(parts[1])

		r = append(r, IDRange{S, E})
	}
	return r
}

// func isDoubleNumber(i int) bool {
// 	digits := 0
// 	for t := i; t > 0; t /= 10 {
// 		digits++
// 	}
// 	if digits%2 == 1 || digits == 0 {
// 		return false
// 	}
// 	divisor := int(math.Pow10(digits / 2))
// 	firstHalf := i / divisor
// 	secondHalf := i % divisor
//
// 	return firstHalf == secondHalf
// }

func containsSequences(i int) bool {
	if i < 10 {
		return false
	}
	number := strconv.Itoa(i)
	for seqLen := 1; seqLen <= len(number)/2; seqLen++ {
		if len(number)%seqLen != 0 {
			continue
		}

		sequences := []string{}
		for s := 0; s < len(number); s += seqLen {
			sequences = append(sequences, number[s:s+seqLen])
		}

		allSame := true
		for s := 0; s < len(sequences)-1; s++ {
			if sequences[s] != sequences[s+1] {
				allSame = false
				break
			}
		}
		if allSame {
			return true
		}
	}
	return false
}

func main() {
	ret := 0
	ranges := ReadInput("inputs/input-2.txt")
	fmt.Printf("Ranges %v\n", ranges)
	for _, v := range ranges {
		for i := v.s; i <= v.e; i++ {
			if containsSequences(i) {
				ret += i
			}
		}
	}
	fmt.Printf("Value is %v\n", ret)
}
