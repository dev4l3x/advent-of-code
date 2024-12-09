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

const (
	BLOCKED_LOCATION = "#"
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
	var initialDirection [2]int
	


	for i, line := range lines {
		cells := strings.Split(line, "")

		for j, cell := range cells {
			if cell == "^" {
				initialDirection = UP
				initialPosition = Position{i, j}
			} else if cell == "<" {
				initialDirection = LEFT
				initialPosition = Position{i, j}
			} else if cell == ">" {
				initialDirection = RIGHT
				initialPosition = Position{i, j}
			} else if cell == "v" {
				initialDirection = DOWN
				initialPosition = Position{i, j}
			}
		} 

		matrix = append(matrix, cells)
	}


	currentPosition := initialPosition
	direction := initialDirection
	const visitedLocation = "X"
	const nonVisitedLocation = "."

	totalLoopOptions := 0

	for isInsideMatrixBoudaries(matrix, currentPosition) {

		nextPosition := Position {currentPosition.x + direction[0], currentPosition.y + direction[1]}

		if !isInsideMatrixBoudaries(matrix, nextPosition) {
			break
		}

		nextLocation := matrix[nextPosition.x][nextPosition.y]

		if nextPosition != initialPosition && nextLocation == nonVisitedLocation && isInLoop(matrix, currentPosition, direction, nextPosition) {
			totalLoopOptions++
		}

		currentLocation := matrix[currentPosition.x][currentPosition.y]

		if nextLocation == BLOCKED_LOCATION {
			direction = rotateDirectionRight(direction)
			matrix[currentPosition.x][currentPosition.y] = "+"
			continue
		} else { 
			if currentLocation == "." && (direction == LEFT || direction == RIGHT) {
				matrix[currentPosition.x][currentPosition.y] = "-"
			} else if currentLocation == "." && (direction == UP || direction == DOWN){
				matrix[currentPosition.x][currentPosition.y] = "|"
			}
		}

		currentPosition = nextPosition
	}


	fmt.Printf("Total visited locations: %v\n", totalLoopOptions)

}

func isInLoop(matrix [][]string, initialPosition Position, direction [2]int, optionalBlockedPosition Position) bool {

	visitedNodes := make(map[Position]map[[2]int]bool, 0)

	//fmt.Printf("Checking loop for position: %v\n", initialPosition)
	currentPosition := initialPosition

	for isInsideMatrixBoudaries(matrix, currentPosition) {

		visitedDirections, visitedPosition := visitedNodes[currentPosition]
		_, visitedDirection := visitedDirections[direction]

		if visitedPosition && visitedDirection {
			fmt.Printf("Loop at position: %v with blocked: %v\n", currentPosition, optionalBlockedPosition)
			return true
		}

		if !visitedPosition {
			visitedNodes[currentPosition] = map[[2]int]bool {direction: true}
		} else {
			visitedNodes[currentPosition][direction] = true
		}

		nextPosition := Position {currentPosition.x + direction[0], currentPosition.y + direction[1]}

		if !isInsideMatrixBoudaries(matrix, nextPosition) {
			break
		}

		nextLocation := matrix[nextPosition.x][nextPosition.y]

		if nextLocation == BLOCKED_LOCATION || nextPosition == optionalBlockedPosition {
			direction = rotateDirectionRight(direction)
			continue
		} 		
		
		currentPosition = nextPosition
	}

	return false
}

func printMatrix(matrix [][]string, position Position) {
	for x, line := range matrix {
		for y, col := range line {
			if position.x == x && position.y == y {
				fmt.Printf("C ")
			} else {
				fmt.Printf("%v ", col)
			}
		}
		fmt.Printf("\n")
	}

	fmt.Printf("\n\n\n")
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