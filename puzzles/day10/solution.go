package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type Machine struct {
	lights  []bool
	buttons [][]bool
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

		lights := []bool{}
		for i := sLights; i <= eLights; i++ {
			if c := line[i]; c == '#' {
				lights = append(lights, true)
			} else {
				lights = append(lights, false)
			}
		}

		jolts := []int{}
		joltsLine := line[sJolts+1 : eJolts]
		joltVals := strings.SplitSeq(joltsLine, ",")
		for v := range joltVals {
			vInt, _ := strconv.Atoi(v)
			jolts = append(jolts, vInt)
		}

		buttons := [][]bool{}
		buttonLine := line[sButtons : eButtons+1]

		buttonVals := strings.FieldsSeq(buttonLine)
		for v := range buttonVals {
			combination := make([]bool, len(lights))
			inner := v[1 : len(v)-1]
			for v := range strings.SplitSeq(inner, ",") {
				vInt, _ := strconv.Atoi(v)
				combination[vInt] = true
			}
			buttons = append(buttons, combination)
		}

		machines = append(machines, Machine{lights: lights, buttons: buttons, jolts: jolts})
	}

	return machines
}

func main() {
	machines := ReadInputs("inputs/input-10-sample.txt")

	spew.Dump(machines)
}
