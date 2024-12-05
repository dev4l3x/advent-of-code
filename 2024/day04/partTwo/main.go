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
			if matrix[i][j] == "A" && isXmasShape(matrix, i, j) {
				ocurrences++ 
			}
		}
	}

	fmt.Printf("The sum of all ocurrences is: %v\n", ocurrences)

}

func isXmasShape(matrix [][]string, i int, j int) bool {
	
	if i - 1 < 0 || i + 1 >= len(matrix) {
		return false
	}

	if j - 1 < 0 || j + 1 >= len(matrix[i]) {
		return false
	}

	hasMasInFirstDiagonal := hasOcurrenceInDirection(matrix, i - 1, j + 1, DOWN_LEFT_DIAGONAL)	

	hasMasInSecondDiagonal := hasOcurrenceInDirection(matrix, i - 1, j - 1, DOWN_RIGHT_DIAGONAL)	

	return hasMasInFirstDiagonal && hasMasInSecondDiagonal 
}

func hasOcurrenceInDirection(matrix [][]string, i int, j int, direction [2]int) bool{
	const xmasSize = 3	

	word := make([]string, 0)

	for k := 0 ; k < xmasSize ; k++ {
		letter := matrix[i + k * direction[0]][j + k * direction[1]]
		word = append(word, letter)
	} 

	joinedWord := strings.Join(word, "")

	return joinedWord == "MAS" || joinedWord == "SAM"
}

func parseNumber(number string) int {
	
	n, err := strconv.Atoi(number)

	if err != nil {
		log.Fatalf("An error has ocurred while parsing the number: %v", err)
	}

	return n
}