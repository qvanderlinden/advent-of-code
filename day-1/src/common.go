package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func GetValuesFromFile(filepath string) ([]int, error) {
	content, err := ioutil.ReadFile(filepath)
	if (err != nil) {
		return nil, err
	}

	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")
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
