package utils

import (
	"bufio"
	"os"
)

type Line interface {
	string | int | int64
}

func ReadLines[T Line](path string, fun func(string) (T, error)) ([]T, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []T
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		value, _ := fun(scanner.Text())
		lines = append(lines, value)
	}

	return lines, scanner.Err()
}
