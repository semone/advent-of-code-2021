// --- Day 4: Giant Squid ---
package main

import (
	"fmt"
	"utils"
)

func main() {
	input, _ := utils.ReadChunks("./day04/input.txt")
	numbers, boards := parseInput(input)
	part1 := solvePart1(numbers, boards)
	part2 := solvePart2(numbers, boards)
	fmt.Printf("Part 1: %v\n", part1)
	fmt.Printf("Part 2: %v\n", part2)
}

func parseInput(input [][]int) ([]int, [][]int) {
	numbers := input[0]
	boards := input[1:]

	return numbers, boards
}

func solvePart1(numbers []int, boards [][]int) int {
	return playBingo(numbers, boards, true)
}

func solvePart2(numbers []int, boards [][]int) int {
	return playBingo(numbers, boards, false)
}

func playBingo(numbers []int, boards [][]int, firstBoard bool) int {
	m := make(map[int]bool)

	for _, number := range numbers {
		for index := range boards {
			markBoard(boards[index], number)
			if isWinningBoard(boards[index]) {
				if firstBoard {
					return calculateScore(boards[index], number)
				} else {
					m[index] = true
					if len(m) == len(boards) {
						return calculateScore(boards[index], number)
					}
				}
			}
		}

	}

	return 0
}

func markBoard(board []int, number int) {
	for index, val := range board {
		if val == number {
			board[index] = -1
		}
	}
}

func isWinningBoard(board []int) bool {
	var size = 5
	return checkHorizontal(board, size) || checkVertical(board, size)
}

func checkHorizontal(board []int, size int) bool {
	for i := 0; i < len(board); i += size {
		if bingoLine(board[i : i+size]) {
			return true
		}
	}

	return false
}

func checkVertical(board []int, size int) bool {
	for i := 0; i < size; i++ {
		var column []int
		for j := 0; j < size; j++ {
			column = append(column, board[i+size*j])
		}

		if bingoLine(column) {
			return true
		}
	}

	return false
}

func bingoLine(row []int) bool {
	for _, val := range row {
		if val != -1 {
			return false
		}
	}
	return true
}

func calculateScore(board []int, number int) int {
	var sum = 0
	for _, val := range board {
		if val != -1 {
			sum += val
		}
	}
	return sum * number
}
