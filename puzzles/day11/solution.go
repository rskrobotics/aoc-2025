package main

import (
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func ReadInputs(path string) map[string][]string {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	racks := make(map[string][]string)
	lines := strings.SplitSeq(strings.TrimSpace(string(bytes)), "\n")

	for l := range lines {
		fields := strings.Fields(l)
		rack := fields[0][:len(fields[0])-1]
		values := fields[1:]
		racks[rack] = values
	}
	return racks
}

type Key struct {
	curr  string
	isDAC bool
	isFFT bool
}

func dfs2(racks map[string][]string, curr string, isDAC, isFFT bool, memo map[Key]int) int {
	if curr == "out" {
		if isDAC && isFFT {
			return 1
		}
		return 0
	}
	if value, ok := memo[Key{curr, isDAC, isFFT}]; ok {
		return value
	}

	count := 0
	for _, c := range racks[curr] {
		newDAC := isDAC || c == "dac"
		newFFT := isFFT || c == "fft"
		count += dfs2(racks, c, newDAC, newFFT, memo)
	}
	memo[Key{curr, isDAC, isFFT}] = count
	return count
}

// func dfs(racks map[string][]string, state []string, ret int) int {
// 	curr := state[len(state)-1]
// 	if curr == "out" {
// 		isDAC := false
// 		isFFT := false
// 		for _, v := range state {
// 			if v == "dac" {
// 				isDAC = true
// 			}
// 			if v == "fft" {
// 				isFFT = true
// 			}
// 		}
//
// 		if isDAC && isFFT {
// 			return ret + 1
// 		}
// 		return ret
// 	}
//
// 	choices := racks[curr]
// 	for _, choice := range choices {
// 		state = append(state, choice)
// 		ret = dfs(racks, state, ret)
// 		state = state[:len(state)-1]
// 	}
// 	return ret
// }

func main() {
	racks := ReadInputs("inputs/input-11.txt")
	result := dfs2(racks, "svr", false, false, make(map[Key]int))

	spew.Dump(racks)
	spew.Dump(result)
}
