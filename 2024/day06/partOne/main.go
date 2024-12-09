package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	UP = [2]int {-1, 0}
	LEFT = [2]int {0, -1}
	RIGHT = [2]int {0, 1}
	DOWN = [2]int {1, 0}
)

type Position struct {
	x int
	y int
}

func main() {

	file, error := os.ReadFile("input.txt")	

	if error != nil {
		fmt.Printf("An error has ocurred while processing the input: %v", error)
		os.Exit(1)
	}

	input := string(file)

	lines := strings.Split(input, "\n")
	matrix := make([][]string, 0)

	var initialPosition Position 
	var direction [2]int
	


	for i, line := range lines {
		cells := strings.Split(line, "")

		for j, cell := range cells {
			if cell == "^" {
				direction = UP
				initialPosition = Position{i, j}
			} else if cell == "<" {
				direction = LEFT
				initialPosition = Position{i, j}
			} else if cell == ">" {
				direction = RIGHT
				initialPosition = Position{i, j}
			} else if cell == "v" {
				direction = DOWN
				initialPosition = Position{i, j}
			}
		} 

		matrix = append(matrix, cells)
	}


	currentPosition := initialPosition
	const blockedLocation = "#"
	const visitedLocation = "X"
	const nonVisitedLocation = "."

	totalLocationsVisited := 1

	for isInsideMatrixBoudaries(matrix, currentPosition) {

		nextPosition := Position {currentPosition.x + direction[0], currentPosition.y + direction[1]}

		if !isInsideMatrixBoudaries(matrix, nextPosition) {
			break
		}

		nextLocation := matrix[nextPosition.x][nextPosition.y]

		if nextLocation == blockedLocation {
			direction = rotateDirectionRight(direction)
			continue
		} else if nextLocation == nonVisitedLocation {
			matrix[nextPosition.x][nextPosition.y] = visitedLocation
			totalLocationsVisited++
		}

		currentPosition = nextPosition
	}

	fmt.Printf("Total visited locations: %v\n", totalLocationsVisited)

}

func isInsideMatrixBoudaries(matrix [][]string, currentPosition Position) bool {
 return currentPosition.x < len(matrix) && currentPosition.x >= 0 && currentPosition.y < len(matrix[0]) && currentPosition.y >= 0
}

func rotateDirectionRight(direction [2]int) [2]int {
	if direction == UP {
		return RIGHT
	} else if direction == RIGHT {
		return DOWN
	} else if direction == DOWN {
		return LEFT
	} else {
		return UP
	}
}