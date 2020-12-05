package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type traversalMethod struct {
	rows int
	cols int
}

func parseInput(input []byte) []string {
	return strings.Split(string(input), "\n")
}

func traverseSlope(mapData []string, method *traversalMethod) int {
	var (
		trees int
		row   int
		col   int
	)

	mapLength := len(mapData)
	mapWidth := len(mapData[0])

	for row < mapLength {
		if string(mapData[row][col]) == "#" {
			trees++
		}
		row = row + method.rows
		col = (col + method.cols) % mapWidth
	}

	return trees
}

func partOne(mapData []string) int {
	method := &traversalMethod{1, 3}
	return traverseSlope(mapData, method)
}

func partTwo(mapData []string) int {
	methods := []*traversalMethod{
		&traversalMethod{1, 1},
		&traversalMethod{1, 3},
		&traversalMethod{1, 5},
		&traversalMethod{1, 7},
		&traversalMethod{2, 1},
	}

	treeValues := make([]int, len(methods))
	for i, m := range methods {
		treeValues[i] = traverseSlope(mapData, m)
	}

	result := treeValues[0]
	for i := 1; i < len(treeValues); i++ {
		result *= treeValues[i]
	}

	return result
}

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("unable to open file: %v", err)
		return
	}

	mapData := parseInput(data)

	partOneAnswer := partOne(mapData)
	fmt.Printf("Part One: %d\n", partOneAnswer)

	partTwoAnswer := partTwo(mapData)
	fmt.Printf("Part Two: %d\n", partTwoAnswer)
}
