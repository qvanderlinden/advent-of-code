package main

import (
	"aoc/lib"
	"fmt"
	"regexp"
	"os"
	"strconv"
)

type Policy struct {
	firstPosition int
	secondPosition int
	char rune
}

func checkPasswordAgainstPolicy(password string, policy Policy) (bool) {
	subStr := string(password[policy.firstPosition - 1]) + string(password[policy.secondPosition - 1])
	count := lib.CountCharOccurrences(subStr, policy.char)
	return count == 1
}

func main() {
	lines, err := lib.GetLinesFromFile(os.Args[1])
	lib.CheckErr(err)

	re := regexp.MustCompile("([0-9]+)-([0-9]+) ([a-z]): ([a-z]*)")

	count := 0
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		
		firstPosition, err := strconv.Atoi(matches[1])
		lib.CheckErr(err)
		
		secondPosition, err := strconv.Atoi(matches[2])
		lib.CheckErr(err)
		
		char := rune(matches[3][0])

		policy := Policy{ firstPosition, secondPosition, char }
		password := matches[4]
		if (checkPasswordAgainstPolicy(password, policy)) {
			count++
		}
	}

	fmt.Println(count)
}
