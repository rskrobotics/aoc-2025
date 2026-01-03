package main

import (
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type Point struct {
	x, y int
}

func Area(a, b Point) int {
	width := int(math.Abs(float64(b.x-a.x))) + 1
	height := int(math.Abs(float64(b.y-a.y))) + 1
	return width * height
}

func ReadInputs(path string) []Point {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	boxes := []Point{}
	for _, l := range lines {
		fields := strings.Split(l, ",")
		x, _ := strconv.Atoi(fields[0])
		y, _ := strconv.Atoi(fields[1])
		boxes = append(boxes, Point{x, y})
	}

	return boxes
}

func main() {
	points := ReadInputs("inputs/input-9.txt")
	maxArea := 0
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			if curr := Area(points[i], points[j]); curr > maxArea {
				maxArea = curr
			}
		}
	}
	spew.Dump(maxArea)
}
