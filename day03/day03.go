// --- Day 3: Binary Diagnostic ---
package main

import (
	"fmt"
	"strconv"
	"utils"
)

func main() {
	input, _ := utils.ReadLines("./day03/test-input.txt", parseLines)
	part1 := solvePart1(input)
	part2 := solvePart2(input)
	fmt.Printf("Part 1: %v\n", part1)
	fmt.Printf("Part 2: %v\n", part2)
}

func parseLines(str string) (string, error) {
	return str, nil
}

func solvePart1(lines []string) int {
	rating1, rating2 := calculateRatingPart1(lines)
	gamma, _ := strconv.ParseInt(rating1, 2, 64)
	epsilon, _ := strconv.ParseInt(rating2, 2, 64)

	return int(gamma) * int(epsilon)
}

func solvePart2(lines []string) int {

	oxygen, _ := strconv.ParseInt(calculateRatingPart2(lines, true), 2, 64)
	scrubber, _ := strconv.ParseInt(calculateRatingPart2(lines, false), 2, 64)

	return int(oxygen) * int(scrubber)
}

func calculateRatingPart1(lines []string) (string, string) {
	positions := len(lines[0])

	ratingGamma := make([]rune, positions)
	ratingEpsilon := make([]rune, positions)

	for i := 0; i < positions; i++ {
		ratingGamma[i] = rune(getBitCriteria(lines, i, true))
		ratingEpsilon[i] = rune(getBitCriteria(lines, i, false))
	}

	return string(ratingGamma), string(ratingEpsilon)
}

func numberOfPositivesAtPostion(lines []string, position int) int {
	positives := 0

	for _, value := range lines {
		if value[position] == '1' {
			positives++
		}
	}

	return positives
}

func calculateRatingPart2(lines []string, mostCommon bool) string {
	total := len(lines)
	linesLeft := total
	iteration := 0
	currentLines := lines

	for linesLeft > 1 {
		criteria := getBitCriteria(currentLines, iteration, mostCommon)
		keptLines := filterByCriteria(currentLines, iteration, criteria)
		iteration++
		linesLeft = len(keptLines)
		currentLines = keptLines
	}

	return currentLines[0]
}

func getBitCriteria(lines []string, position int, oxygen bool) byte {
	positives := numberOfPositivesAtPostion(lines, position)
	negatives := len(lines) - positives
	if oxygen {
		if positives >= negatives {
			return '1'
		} else {
			return '0'
		}
	} else {
		if positives >= negatives {
			return '0'
		} else {
			return '1'
		}
	}

}

func filterByCriteria(lines []string, iteration int, criteria byte) []string {
	var keptLines []string
	for _, value := range lines {
		if value[iteration] == criteria {
			keptLines = append(keptLines, value)
		}
	}
	return keptLines
}
