package main

import (
	"aoc/lib"
	"fmt"
	"regexp"
	"os"
	"strconv"
)

type Policy struct {
	min int
	max int
	char rune
}

func checkPasswordAgainstPolicy(password string, policy Policy) (bool) {
	count := lib.CountCharOccurrences(password, policy.char)
	return (policy.min <= count && count <= policy.max)
}

func main() {
	lines, err := lib.GetLinesFromFile(os.Args[1])
	lib.CheckErr(err)

	re := regexp.MustCompile("([0-9]+)-([0-9]+) ([a-z]): ([a-z]*)")

	count := 0
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		
		min, err := strconv.Atoi(matches[1])
		lib.CheckErr(err)
		
		max, err := strconv.Atoi(matches[2])
		lib.CheckErr(err)
		
		char := rune(matches[3][0])

		policy := Policy{ min, max, char }
		password := matches[4]
		if (checkPasswordAgainstPolicy(password, policy)) {
			count++
		}
	}

	fmt.Println(count)
}
