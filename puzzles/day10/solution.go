package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type Machine struct {
	lights  int
	buttons []int
	jolts   []int
}

func ReadInputs(path string) []Machine {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	machines := []Machine{}

	for _, line := range lines {

		sLights := 1
		eLights := strings.Index(line, "]") - 1
		sJolts := strings.Index(line, "{")
		eJolts := strings.Index(line, "}")
		eButtons := sJolts - 2
		sButtons := eLights + 2

		lights := 0
		for i := sLights; i <= eLights; i++ {
			if c := line[i]; c == '#' {
				lights |= (1 << (i - sLights))
			}
		}

		jolts := []int{}
		joltsLine := line[sJolts+1 : eJolts]
		joltVals := strings.SplitSeq(joltsLine, ",")
		for v := range joltVals {
			vInt, _ := strconv.Atoi(v)
			jolts = append(jolts, vInt)
		}

		buttons := []int{}
		buttonLine := line[sButtons : eButtons+1]

		buttonVals := strings.FieldsSeq(buttonLine)
		for v := range buttonVals {
			inner := v[1 : len(v)-1]
			mask := 0
			for v := range strings.SplitSeq(inner, ",") {
				vInt, _ := strconv.Atoi(v)
				mask |= (1 << vInt)
			}
			buttons = append(buttons, mask)
		}

		machines = append(machines, Machine{lights: lights, buttons: buttons, jolts: jolts})
	}

	return machines
}

func main() {
	machines := ReadInputs("inputs/input-10.txt")
	result := 0

	// spew.Dump(machines)
	for _, m := range machines {
		queue := [][2]int{{0, 0}}
		visited := make(map[int]bool)
		for len(queue) > 0 {
			state, presses := queue[0][0], queue[0][1]
			queue = queue[1:]
			if state == m.lights {
				result += presses
				break
			}
			for _, b := range m.buttons {
				newState := state ^ b
				if !visited[newState] {
					visited[newState] = true
					queue = append(queue, [2]int{newState, presses + 1})
				}

			}
		}
	}
	spew.Dump(result)
}
