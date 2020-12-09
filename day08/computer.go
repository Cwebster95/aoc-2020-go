package main

import "fmt"

type computer struct {
	accumulator        int
	exitCode           int
	halted             bool
	instructionPointer int
	instructions       []*instruction
}

const (
	endOfProgram        = 0
	repeatedInstruction = 1
)

func (c *computer) String() string {
	return fmt.Sprintf("instruction pointer: %d\naccumulator value: %d\nhalted: %v\nexit code: %d", c.instructionPointer, c.accumulator, c.halted, c.exitCode)
}

func (c *computer) halt(exit int) {
	c.halted = true
	c.exitCode = exit
}

func (c *computer) nextInstruction() {
	if c.instructionPointer >= len(c.instructions) {
		c.halt(endOfProgram)
		return
	}

	currentInstruction := c.instructions[c.instructionPointer]
	if currentInstruction.processed {
		c.halt(repeatedInstruction)
		return
	}

	switch currentInstruction.operand {
	case "acc":
		c.accumulator += currentInstruction.argument
		c.instructionPointer++
	case "jmp":
		c.instructionPointer += currentInstruction.argument
	case "nop":
		c.instructionPointer++
	}
	currentInstruction.processed = true
}

func (c *computer) run() int {
	c.accumulator = 0
	c.exitCode = 0
	c.halted = false
	c.instructionPointer = 0

	for !c.halted {
		c.nextInstruction()
	}

	c.resetInstructions()

	return c.exitCode
}

func (c *computer) resetInstructions() {
	for _, in := range c.instructions {
		in.processed = false
	}
}

func newComputer(instructions []*instruction) *computer {
	return &computer{instructions: instructions}
}
