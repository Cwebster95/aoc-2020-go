package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type answers struct {
	answerString string
	count        int
}

func parseInput(input []byte) []*answers {
	tempStrings := strings.Split(string(input), "\n")
	tempString := strings.Join(tempStrings, "|")
	tempStrings = strings.Split(tempString, "||")

	answerGroups := make([]*answers, len(tempStrings))
	for i, s := range tempStrings {
		answerGroups[i] = &answers{
			answerString: strings.ReplaceAll(s, "|", ""),
			count:        strings.Count(s, "|") + 1,
		}
	}
	return answerGroups
}

func partOne(answerGroups []*answers) int {
	total := 0

	for _, a := range answerGroups {
		current := 0

		for pos, ch := range a.answerString {
			found := false

			for i := 0; i < pos; i++ {
				if a.answerString[i] == byte(ch) {
					found = true
				}
			}

			if !found {
				current++
			}
		}

		total += current
	}

	return total
}

func partTwo(answerGroups []*answers) int {
	total := 0

	for _, a := range answerGroups {
		current := 0

		for pos, ch := range a.answerString {
			found := false

			for i := 0; i < pos; i++ {
				if a.answerString[i] == byte(ch) {
					found = true
				}
			}

			if !found && strings.Count(a.answerString, string(ch)) == a.count {
				current++
			}
		}

		total += current
	}

	return total
}

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("unable to open file: %v", err)
		return
	}

	input := parseInput(data)

	partOneAnswer := partOne(input)
	fmt.Printf("Part One: %d\n", partOneAnswer)

	partTwoAnswer := partTwo(input)
	fmt.Printf("Part Two: %d\n", partTwoAnswer)
}
