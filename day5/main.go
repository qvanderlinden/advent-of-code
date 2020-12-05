package main

import (
	"aoc/lib"
	"fmt"
	"math"
	"sort"
)

func main() {
	lines, err := lib.GetLinesFromFile("data/input-day-5.txt")
	lib.CheckErr(err)
	
	boardingPasses := make([]BoardingPass, 0)
	for _, line := range lines {
		boardingPass, err := DecodeBoardingPass(line)
		if (err != nil) {
			continue
		}
		
		boardingPasses = append(boardingPasses, boardingPass)
	}
	
	// part 1
	maxSeatId := math.Inf(-1)
	for _, boardingPass := range boardingPasses {
		maxSeatId = math.Max(maxSeatId, float64(boardingPass.seatId))
	}
	fmt.Println("Part 1 result:", maxSeatId)
	
	// part 2
	minRow := math.Inf(1)
	maxRow := math.Inf(-1)
	for _, boardingPass := range boardingPasses {
		row := float64(boardingPass.row)
		if (row > maxRow) {
			maxRow = row
		}
		if (row < minRow) {
			minRow = row
		}
	}

	nonExtremeRowSeatIds := make([]float64, 0)
	for _, boardingPass := range boardingPasses {
		row := float64(boardingPass.row)
		if (row > minRow && row < maxRow) {
			nonExtremeRowSeatIds = append(nonExtremeRowSeatIds, float64(boardingPass.seatId))
		}
	}
	sort.Float64s(nonExtremeRowSeatIds)
	for i, seatId := range nonExtremeRowSeatIds {
		if (i == 0) {
			continue
		}

		if (seatId != nonExtremeRowSeatIds[i - 1] + 1) {
			fmt.Println("Part 2 result:", nonExtremeRowSeatIds[i - 1] + 1)
		}
	}
}