// --- Day 5: Hydrothermal Venture ---
package main

import (
	"fmt"
	"math"
	"strings"
	"utils"
)

type Point struct {
	x int
	y int
}

type Line struct {
	start Point
	end   Point
}

func main() {
	input, _ := utils.ReadLines("./day05/input.txt", parseLines)
	lines := parseInput(input)
	part1 := solvePart1(lines)
	part2 := solvePart2(lines)
	fmt.Printf("Part 1: %v\n", part1)
	fmt.Printf("Part 2: %v\n", part2)

}

func parseLines(str string) (string, error) {
	return str, nil
}

func parseInput(input []string) []Line {
	var lines []Line
	for _, val := range input {
		split := strings.Split(val, "->")

		stringStartPoint := strings.Trim(split[0], " ")
		stringEndPoint := strings.Trim(split[1], " ")

		var startPointX, startPointY, endPointX, endPointY int

		fmt.Sscanf(stringStartPoint, "%d,%d", &startPointX, &startPointY)
		fmt.Sscanf(stringEndPoint, "%d,%d", &endPointX, &endPointY)

		line := Line{start: Point{x: startPointX, y: startPointY}, end: Point{x: endPointX, y: endPointY}}

		lines = append(lines, line)

	}
	return lines
}

func solvePart1(lines []Line) int {
	return runProgram(lines, shouldBeConsideredPart1)
}

func solvePart2(lines []Line) int {
	return runProgram(lines, shouldBeConsideredPart2)
}

func runProgram(lines []Line, consideredFunc func(Line) bool) int {

	diagram := make(map[string]int)

	for _, line := range lines {
		if consideredFunc(line) {
			coverPoints(line, diagram)
		}
	}
	return calculateScore(diagram)
}

func calculateScore(diagram map[string]int) int {
	sum := 0

	for _, element := range diagram {
		if element > 1 {
			sum += 1
		}
	}
	return sum
}

func coverPoints(line Line, diagram map[string]int) {
	factorX := 1
	if line.start.x > line.end.x {
		factorX = -1
	}
	if isVertical(line) {
		factorY := 1
		if line.start.y > line.end.y {
			factorY = -1
		}
		for i := 0; i <= int(math.Abs(float64(line.start.y)-float64(line.end.y))); i++ {
			coverPoint(Point{x: line.start.x, y: line.start.y + i*factorY}, diagram)
		}
	} else if isHoriozntal(line) {
		for i := 0; i <= int(math.Abs(float64(line.start.x)-float64(line.end.x))); i++ {
			coverPoint(Point{x: line.start.x + factorX*i, y: line.start.y}, diagram)
		}
	} else if isDiagonal(line) {
		length := int(math.Abs(float64(line.start.x) - float64(line.end.x)))
		if positiveSlope(line) {
			for i := 0; i <= length; i++ {
				coverPoint(Point{x: line.start.x + factorX*i, y: line.start.y + factorX*i}, diagram)
			}
		}
		if negativeSlope(line) {
			for i := 0; i <= length; i++ {
				coverPoint(Point{x: line.start.x + factorX*i, y: line.start.y - factorX*i}, diagram)
			}
		}
	} else {
		fmt.Println("Not an accepted line!")
	}

}

func coverPoint(point Point, diagram map[string]int) {
	diagramKey := fmt.Sprintf("%d,%d", point.x, point.y)
	diagram[diagramKey] += 1
}

func isHoriozntal(line Line) bool {
	return line.start.y == line.end.y
}

func isVertical(line Line) bool {
	return line.start.x == line.end.x
}

func negativeSlope(line Line) bool {
	return line.end.y-line.start.y == line.start.x-line.end.x
}

func positiveSlope(line Line) bool {
	return line.end.y-line.start.y == line.end.x-line.start.x
}

func isDiagonal(line Line) bool {
	return positiveSlope(line) || negativeSlope(line)
}

func shouldBeConsideredPart1(line Line) bool {
	return isHoriozntal(line) || isVertical(line)
}

func shouldBeConsideredPart2(line Line) bool {
	return shouldBeConsideredPart1(line) || isDiagonal(line)
}
