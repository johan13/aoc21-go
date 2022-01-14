package day01

import "testing"

func TestPart1(t *testing.T) {
	cases := []struct {
		file string
		want int
	}{
		{"example.txt", 7},
		{"input.txt", 1583},
	}

	for _, c := range cases {
		got, err := Part1(c.file)
		if err != nil {
			t.Errorf("Day 01 part 1 failed at %s: %v", c.file, err)
		} else if got != c.want {
			t.Errorf("Day 01 part 1: Expected %d, got %d.", c.want, got)
		}
	}
}

func TestPart2(t *testing.T) {
	cases := []struct {
		file string
		want int
	}{
		{"example.txt", 5},
		{"input.txt", 1627},
	}

	for _, c := range cases {
		got, err := Part2(c.file)
		if err != nil {
			t.Errorf("Day 01 part 2 failed at %s: %v", c.file, err)
		} else if got != c.want {
			t.Errorf("Day 01 part 2: Expected %d, got %d.", c.want, got)
		}
	}
}
