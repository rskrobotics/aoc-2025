package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type Grid [][]bool

type Present struct {
	shape     [3][3]bool
	cellCount int
}

type ShapeOption struct {
	orientations []Present
	cellCount    int
}

type Region struct {
	width   int
	heigth  int
	amounts [6]int
}

func countCells(shape [3][3]bool) int {
	count := 0
	for y := range 3 {
		for x := range 3 {
			if shape[y][x] {
				count++
			}
		}
	}
	return count
}

// +90 rotation
func (p Present) Rotate() Present {
	var new [3][3]bool
	for y := range 3 {
		for x := range 3 {
			new[x][2-y] = p.shape[y][x]
		}
	}
	return Present{new, p.cellCount}
}

func (p Present) Flip() Present {
	var new [3][3]bool
	for y := range 3 {
		for x := range 3 {
			new[2-y][x] = p.shape[y][x]
		}
	}
	return Present{new, p.cellCount}
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

func CanPlaceAt(grid Grid, shape Present, px, py int) bool {
	for y := range 3 {
		for x := range 3 {
			if shape.shape[y][x] {
				gx, gy := px+x, py+y
				if gx < 0 || gx >= len(grid[0]) || gy < 0 || gy >= len(grid) {
					return false
				}
				if grid[gy][gx] {
					return false
				}
			}
		}
	}
	return true
}

func Place(grid Grid, shape Present, px, py int) {
	for y := range 3 {
		for x := range 3 {
			if shape.shape[y][x] {
				grid[py+y][px+x] = true
			}
		}
	}
}

func Unplace(grid Grid, shape Present, px, py int) {
	for y := range 3 {
		for x := range 3 {
			if shape.shape[y][x] {
				grid[py+y][px+x] = false
			}
		}
	}
}

func countEmpty(grid Grid) int {
	count := 0
	for _, row := range grid {
		for _, cell := range row {
			if !cell {
				count++
			}
		}
	}
	return count
}

func CanPlace(grid Grid, toPlace []ShapeOption, cellsNeeded int) bool {
	if len(toPlace) == 0 {
		return true
	}

	// Pruning: check if enough empty cells remain
	if countEmpty(grid) < cellsNeeded {
		return false
	}

	// Take first shape and try all valid positions
	shape := toPlace[0]
	remaining := toPlace[1:]

	for _, orient := range shape.orientations {
		for py := 0; py < len(grid); py++ {
			for px := 0; px < len(grid[0]); px++ {
				if CanPlaceAt(grid, orient, px, py) {
					Place(grid, orient, px, py)
					if CanPlace(grid, remaining, cellsNeeded-shape.cellCount) {
						return true
					}
					Unplace(grid, orient, px, py)
				}
			}
		}
	}
	return false
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
		presents = append(presents, Present{shape, countCells(shape)})
	}
	for _, r := range regionLines {
		fields := strings.Fields(r)
		x := strings.Index(fields[0], "x")
		width, _ := strconv.Atoi(fields[0][:x])
		heigth, _ := strconv.Atoi(fields[0][x+1 : len(fields[0])-1])
		amounts := [6]int{}
		for i, v := range fields[1:] {
			amounts[i], _ = strconv.Atoi(v)
		}

		regions = append(regions, Region{width, heigth, amounts})
	}
	return presents, regions
}

func SolveRegion(region Region, presents []Present) bool {
	var toPlace []ShapeOption
	totalCells := 0
	for shapeIdx, count := range region.amounts {
		if count == 0 {
			continue
		}
		orients := presents[shapeIdx].AllOrientations()
		cellCount := presents[shapeIdx].cellCount
		for range count {
			toPlace = append(toPlace, ShapeOption{orients, cellCount})
			totalCells += cellCount
		}
	}

	// Quick check: do we have enough space?
	if totalCells > region.width*region.heigth {
		return false
	}

	grid := make(Grid, region.heigth)
	for i := range grid {
		grid[i] = make([]bool, region.width)
	}

	return CanPlace(grid, toPlace, totalCells)
}

func main() {
	presents, regions := ReadInputs("inputs/input-12.txt")
	result := 0
	for _, r := range regions {
		if SolveRegion(r, presents) {
			result += 1
		}
	}

	spew.Dump(result)
}
