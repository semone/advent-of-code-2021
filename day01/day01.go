// --- Day 1: Sonar Sweep ---

package main

import (
	"fmt"
	"strconv"
	"utils"
)

func main() {
	// TODO fix path/workspace/module
	input, _ := utils.ReadLines("./day01/input.txt", strconv.Atoi)
	part1 := solvePart1(input)
	part2 := solvePart2(input)
	fmt.Printf("Part 1: %v\n", part1)
	fmt.Printf("Part 2: %v\n", part2)
}

func solvePart1(lines []int) int {
	sum := 0
	var prev int
	for _, value := range lines {
		current := value

		if current > prev && prev != 0 {
			sum += 1
		}
		prev = current
	}
	return sum
}

func solvePart2(lines []int) int {
	threeLines := threeLines(lines)
	sum := solvePart1(threeLines)
	return sum
}

func threeLines(lines []int) []int {
	var threeLines []int

	for i := 0; i < len(lines)-2; i++ {
		r := lines[i : i+3]
		result := 0
		for _, v := range r {
			result += v
		}
		threeLines = append(threeLines, result)

	}
	return threeLines
}
