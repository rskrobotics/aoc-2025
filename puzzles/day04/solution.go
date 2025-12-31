package main

import (
	"fmt"
	"os"
	"strings"
)

type Grid struct {
	data [][]rune
	xlen int
	ylen int
}

func NewGrid(lines []string) *Grid {
	g := new(Grid)
	for _, line := range lines {
		g.data = append(g.data, []rune(line))
	}
	g.xlen = len(g.data[0]) - 1
	g.ylen = len(g.data) - 1
	return g
}

func (g *Grid) Get(x, y int) (rune, bool) {
	if y < 0 || y > g.ylen || x < 0 || x > g.xlen {
		return 0, false
	}
	return g.data[y][x], true
}

func (g *Grid) CheckGrid() int {
	accessible := 0
	directions := [][]int{
		{0, 1},
		{0, -1},
		{-1, 0},
		{1, 0},
		{-1, -1},
		{-1, 1},
		{1, -1},
		{1, 1},
	}
	for y, line := range g.data {
		for x, c := range line {
			if c != '@' {
				continue
			}
			monkeCount := 0
			for i := range len(directions) {
				n, ok := g.Get(x+directions[i][0], y+directions[i][1])
				if !ok {
					continue
				}
				if n == '@' {
					monkeCount++
				}

			}
			if monkeCount < 4 {
				accessible++
			}
		}
	}
	return accessible
}

func ReadInputs(path string) []string {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	return lines
}

func main() {
	lines := ReadInputs("inputs/input-4.txt")
	grid := NewGrid(lines)
	result := grid.CheckGrid()
	fmt.Printf("Grid: %v\n", grid)
	fmt.Printf("Result: %v\n", result)
}
