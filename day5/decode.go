package main

import (
	"strconv"
	"strings"
)

func DecodeRow(sequence string) (uint64, error) {
	sequence = strings.ReplaceAll(sequence, "B", "1")
	sequence = strings.ReplaceAll(sequence, "F", "0")
	return strconv.ParseUint(sequence, 2, 7)
}

func DecodeSeat(sequence string) (uint64, error) {
	sequence = strings.ReplaceAll(sequence, "R", "1")
	sequence = strings.ReplaceAll(sequence, "L", "0")
	return strconv.ParseUint(sequence, 2, 3)
}