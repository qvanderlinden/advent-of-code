package main

import (
	"aoc/lib"
	"fmt"
	"log"
	"os"
)

func main() {
	values, err := lib.GetValuesFromFile(os.Args[1])
	if (err != nil) {
		log.Fatal(err)
	}

	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			if (values[i] + values[j] == 2020) {
				fmt.Printf("%d\n", values[i] * values[j])
			}
		}
	}
}