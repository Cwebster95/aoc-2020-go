package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

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

func partOne(values []int) int {
	currentRating := 0
	differences := make(map[int]int)

	for _, v := range values {
		difference := v - currentRating
		differences[difference]++
		currentRating = v
	}
	differences[3]++

	return differences[1] * differences[3]
}

func buildToEnd(values []int, currentValue int) int {
	if len(values) == 0 {
		return 1
	}

	possiblePathCount := 0
	for i := 0; i < len(values) && values[i]-currentValue <= 3; i++ {
		possiblePathCount += buildToEnd(values[i+1:], values[i])
	}

	return possiblePathCount
}

func partTwo(values []int) int {
	splits := make([][]int, 1)
	currentValue, currentSplit := 0, 0

	for i := 0; i < len(values); i++ {
		if values[i]-currentValue == 3 {
			splits = append(splits, make([]int, 0))
			currentSplit++
		}

		splits[currentSplit] = append(splits[currentSplit], values[i])
		currentValue = values[i]
	}

	currentPaths := 1
	for i, split := range splits {
		currentValue := 0
		if i > 0 {
			currentValue = splits[i-1][len(splits[i-1])-1]
		}

		currentPaths *= buildToEnd(split, currentValue)
	}
	return currentPaths
}

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("unable to open file: %v\n", err)
		return
	}

	values, err := parseInput(data)
	sort.Ints(values)
	if err != nil {
		fmt.Printf("error parsing input data: %v\n", err)
		return
	}

	partOneAnswer := partOne(values)
	fmt.Printf("Part One: %d\n", partOneAnswer)

	partTwoAnswer := partTwo(values)
	fmt.Printf("Part Two: %d\n", partTwoAnswer)
}
