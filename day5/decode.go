package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type BoardingPass struct{
	row uint64
	seat uint64
	seatId uint64
}

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

func DecodeBoardingPass(sequence string) (BoardingPass, error) {
	re := regexp.MustCompile("^([F|B]{7})([L|R]{3})$")
	matches := re.FindStringSubmatch(sequence)
	if (len(matches) == 0) {
		return BoardingPass{}, errors.New("Incorrect input")
	}

	rowSequence := matches[1]
	seatSequence := matches[2]
	
	row, err := DecodeRow(rowSequence)
	if (err != nil) {
		return BoardingPass{}, err
	}
	
	seat, err := DecodeSeat(seatSequence)
	if (err != nil) {
		return BoardingPass{}, err
	}

	seatId := row * 8 + seat
	return BoardingPass{ row, seat, seatId }, nil
}