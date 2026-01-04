package main

import (
	"iter"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type Point struct {
	x, y int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func rangeInclusive(a, b int) iter.Seq[int] {
	if a > b {
		a, b = b, a
	}
	return func(yield func(int) bool) {
		for i := a; i <= b; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

func compressedSpace(points []Point) (map[int]int, map[int]int) {
	xSet := make(map[int]bool)
	ySet := make(map[int]bool)
	for _, p := range points {
		xSet[p.x] = true
		ySet[p.y] = true
	}
	xKeys := []int{}
	for k := range xSet {
		xKeys = append(xKeys, k)
	}
	slices.Sort(xKeys)

	yKeys := []int{}
	for k := range ySet {
		yKeys = append(yKeys, k)
	}
	slices.Sort(yKeys)

	xToCompressed := make(map[int]int)
	for i, x := range xKeys {
		xToCompressed[x] = i
	}
	yToCompressed := make(map[int]int)
	for i, y := range yKeys {
		yToCompressed[y] = i
	}
	return xToCompressed, yToCompressed
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

func isValidRectangle(cx1, cy1, cx2, cy2 int, visited map[[2]int]bool) bool {
	for x := min(cx1, cx2); x <= max(cx1, cx2); x++ {
		for y := min(cy1, cy2); y <= max(cy1, cy2); y++ {
			px, py := x+1, y+1
			if visited[[2]int{py, px}] {
				return false
			}
		}
	}
	return true
}

func main() {
	points := ReadInputs("inputs/input-9.txt")
	xToCompressed, yToCompressed := compressedSpace(points)
	border := make(map[Point]bool)
	for i := range points {
		curr := points[i]
		next := points[(i+1)%len(points)]
		cx1, cy1 := xToCompressed[curr.x], yToCompressed[curr.y]
		cx2, cy2 := xToCompressed[next.x], yToCompressed[next.y]
		if cx1 == cx2 {
			// vertical
			for y := range rangeInclusive(cy1, cy2) {
				border[Point{cx1, y}] = true
			}
		} else {
			// horizontal
			for x := range rangeInclusive(cx1, cx2) {
				border[Point{x, cy1}] = true
			}
		}
	}
	// flood fill
	visited := make(map[[2]int]bool)
	stack := [][2]int{{0, 0}}
	paddedY, paddedX := len(yToCompressed)+2, len(xToCompressed)+2
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(stack) > 0 {
		pos := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		py, px := pos[0], pos[1]

		if py < 0 || py >= paddedY || px < 0 || px >= paddedX {
			continue
		}

		if visited[pos] {
			continue
		}
		rx, ry := px-1, py-1
		if ry >= 0 && ry < len(yToCompressed) && rx >= 0 && rx < len(xToCompressed) {
			if border[Point{rx, ry}] {
				continue
			}
		}

		visited[pos] = true

		for _, d := range directions {
			stack = append(stack, [2]int{py + d[0], px + d[1]})
		}

	}
	maxArea := 0
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			p1, p2 := points[i], points[j]
			cx1, cy1 := xToCompressed[p1.x], yToCompressed[p1.y]
			cx2, cy2 := xToCompressed[p2.x], yToCompressed[p2.y]
			if isValidRectangle(cx1, cy1, cx2, cy2, visited) {
				width := abs(p2.x-p1.x) + 1
				height := abs(p2.y-p1.y) + 1
				area := width * height
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	spew.Dump(maxArea)
}
