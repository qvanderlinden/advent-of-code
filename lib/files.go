package lib

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func GetLinesFromFile(filepath string) ([]string, error) {
	content, err := ioutil.ReadFile(filepath)
	if (err != nil) {
		return nil, err
	}

	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")
	return lines, nil
}

func GetLineGroupsFromFile(filepath string) ([]string) {
	rawContent, err := ioutil.ReadFile(filepath)
	CheckErr(err)

	content := string(rawContent)
	re := regexp.MustCompile(`(?m:\n\s*\n+)`)
	groups := re.Split(content, -1)

	return groups
}

func GetValuesFromFile(filepath string) ([]int, error) {
	lines, err := GetLinesFromFile(filepath)
	if (err != nil) {
		return nil, err
	}

	values := make([]int, len(lines))
	for i, v := range lines {
		val, err := strconv.Atoi(v)
		if (err != nil) {
			return nil, err
		}
		values[i] = val
	}
	return values, nil
}
