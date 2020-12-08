package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type passwordPolicy struct {
	valueOne  int
	valueTwo  int
	character string
	password  string
}

func parseInput(input []byte) []string {
	return strings.Split(string(input), "\n")
}

func parseValue(value string) (*passwordPolicy, error) {
	r := regexp.MustCompile(`(\d+)-(\d+) ([a-z]): (.*)`)
	matches := r.FindStringSubmatch(value)
	if len(matches) != 5 {
		return nil, fmt.Errorf("unable to parse input \"%s\"", value)
	}

	valueOne, err := strconv.Atoi(matches[1])
	if err != nil {
		return nil, fmt.Errorf("unable to get first int from value \"%s\"", matches[1])
	}

	valueTwo, err := strconv.Atoi(matches[2])
	if err != nil {
		return nil, fmt.Errorf("unable to get second int from value \"%s\"", matches[2])
	}

	return &passwordPolicy{valueOne, valueTwo, matches[3], matches[4]}, nil
}

func partOne(values []string) (int, error) {
	valid := 0

	for _, v := range values {
		policy, err := parseValue(v)
		if err != nil {
			return 0, fmt.Errorf("error parsing value: %v", err)
		}

		r := regexp.MustCompile(policy.character)
		matches := r.FindAllStringIndex(policy.password, -1)
		if len(matches) >= policy.valueOne && len(matches) <= policy.valueTwo {
			valid++
		}
	}

	return valid, nil
}

func partTwo(values []string) (int, error) {
	valid := 0

	for _, v := range values {
		policy, err := parseValue(v)
		if err != nil {
			return 0, fmt.Errorf("error parsing value: %v", err)
		}

		validPosOne := string(policy.password[policy.valueOne-1]) == policy.character
		validPosTwo := string(policy.password[policy.valueTwo-1]) == policy.character
		if validPosOne != validPosTwo {
			valid++
		}
	}

	return valid, nil
}

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("unable to open file: %v\n", err)
		return
	}

	values := parseInput(data)

	partOneAnswer, err := partOne(values)
	if err != nil {
		fmt.Printf("error during part one: %v\n", err)
		return
	}

	fmt.Printf("Part One: %d\n", partOneAnswer)

	partTwoAnswer, err := partTwo(values)
	if err != nil {
		fmt.Printf("error during part two: %v\n", err)
		return
	}

	fmt.Printf("Part Two: %d\n", partTwoAnswer)
}
