package day04

import (
	"fmt"
	"regexp"
	"strconv"
)

var heightRE = regexp.MustCompile("^(\\d+)(cm|in)$")
var hclRE = regexp.MustCompile("^#[0-9a-f]{6}$")
var eclRE = regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$")
var pidRE = regexp.MustCompile("^[0-9]{9}$")

func validateNumber(s string, min int, max int) bool {
	if num, err := strconv.Atoi(s); err != nil || num < min || num > max {
		return false
	}
	return true
}

func isPassportValid2(p passport) bool {
	if p.byr == "" {
		return false
	}
	if !validateNumber(p.byr, 1920, 2002) {
		return false
	}
	if p.iyr == "" {
		return false
	}
	if !validateNumber(p.iyr, 2010, 2020) {
		return false
	}
	if p.eyr == "" {
		return false
	}
	if !validateNumber(p.eyr, 2020, 2030) {
		return false
	}
	if p.hgt == "" {
		return false
	}
	if !heightRE.Match([]byte(p.hgt)) {
		return false
	}
	submatch := heightRE.FindAllStringSubmatch(p.hgt, -1)
	if submatch[0][2] == "cm" {
		if !validateNumber(submatch[0][1], 150, 193) {
			return false
		}
	} else {
		if !validateNumber(submatch[0][1], 59, 76) {
			return false
		}
	}
	if p.hcl == "" {
		return false
	}
	if !hclRE.Match([]byte(p.hcl)) {
		return false
	}
	if p.ecl == "" {
		return false
	}
	if !eclRE.Match([]byte(p.ecl)) {
		return false
	}
	if p.pid == "" {
		return false
	}
	if !pidRE.Match([]byte(p.pid)) {
		return false
	}
	return true
}

func CountValidPassports2() {
	passports := loadPassportDB()
	count := 0
	for _, p := range passports {
		if isPassportValid2(p) {
			count++
		}
	}

	fmt.Println("valid passports: ", count)
}
