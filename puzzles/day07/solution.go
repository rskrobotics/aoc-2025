package main

import (
	"fmt"
	"os"
	"strings"
)

type Grid struct {
	data [][]rune
}

func ReadInputs(path string) Grid {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	data := make([][]rune, len(lines))
	for y, s := range lines {
		data[y] = []rune(s)
	}

	return Grid{data: data}
}

func (g *Grid) countSplitters() int {
	splitters := 0

	for y := 0; y < len(g.data); y++ {
		for x := 0; x < len(g.data[y]); x++ {
			r := g.data[y][x]
			if r == 'S' {
				g.data[y+1][x] = '|'
			}
			if y != len(g.data)-1 {
				below := g.data[y+1][x]
				if r == '|' {
					switch below {
					case '^':
						g.data[y+1][x-1] = '|'
						g.data[y+2][x+1] = '|'
						splitters += 1
					case '.':
						g.data[y+1][x] = '|'
					}
				}
			}
		}
	}
	return splitters
}

func main() {
	// g := ReadInputs("inputs/input-7-sample.txt")
	g := ReadInputs("inputs/input-7.txt")
	fmt.Printf("Grid: %v\n", g)
	splitters := g.countSplitters()
	fmt.Printf("Splitters : %v\n", splitters)
}
