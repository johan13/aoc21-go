package day03

import (
	"strconv"

	"github.com/johan13/aoc21-go/pkg/common"
)

func Part1(path string) (int, error) {
	input, err := common.ReadInput(path, func(line string) (string, error) { return line, nil })
	if err != nil {
		return 0, err
	}

	gamma, epsilon := 0, 0
	for pos := range input[0] {
		numOnesAtPos := 0
		for _, line := range input {
			if line[pos] == '1' {
				numOnesAtPos++
			}
		}
		if numOnesAtPos > len(input)/2 {
			gamma = gamma<<1 | 1
			epsilon = epsilon << 1
		} else {
			gamma = gamma << 1
			epsilon = epsilon<<1 | 1
		}
	}

	return gamma * epsilon, nil
}

func Part2(path string) (int, error) {
	input, err := common.ReadInput(path, func(line string) (string, error) { return line, nil })
	if err != nil {
		return 0, err
	}

	var o2, co2 int
	if o2, err = findRating(input, true, 0); err != nil {
		return 0, err
	}
	if co2, err = findRating(input, false, 0); err != nil {
		return 0, err
	}
	return o2 * co2, nil
}

func findRating(input []string, mostCommon bool, pos int) (int, error) {
	if len(input) == 1 {
		rating, error := strconv.ParseInt(input[0], 2, 0)
		if error != nil {
			return 0, error
		}
		return int(rating), nil
	}
	numOnesAtPos := 0
	for _, line := range input {
		if line[pos] == '1' {
			numOnesAtPos++
		}
	}
	digitToKeep := '0'
	if mostCommon == (numOnesAtPos*2 >= len(input)) {
		digitToKeep = '1'
	}
	var filtered []string
	for _, line := range input {
		if line[pos] == byte(digitToKeep) {
			filtered = append(filtered, line)
		}
	}
	return findRating(filtered, mostCommon, pos+1)
}
