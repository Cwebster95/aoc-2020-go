package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func (p *passport) hasRequiredProperties() bool {
	return p.byr != "" &&
		p.iyr != "" &&
		p.eyr != "" &&
		p.hgt != "" &&
		p.hcl != "" &&
		p.ecl != "" &&
		p.pid != ""
}

func validFourDigitYear(year string, low int, high int) bool {
	matchesPattern, err := regexp.Match(`^\d{4}$`, []byte(year))
	if err != nil || !matchesPattern {
		return false
	}

	intValue, err := strconv.Atoi(year)
	if err != nil || intValue < low || intValue > high {
		return false
	}

	return true
}

func validAgainstPattern(value string, pattern string) bool {
	match, err := regexp.Match(pattern, []byte(value))
	if err != nil {
		return false
	}

	return match
}

func (p *passport) validBYR() bool {
	return validFourDigitYear(p.byr, 1920, 2002)
}

func (p *passport) validIYR() bool {
	return validFourDigitYear(p.iyr, 2010, 2020)
}

func (p *passport) validEYR() bool {
	return validFourDigitYear(p.eyr, 2020, 2030)
}

func (p *passport) validHGT() bool {
	r := regexp.MustCompile(`^(\d+)(cm|in)$`)
	if !r.Match([]byte(p.hgt)) {
		return false
	}

	matchData := r.FindStringSubmatch(p.hgt)
	height, unit := matchData[1], matchData[2]
	intHeight, err := strconv.Atoi(height)
	if err != nil {
		return false
	}

	switch unit {
	case "cm":
		return intHeight >= 150 && intHeight <= 193
	case "in":
		return intHeight >= 59 && intHeight <= 76
	default:
		return false
	}
}

func (p *passport) validHCL() bool {
	return validAgainstPattern(p.hcl, `^#[0-9a-f]{6}$`)
}

func (p *passport) validECL() bool {
	return validAgainstPattern(p.ecl, `^amb$|^blu$|^brn$|^gry$|^grn$|^hzl$|^oth$`)
}

func (p *passport) validPID() bool {
	return validAgainstPattern(p.pid, `^\d{9}$`)
}

func (p *passport) valid() bool {
	return p.validBYR() &&
		p.validIYR() &&
		p.validEYR() &&
		p.validHGT() &&
		p.validHCL() &&
		p.validECL() &&
		p.validPID()
}

func parseInput(input []byte) []string {
	return strings.Split(string(input), "\n")
}

func inputToPassportStrings(data []string) []string {
	passportStrings := make([]string, 0)
	currentPassportString := ""

	for _, s := range data {
		if s == "" {
			passportStrings = append(passportStrings, currentPassportString)
			currentPassportString = ""
		} else {
			if currentPassportString == "" {
				currentPassportString = s
			} else {
				currentPassportString = fmt.Sprintf("%s %s", currentPassportString, s)
			}
		}
	}

	if currentPassportString != "" {
		passportStrings = append(passportStrings, currentPassportString)
	}

	return passportStrings
}

func stringToPassport(passportString string) *passport {
	r := regexp.MustCompile(`(byr|iyr|eyr|hgt|hcl|ecl|pid|cid):([^\s]+)`)
	matches := r.FindAllStringSubmatch(passportString, -1)

	pass := &passport{}
	for _, match := range matches {
		switch match[1] {
		case "byr":
			pass.byr = match[2]
		case "iyr":
			pass.iyr = match[2]
		case "eyr":
			pass.eyr = match[2]
		case "hgt":
			pass.hgt = match[2]
		case "hcl":
			pass.hcl = match[2]
		case "ecl":
			pass.ecl = match[2]
		case "pid":
			pass.pid = match[2]
		case "cid":
			pass.cid = match[2]
		}
	}

	return pass
}

func partOne(passports []*passport) (count int) {
	for _, p := range passports {
		if p.hasRequiredProperties() {
			count++
		}
	}
	return
}

func partTwo(passports []*passport) (count int) {
	for _, p := range passports {
		if p.valid() {
			count++
		}
	}
	return
}

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("unable to open file: %v", err)
		return
	}

	inputStrings := parseInput(data)
	passportStrings := inputToPassportStrings(inputStrings)
	passports := make([]*passport, len(passportStrings))
	for i, ps := range passportStrings {
		passports[i] = stringToPassport(ps)
	}

	partOneAnswer := partOne(passports)
	fmt.Printf("Part One: %d\n", partOneAnswer)

	partTwoAnswer := partTwo(passports)
	fmt.Printf("Part Two: %d\n", partTwoAnswer)
}
