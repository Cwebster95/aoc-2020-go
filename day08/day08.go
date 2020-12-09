package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func parseInput(input []byte) ([]*instruction, error) {
	instructionStrings := strings.Split(string(input), "\n")
	instructions := make([]*instruction, len(instructionStrings))

	for i, s := range instructionStrings {
		currentInstruction, err := newInstruction(s)
		if err != nil {
			return nil, err
		}

		instructions[i] = currentInstruction
	}

	return instructions, nil
}

func partOne(instructions []*instruction) (int, error) {
	comp := newComputer(instructions)
	comp.run()

	if comp.exitCode == endOfProgram {
		return 0, fmt.Errorf("program ended before repeated instruction")
	}

	return comp.accumulator, nil
}

func partTwo(instructions []*instruction) (int, error) {
	for i, in := range instructions {
		if in.operand == "jmp" || in.operand == "nop" {
			newInstructions := make([]*instruction, len(instructions))
			copy(newInstructions, instructions)
			newInstructions[i] = in.switchOperand()

			comp := newComputer(newInstructions)
			comp.run()

			if comp.exitCode == endOfProgram {
				return comp.accumulator, nil
			}
		}
	}

	return 0, fmt.Errorf("unable to find which instruction to switch")
}

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("unable to open file: %v\n", err)
		return
	}

	instructions, err := parseInput(data)
	if err != nil {
		fmt.Printf("error parsing input data: %v\n", err)
		return
	}

	partOneAnswer, err := partOne(instructions)
	if err != nil {
		fmt.Printf("error during part one: %v\n", err)
		return
	}

	fmt.Printf("Part One: %d\n", partOneAnswer)

	partTwoAnswer, err := partTwo(instructions)
	if err != nil {
		fmt.Printf("error during part two: %v\n", err)
		return
	}

	fmt.Printf("Part Two: %d\n", partTwoAnswer)
}
