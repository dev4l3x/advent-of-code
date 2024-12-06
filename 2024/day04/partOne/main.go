package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	UP = [2]int {1, 0} 
	DOWN = [2]int {-1, 0} 
	LEFT = [2]int {0, -1}
	RIGHT = [2]int {0, 1}
	UPPER_LEFT_DIAGONAL = [2]int {-1, -1}
	UPPER_RIGHT_DIAGONAL = [2]int {-1, 1}
	DOWN_LEFT_DIAGONAL = [2]int {1, -1}
	DOWN_RIGHT_DIAGONAL = [2]int {1, 1}
)

var directions [8][2]int = [8][2]int {
	UP, DOWN, LEFT, RIGHT, UPPER_LEFT_DIAGONAL, UPPER_RIGHT_DIAGONAL, DOWN_LEFT_DIAGONAL, DOWN_RIGHT_DIAGONAL, 
}

func main() {

	file, error := os.ReadFile("input.txt")	

	if error != nil {
		fmt.Printf("An error has ocurred while processing the input: %v", error)
		os.Exit(1)
	}

	input := string(file)

	lines := strings.Split(input, "\n")

	// Build the matrix to search for words with XMAS

	matrix := make([][]string, 0)
	for _, line := range lines {
		row := strings.Split(line, "")
		matrix = append(matrix, row)
	}

	// Search for X letters in the matrix, that starts the pattern
	ocurrences := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == "X" {
				ocurrences += countXmasOcurrences(matrix, i, j)
			}
		}
	}

	fmt.Printf("The sum of all ocurrences is: %v\n", ocurrences)

}

func countXmasOcurrences(matrix [][]string, i int, j int) int {
	
	ocurrences := 0
	
	for _, direction := range directions {
		ocurrences += countXmasOcurrencesInDirection(matrix, i, j, direction)
	}


	return ocurrences
}

func countXmasOcurrencesInDirection(matrix [][]string, i int, j int, direction [2]int) int {

	const xmasSize = 4

	if i + (xmasSize - 1) * direction[0] < 0 && i + (xmasSize - 1) * direction[0] >= len(matrix) {
		return 0
	} else if j 
	
	word := make([]string, 0)

	for k := 0 ; k < xmasSize ; k++ {
		letter := matrix[i + k * direction[0]][j + k * direction[1]]
		word = append(word, letter)
	} 

	joinedWord := strings.Join(word, "")

	if joinedWord == "XMAS" {
		return 1
	}

	return 0
}

func parseNumber(number string) int {
	
	n, err := strconv.Atoi(number)

	if err != nil {
		log.Fatalf("An error has ocurred while parsing the number: %v", err)
	}

	return n
}