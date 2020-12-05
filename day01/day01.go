package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func parseInput(input []byte) ([]int, error) {
	strings := strings.Split(string(input), "\n")
	ints := make([]int, len(strings))

	for i, stringValue := range strings {
		intValue, err := strconv.Atoi(stringValue)

		if err != nil {
			return nil, fmt.Errorf("Unable to parse value \"%s\" to int", stringValue)
		}

		ints[i] = intValue
	}

	return ints, nil
}

func partOne(values []int) (int, error) {
	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			valueOne, valueTwo := values[i], values[j]
			if valueOne+valueTwo == 2020 {
				return valueOne * valueTwo, nil
			}
		}
	}

	return 0, fmt.Errorf("unable to find two values that sum to 2020")
}

func partTwo(values []int) (int, error) {
	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			for k := j + 1; k < len(values); k++ {
				valueOne, valueTwo, valueThree := values[i], values[j], values[k]
				if valueOne+valueTwo+valueThree == 2020 {
					return valueOne * valueTwo * valueThree, nil
				}
			}
		}
	}

	return 0, fmt.Errorf("unable to find three values that sum to 2020")
}

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("unable to open file: %v", err)
		return
	}

	values, err := parseInput(data)
	if err != nil {
		fmt.Printf("unable to parse input: %v", err)
	}

	partOneAnswer, err := partOne(values)
	if err != nil {
		fmt.Printf("error during part one: %v", err)
	}

	fmt.Printf("Part One: %d\n", partOneAnswer)

	partTwoAnswer, err := partTwo(values)
	if err != nil {
		fmt.Printf("error during part two: %v", err)
	}

	fmt.Printf("Part Two: %d\n", partTwoAnswer)

}
