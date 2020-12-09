package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// internalBag refers to a type of bag that can be stored inside another
type internalBag struct {
	name  string
	count int
}

func parseInput(input []byte) (map[string][]*internalBag, error) {
	inputStrings := strings.Split(string(input), "\n")
	lineRegex := regexp.MustCompile(`^([a-zA-Z ]+) bags contain (.*).$`)
	internalBagRegex := regexp.MustCompile(`(\d+) ([a-zA-Z ]+) bags?`)
	externalBags := make(map[string][]*internalBag)

	for _, s := range inputStrings {
		matches := lineRegex.FindStringSubmatch(s)
		if len(matches) != 3 {
			return nil, fmt.Errorf("error parsing input string \"%s\"", s)
		}

		key, internalBagString := matches[1], matches[2]
		if internalBagString == "no other bags" {
			externalBags[key] = make([]*internalBag, 0)
		} else {
			internalMatches := internalBagRegex.FindAllStringSubmatch(internalBagString, -1)
			if internalMatches == nil {
				return nil, fmt.Errorf("error parsing internal bags for \"%s\"", s)
			}

			externalBags[key] = make([]*internalBag, len(internalMatches))
			for i, m := range internalMatches {
				count, err := strconv.Atoi(m[1])
				if err != nil {
					return nil, fmt.Errorf("error parsing internal bag count: %s for \"%s\"", m[1], m[2])
				}
				externalBags[key][i] = &internalBag{name: m[2], count: count}
			}
		}
	}

	return externalBags, nil
}

func findTarget(bags map[string][]*internalBag, currentBag string, targetBag string) bool {
	if len(bags[currentBag]) == 0 {
		return false
	}

	for _, internal := range bags[currentBag] {
		if internal.name == targetBag {
			return true
		}
	}

	findInside := false
	for _, internal := range bags[currentBag] {
		findInside = findInside || findTarget(bags, internal.name, targetBag)
	}
	return findInside
}

func countInternal(bags map[string][]*internalBag, currentBag *internalBag) int {
	if len(bags[currentBag.name]) == 0 {
		return 0
	}

	count := 0
	for _, innerBag := range bags[currentBag.name] {
		count += innerBag.count + (innerBag.count * countInternal(bags, innerBag))
	}
	return count
}

func partOne(bags map[string][]*internalBag, targetBag string) int {
	count := 0

	for externalBag := range bags {
		if externalBag != targetBag {
			if findTarget(bags, externalBag, targetBag) {
				count++
			}
		}
	}

	return count
}

func partTwo(bags map[string][]*internalBag, startingBag string) int {
	count := 0

	for _, innerBag := range bags[startingBag] {
		count += innerBag.count + (innerBag.count * countInternal(bags, innerBag))
	}

	return count
}

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("unable to open file: %v\n", err)
		return
	}

	bags, err := parseInput(data)
	if err != nil {
		fmt.Printf("error parsing input data: %v\n", err)
		return
	}

	partOneAnswer := partOne(bags, "shiny gold")
	fmt.Printf("Part One: %d\n", partOneAnswer)

	partTwoAnswer := partTwo(bags, "shiny gold")
	fmt.Printf("Part Two: %d\n", partTwoAnswer)
}
