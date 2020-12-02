package main

import (
	"aoc/lib"
	"fmt"
	"os"
)

func main() {
	values, err := lib.GetValuesFromFile(os.Args[1])
	lib.CheckErr(err)

	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			for k := j + 1; k < len(values); k++ {
				if (values[i] + values[j] + values[k] == 2020) {
					fmt.Printf("%d\n", values[i] * values[j] * values[k])
				}
			}
		}
	}
}