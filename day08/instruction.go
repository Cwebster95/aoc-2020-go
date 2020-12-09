package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type instruction struct {
	argument  int
	operand   string
	processed bool
}

func (i *instruction) String() string {
	processedString := ""
	if i.processed {
		processedString = "!"
	}
	return fmt.Sprintf("%s(%d)%s", i.operand, i.argument, processedString)
}

func (i *instruction) switchOperand() *instruction {
	switch i.operand {
	case "jmp":
		return &instruction{argument: i.argument, operand: "nop", processed: false}
	case "nop":
		return &instruction{argument: i.argument, operand: "jmp", processed: false}
	default:
		return &instruction{argument: i.argument, operand: i.operand, processed: false}
	}
}

func newInstruction(instructionString string) (*instruction, error) {
	r := regexp.MustCompile(`^(acc|jmp|nop) ([+-]\d+)$`)
	if !r.Match([]byte(instructionString)) {
		return nil, fmt.Errorf("unable to parse instruction string \"%s\"", instructionString)
	}

	matches := r.FindStringSubmatch(instructionString)
	argumentValue, err := strconv.Atoi(matches[2])
	if err != nil {
		return nil, fmt.Errorf("unable to parse argument value \"%s\"", matches[2])
	}

	return &instruction{argument: argumentValue, operand: matches[1], processed: false}, nil
}
