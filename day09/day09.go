package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const preambleLength = 25

func parseInput(input []byte) ([]int, error) {
	stringValues := strings.Split(string(input), "\n")
	intValues := make([]int, len(stringValues))

	for i, s := range stringValues {
		intValue, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("unable to parse input value \"%s\"", s)
		}

		intValues[i] = intValue
	}

	return intValues, nil
}

func checkSum(preambleValues []int, target int) bool {
	for i := 0; i < len(preambleValues)-1; i++ {
		for j := i + 1; j < len(preambleValues); j++ {
			if preambleValues[i]+preambleValues[j] == target {
				return true
			}
		}
	}

	return false
}

func firstInvalid(values []int, preamble int) (int, error) {
	for i := preamble; i < len(values); i++ {
		currentValue := values[i]
		preambleValues := make([]int, preamble)
		copy(preambleValues, values[i-preamble:i])

		if !checkSum(preambleValues, currentValue) {
			return currentValue, nil
		}
	}

	return 0, fmt.Errorf("could not find a value that wasn't summable from its' preamble")
}

func sum(values []int) int {
	result := 0

	for _, v := range values {
		result += v
	}

	return result
}

func findConsecutiveSet(values []int, target int) ([]int, error) {
	for groupSize := 2; groupSize <= len(values); groupSize++ {
		for start := 0; start+groupSize <= len(values); start++ {
			checking := values[start : start+groupSize]

			if sum(checking) == target {
				validSet := make([]int, len(checking))
				copy(validSet, checking)

				return validSet, nil
			}
		}
	}

	return nil, fmt.Errorf("unable to find a set of consecutive numbers adding to %d", target)
}

func addMinMax(values []int) int {
	min, max := values[0], values[0]

	for _, v := range values {
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}

	return min + max
}

func partOne(values []int) (int, error) {
	return firstInvalid(values, preambleLength)
}

func partTwo(values []int) (int, error) {
	targetSum, err := firstInvalid(values, preambleLength)
	if err != nil {
		return 0, err
	}

	validSet, err := findConsecutiveSet(values, targetSum)
	if err != nil {
		return 0, err
	}

	return addMinMax(validSet), nil
}

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("unable to open file: %v\n", err)
		return
	}

	values, err := parseInput(data)
	if err != nil {
		fmt.Printf("error parsing input data: %v\n", err)
		return
	}

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
