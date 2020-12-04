package main

import (
	"aoc/lib"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

func Between(value int, lower int, upper int) (bool) {
	return value >= lower && value <= upper
}

func ValidateStringNumber(number string, nDigits int, min int, max int) (bool) {
	if (len(number) != nDigits) {
		return false
	}
	
	value, err := strconv.Atoi(number)
	if (err != nil) {
		return false
	}

	return Between(value, min, max)
}

func ValidateHeight(height string) (bool) {
	re := regexp.MustCompile("^([0-9]+)(in|cm)$")
	matches := re.FindStringSubmatch(height)
	if (len(matches) == 0) {
		return false
	}

	value, err := strconv.Atoi(matches[1])
	if (err != nil) {
		return false
	}
	unit := matches[2]

	if (unit == "cm") {
		return Between(value, 150, 193)
	}
	if (unit == "in") {
		return Between(value, 59, 76)
	}

	return false
}

func ValidateHairColor(color string) (bool) {
	re := regexp.MustCompile("^#[0-9a-f]{6}$")
	return re.MatchString(color)
}

func ValidateEyeColor(color string) (bool) {
	switch (color) {
	case "amb":
		return true
	case "blu":
		return true
	case "brn":
		return true
	case "gry":
		return true
	case "grn":
		return true
	case "hzl":
		return true
	case "oth":
		return true
	}
	return false
}

func ValidatePasssportId(str string) (bool) {
	re := regexp.MustCompile("^[0-9]{9}$")
	return re.MatchString(str)
}

func IsPassportValid(passportMap map[string]string) (bool) {
	if (!ValidateStringNumber(passportMap["byr"], 4, 1920, 2002)) {
		return false
	}
	if (!ValidateStringNumber(passportMap["iyr"], 4, 2010, 2020)) {
		return false
	}
	if (!ValidateStringNumber(passportMap["eyr"], 4, 2020, 2030)) {
		return false
	}
	if (!ValidateHeight(passportMap["hgt"])) {
		return false
	}
	if (!ValidateHairColor(passportMap["hcl"])) {
		return false
	}
	if (!ValidateEyeColor(passportMap["ecl"])) {
		return false
	}
	if (!ValidatePasssportId(passportMap["pid"])) {
		return false
	}

	// nothing to check for cid
	
	return true
}

func main() {
	rawContent, err := ioutil.ReadFile(os.Args[1])
	lib.CheckErr(err)
	
	content := string(rawContent)
	blankLineRe := regexp.MustCompile("(?m:^[ ]*\n)")
	allPassportGroups := blankLineRe.Split(content, -1)
	
	count := 0
	pairsRe := regexp.MustCompile("(?m:([a-z]+):([\\S]+))")

	for _, passport := range allPassportGroups {
		passportMap := make(map[string]string)

		matches := pairsRe.FindAllStringSubmatch(passport, -1)
		for _, match := range matches {
			passportMap[match[1]] = match[2]
		}
		
		valid := IsPassportValid(passportMap)	
		if (valid) {
			count++
		}
	}

	fmt.Println(count)
}
