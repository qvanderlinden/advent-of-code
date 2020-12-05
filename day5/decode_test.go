package main

import (
	"testing"
)

type UnitTest struct {
	fName string
	f func(string) (uint64, error)
	sequence string
	expectedResult uint64
}

func TestDecodings(t *testing.T) {
	unitTests := []UnitTest{
		{ "DecodeRow", DecodeRow, "FBFBBFF", 44 },
		{ "DecodeSeat", DecodeSeat, "RLR", 5 },
	}

	for _, unitTest := range unitTests {
		result, err := unitTest.f(unitTest.sequence)
		if (err != nil) {
			t.Errorf("%s returned the following error: %v", unitTest.fName, err)
			continue
		}

		if (result != unitTest.expectedResult) {
			t.Errorf("%s(%s) = %d; expected %d", unitTest.fName, unitTest.sequence, result, unitTest.expectedResult)
		}
	}
}
