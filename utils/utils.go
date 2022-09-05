package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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

func ReadChunks(path string) ([][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]int
	scanner := bufio.NewScanner(file)
	var board []int

	for scanner.Scan() {
		value := scanner.Text()
		if strings.Contains(value, ",") {
			vals := strings.Split(value, ",")
			intVals := make([]int, len(vals))
			for i := range vals {
				intVals[i], _ = strconv.Atoi(vals[i])
			}
			lines = append(lines, intVals)

		} else {
			if len(value) == 0 {
				if len(board) > 0 {
					lines = append(lines, board)
					board = nil
				}
			} else {
				vals := strings.Fields(value)
				intVals := make([]int, len(vals))
				for i := range vals {
					intVals[i], _ = strconv.Atoi(vals[i])
				}
				board = append(board, intVals...)
			}
		}
	}
	lines = append(lines, board)
	return lines, scanner.Err()
}
