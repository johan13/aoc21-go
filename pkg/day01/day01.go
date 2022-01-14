package day01

import (
	"strconv"

	"github.com/johan13/aoc21-go/pkg/common"
)

func Part1(path string) (int, error) {
	return countIncreases(path, 1)
}

func Part2(path string) (int, error) {
	return countIncreases(path, 3)
}

func countIncreases(path string, windowSize int) (int, error) {
	input, err := common.ReadInput(path, strconv.Atoi)
	if err != nil {
		return 0, err
	}
	increases := 0
	for i := 0; i < len(input)-windowSize; i++ {
		if input[i+windowSize] > input[i] {
			increases++
		}
	}
	return increases, nil
}
