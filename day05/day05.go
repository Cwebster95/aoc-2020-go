package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"sort"
	"strings"
)

type ticket struct {
	rowSelector string
	colSelector string
}

type seat struct {
	row int
	col int
	id  int
}

func (t *ticket) String() string {
	return fmt.Sprintf("Row: %s\tCol: %s", t.rowSelector, t.colSelector)
}

func (s *seat) String() string {
	return fmt.Sprintf("Row: %d  \tCol: %d  \tID: %d", s.row, s.col, s.id)
}

func (t *ticket) getSeat() (*seat, error) {
	row, err := calculatePosition(t.rowSelector)
	if err != nil {
		return nil, fmt.Errorf("unable to calculate row: %v", err)
	}

	col, err := calculatePosition(t.colSelector)
	if err != nil {
		return nil, fmt.Errorf("unable to calculate col: %v", err)
	}

	return &seat{
		row: row,
		col: col,
		id:  ((row * 8) + col),
	}, nil
}

func calculatePosition(selector string) (int, error) {
	var min byte = 0
	var max byte = byte(math.Pow(2, float64(len(selector))) - 1)

	for i := 0; i < len(selector); i++ {
		power := float64(len(selector)-i) - 1

		switch selector[i] {
		case 'B', 'R': // Take upper half, so remove lower half
			min = min | byte(math.Pow(2, power))
		case 'F', 'L': // Take lower half, so remove upper half
			max = max & byte(127-math.Pow(2, power))
		default:
			return 0, fmt.Errorf("unknown selector: %s", string(selector[i]))
		}
	}

	if min != max {
		return 0, fmt.Errorf("error during calculation, min (%d) != max (%d)", min, max)
	}

	return int(min), nil
}

func parseInput(input []byte) ([]*ticket, error) {
	inputStrings := strings.Split(string(input), "\n")
	tickets := make([]*ticket, len(inputStrings))
	r := regexp.MustCompile(`^([BF]+)([LR]+)$`)

	for i, ts := range inputStrings {
		match := r.FindStringSubmatch(ts)
		if len(match) != 3 {
			return nil, fmt.Errorf("unable to parse input string \"%s\"", ts)
		}

		tickets[i] = &ticket{rowSelector: match[1], colSelector: match[2]}
	}

	return tickets, nil
}

func partOne(tickets []*ticket) (int, error) {
	highestID := -1 // Can have id of 0

	for _, t := range tickets {
		seat, err := t.getSeat()
		if err != nil {
			return -1, err
		}

		if seat.id > highestID {
			highestID = seat.id
		}
	}

	return highestID, nil
}

func partTwo(tickets []*ticket) (int, error) {
	seatIDs := make([]int, len(tickets))

	for i, t := range tickets {
		seat, err := t.getSeat()
		if err != nil {
			return -1, err
		}

		seatIDs[i] = seat.id
	}

	sort.Ints(seatIDs)

	for i := 0; i < len(seatIDs)-1; i++ {
		if seatIDs[i+1] != seatIDs[i]+1 {
			return seatIDs[i] + 1, nil
		}
	}

	return -1, fmt.Errorf("unable to find a missing seat ID")
}

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("unable to open file: %v\n", err)
		return
	}

	tickets, err := parseInput(data)
	if err != nil {
		fmt.Printf("error parsing input data: %v\n", err)
		return
	}

	partOneAnswer, err := partOne(tickets)
	if err != nil {
		fmt.Printf("error during part one: %v\n", err)
		return
	}

	fmt.Printf("Part One: %d\n", partOneAnswer)

	partTwoAnswer, err := partTwo(tickets)
	if err != nil {
		fmt.Printf("error during part two: %v\n", err)
		return
	}

	fmt.Printf("Part Two: %d\n", partTwoAnswer)
}
