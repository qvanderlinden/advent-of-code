package lib

import (
	"io/ioutil"
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
