package main

import (
	"aoc/lib"
	"fmt"
	"log"
	"strconv"
	"strings"
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
	if (len(os.Args) != 3) {
		log.Fatal("[Usage]: <bin> <input-file> <slopes>\n<slopes> contains comma separated slope entities. A slope entity has the <dx>:<dy> format")
	}

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
	
	// get slopes
	slopesArg := os.Args[2]
	slopesStrs := strings.Split(slopesArg, ",")
	slopes := make([]Slope, len(slopesStrs))
	for i, slopeStr := range(slopesStrs) {
		d := strings.Split(slopeStr, ":")
		dx, err := strconv.Atoi(d[0])
		lib.CheckErr(err)
		dy, err := strconv.Atoi(d[1])
		lib.CheckErr(err)
		slopes[i] = Slope{ dx, dy }
	}
	
	result := 1
	for _, slope := range slopes {
		result *= CountTrees(forest, slope)
	}
	fmt.Println(result)
}
