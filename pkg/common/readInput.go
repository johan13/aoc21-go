package common

import (
	"bufio"
	"os"
)

func ReadInput[T any](path string, parser func(string) (T, error)) ([]T, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var input []T
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value, err := parser(scanner.Text())
		if err != nil {
			return nil, err
		}
		input = append(input, value)
	}
	return input, scanner.Err()
}
