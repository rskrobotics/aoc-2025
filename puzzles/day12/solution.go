package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type Present struct {
	shape [3][3]bool
}

type Region struct {
	width   int
	heigth  int
	amounts [6]int
}

// +90 rotation
func (p Present) Rotate() Present {
	var new [3][3]bool
	for y := range 3 {
		for x := range 3 {
			new[x][2-y] = p.shape[y][x]
		}
	}
	return Present{new}
}

func (p Present) Flip() Present {
	var new [3][3]bool
	for y := range 3 {
		for x := range 3 {
			new[2-y][x] = p.shape[y][x]
		}
	}
	return Present{new}
}

func (p Present) AllOrientations() []Present {
	seen := map[[3][3]bool]bool{}
	var results []Present

	current := p
	for range 2 {
		for range 4 {
			if !seen[current.shape] {
				seen[current.shape] = true
				results = append(results, current)
			}
			current = current.Rotate()
		}
		current = current.Flip()

	}

	return results
}

func ReadInputs(path string) ([]Present, []Region) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	presents := []Present{}
	regions := []Region{}
	sections := strings.Split(strings.TrimSpace(string(bytes)), "\n\n")
	presentSections := sections[:(len(sections) - 1)]
	regionSection := sections[(len(sections) - 1)]
	regionLines := strings.Split(regionSection, "\n")
	for _, p := range presentSections {
		presentLine := strings.Split(p, "\n")[1:]
		shape := [3][3]bool{}
		for y, v := range presentLine {
			for x, r := range v {
				if r == '#' {
					shape[y][x] = true
				}
			}
		}
		presents = append(presents, Present{shape})

	}
	for _, r := range regionLines {
		fields := strings.Fields(r)
		x := strings.Index(fields[0], "x")
		width, _ := strconv.Atoi(fields[0][:x])
		heigth, _ := strconv.Atoi(fields[0][x+1 : len(fields[0])-1])
		amounts := [6]int{}
		for i, v := range fields[1:] {
			spew.Dump(v)
			amounts[i], _ = strconv.Atoi(v)
		}

		regions = append(regions, Region{width, heigth, amounts})
	}
	return presents, regions
}

func main() {
	presents, regions := ReadInputs("inputs/input-12-sample.txt")

	spew.Dump(presents)
	spew.Dump(regions)
}
