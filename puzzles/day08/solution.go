package main

import (
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type Box struct {
	x, y, z int
}
type Pair struct {
	a, b     int
	distance float64
}

func Part1(connectionsRequired int, pairs []Pair, parent map[int]int, boxes []Box) {
	for i := range connectionsRequired {
		if Find(pairs[i].a, parent) != Find(pairs[i].b, parent) {
			Union(pairs[i].a, pairs[i].b, parent)
		}
	}
	count := make(map[int]int)
	for i := range boxes {
		count[Find(i, parent)] += 1
	}

	values := make([]int, 0, len(count))
	for _, v := range count {
		values = append(values, v)
	}
	slices.Sort(values)
	slices.Reverse(values)

	top3 := values[:min(3, len(values))]
	result := top3[0] * top3[1] * top3[2]
	spew.Dump(result)
}

func Find(idx int, parent map[int]int) int {
	if parent[idx] != idx {
		return Find(parent[idx], parent)
	}
	return idx
}

func Union(x, y int, parent map[int]int) {
	parent[Find(y, parent)] = Find(x, parent)
}

func Distance(a, b Box) float64 {
	dx := float64(b.x - a.x)
	dy := float64(b.y - a.y)
	dz := float64(b.z - a.z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func ReadInputs(path string) []Box {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	boxes := []Box{}
	for _, l := range lines {
		fields := strings.Split(l, ",")
		x, _ := strconv.Atoi(fields[0])
		y, _ := strconv.Atoi(fields[1])
		z, _ := strconv.Atoi(fields[2])
		boxes = append(boxes, Box{x, y, z})
	}

	return boxes
}

func main() {
	boxes := ReadInputs("inputs/input-8.txt")
	// connectionsRequired := 1000
	pairs := []Pair{}

	for i := 0; i < len(boxes)-1; i++ {
		for j := i + 1; j < len(boxes); j++ {
			pairs = append(pairs, Pair{i, j, Distance(boxes[j], boxes[i])})
		}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].distance < pairs[j].distance
	})
	parent := make(map[int]int)
	for i := range boxes {
		parent[i] = i
	}
	// Part1(connectionsRequired, pairs, parent, boxes)
	lastConnected := [2]int{}

	for i := range pairs {
		if Find(pairs[i].a, parent) != Find(pairs[i].b, parent) {
			Union(pairs[i].a, pairs[i].b, parent)
			lastConnected = [2]int{pairs[i].a, pairs[i].b}
		}
	}
	result := boxes[lastConnected[0]].x * boxes[lastConnected[1]].x

	spew.Dump(result)
}
