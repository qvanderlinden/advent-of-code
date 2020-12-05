package main

import (
	"testing"
)

func TestDecodeBoardingPass(t *testing.T) {
	sequence := "FBFBBFFRLR"
	expectedResult := BoardingPass{ 44, 5, 357 }
	
	result, err := DecodeBoardingPass(sequence)
	if (err != nil) {
		t.Logf("DecodeBoardingPass returned the following error: %v", err)
		t.FailNow()
	}

	isResultCorrect := result.row == expectedResult.row && result.seat == expectedResult.seat && result.seatId == expectedResult.seatId
	if (!isResultCorrect) {
		t.Errorf("DecodeBoardingPass(%s) = %v; expected %v", sequence, result, expectedResult)
	}
}
