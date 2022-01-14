package day02

import (
	"regexp"
	"strconv"

	"github.com/johan13/aoc21-go/pkg/common"
)

func Part1(path string) (int, error) {
	input, err := common.ReadInput(path, parseLine)
	if err != nil {
		return 0, err
	}

	x, y := 0, 0
	for _, move := range input {
		switch move.dir {
		case "forward":
			x += move.units
		case "down":
			y += move.units
		case "up":
			y -= move.units
		}
	}
	return x * y, nil
}

func Part2(path string) (int, error) {
	input, err := common.ReadInput(path, parseLine)
	if err != nil {
		return 0, err
	}

	x, y, aim := 0, 0, 0
	for _, move := range input {
		switch move.dir {
		case "forward":
			x += move.units
			y += aim * move.units
		case "down":
			aim += move.units
		case "up":
			aim -= move.units
		}
	}
	return x * y, nil
}

type Move struct {
	dir   string
	units int
}

var re = regexp.MustCompile(`^(forward|down|up) (\d+)$`)

func parseLine(line string) (Move, error) {
	matches := re.FindStringSubmatch(line)
	units, err := strconv.Atoi(matches[2])
	if err != nil {
		return Move{}, err
	}
	return Move{matches[1], units}, nil
}
