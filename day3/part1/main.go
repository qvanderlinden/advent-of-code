package main

import (
	"aoc/lib"
	"fmt"
	"os"
)

type Forest struct {
	width int
	height int
	trees [][]bool
}

type Slope struct {
	dx int
	dy int
}

func CountTrees(forest Forest, slope Slope) (int) {
	count := 0
	i := 0
	j := 0
	for i < forest.height {
		if (forest.trees[i][j]) {
			count++
		}
		i += slope.dy
		j = (j + slope.dx) % forest.width
	}
	return count
}

func main() {
	lines, err := lib.GetLinesFromFile(os.Args[1])
	lib.CheckErr(err)
	
	// create the Forest Struct
	height := len(lines)
	width := len(lines[0])
	trees := make([][]bool, height)
	for i := range trees {
		trees[i] = make([]bool, width)
	}
	for i, line := range lines {
		for j, char := range line {
			trees[i][j] = (char == '#')
		}
	}

	forest := Forest{ width, height, trees }
	slope := Slope{ 3, 1 }

	result := CountTrees(forest, slope)
	fmt.Println(result)
}
