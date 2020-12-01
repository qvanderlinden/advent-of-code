package main

import (
	"fmt"
	"log"
)

func main() {
	values, err := GetValuesFromFile("./data/input.txt")
	if (err != nil) {
		log.Fatal(err)
	}

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