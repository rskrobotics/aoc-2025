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

func dfs(racks map[string][]string, state []string, ret int) int {
	curr := state[len(state)-1]
	if curr == "out" {
		return ret + 1
	}

	choices := racks[curr]
	for _, choice := range choices {
		state = append(state, choice)
		ret = dfs(racks, state, ret)
		state = state[:len(state)-1]
	}
	return ret
}

func main() {
	racks := ReadInputs("inputs/input-11.txt")
	result := dfs(racks, []string{"you"}, 0)

	spew.Dump(racks)
	spew.Dump(result)
}
