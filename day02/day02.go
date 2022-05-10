// --- Day 2: Dive! ---

package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

func main() {
	// TODO fix path/workspace/module
	input, _ := utils.ReadLines("./day02/input.txt", parseLines)
	part1 := solvePart1(input)
	part2 := solvePart2(input)
	fmt.Printf("Part 1: %v\n", part1)
	fmt.Printf("Part 2: %v\n", part2)
}

func parseLines(str string) (string, error) {
	return str, nil
}

func solve(lines []string, isSecondPart bool) int {
	var horizontal int
	var depth int
	var trueDepth int
	for _, value := range lines {
		keyValue := strings.Fields((value))
		intVal, _ := strconv.Atoi(keyValue[1])
		if keyValue[0] == "up" {
			depth -= intVal
		} else if keyValue[0] == "down" {
			depth += intVal
		} else {
			horizontal += intVal
			trueDepth += depth * intVal
		}
	}
	if isSecondPart {
		return trueDepth * horizontal
	} else {
		return depth * horizontal
	}
}

func solvePart1(lines []string) int {
	return solve(lines, false)
}

func solvePart2(lines []string) int {
	return solve(lines, true)
}
