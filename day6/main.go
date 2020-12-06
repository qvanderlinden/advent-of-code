package main

import (
	"aoc/lib"
	"fmt"
	"strings"
)

func CountQuestionsPositiveAnswers(group string) (map[rune]int) {
	positiveAnswersCounts := make(map[rune]int)
	for _, char := range group {
		if (char != '\n') {
			positiveAnswersCounts[char] = positiveAnswersCounts[char] + 1
		}
	}
	return positiveAnswersCounts
}

func CountQuestionsWithPositiveAnswers(group string) (int) {
	positiveAnswersCounts := CountQuestionsPositiveAnswers(group)
	return len(positiveAnswersCounts)
}

func CountQuestionsWithOnlyPositiveAnswers(group string) (int) {
	positiveAnswersCounts := CountQuestionsPositiveAnswers(group)
	nPeople := len(strings.Split(group, "\n"))
	res := 0
	for _, value := range positiveAnswersCounts {
		if (value == nPeople) {
			res++
		}
	}
	return res
}

func main() {
	groups := lib.GetLineGroupsFromFile("data/input-day-6.txt")

	resPart1 := 0
	resPart2 := 0
	for _, group := range groups {
		resPart1 += CountQuestionsWithPositiveAnswers(group)
		resPart2 += CountQuestionsWithOnlyPositiveAnswers(group)
	}
	fmt.Println("Part 1 result:", resPart1)
	fmt.Println("Part 2 result:", resPart2)
}
