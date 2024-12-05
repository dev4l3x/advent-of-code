package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
			if matrix[i][j] == "X" {
				ocurrences += countXmasOcurrences(matrix, i, j)
			}
		}
	}

	fmt.Printf("The sum of all ocurrences is: %v\n", ocurrences)

}

func countXmasOcurrences(matrix [][]string, i int, j int) int {
	
	const xmasSize = 3 // XMAS
	xmas := "XMAS" 
	ocurrences := 0
	
	// Search for the word XMAS in the row with right direction
	if j + xmasSize < len(matrix[i]) {
		row := matrix[i]
		word := row[j:j+xmasSize+1]
		if strings.Join(word, "") == xmas {
			ocurrences++
		}
	}

	// Search for the word XMAS in the row with left direction
	if j - xmasSize >= 0 {
		row := matrix[i]
		word := row[j-xmasSize:j+1]
		if strings.Join(word, "") == "SAMX" {
			ocurrences++
		}
	}

	// Search for the word XMAS in the column with down direction
	if i + xmasSize < len(matrix) {
		column := make([]string, 0)
		for k := i; k <= i + xmasSize; k++ {
			column = append(column, matrix[k][j])
		}
		if strings.Join(column, "") == xmas {
			ocurrences++
		}
	}

	// Search for the word XMAS in the column with up direction
	if i - xmasSize >= 0 {
		column := make([]string, 0)
		for k := 0; k <= xmasSize; k++ {
			column = append(column, matrix[i-k][j])
		}
		if strings.Join(column, "") == xmas {
			ocurrences++
		}
	}

	// Search for the word XMAS in the diagonal with up direction
	if i - xmasSize >= 0 && j - xmasSize >= 0 {
		diagonal := make([]string, 0)
		for k := 0; k <= xmasSize; k++ {
			diagonal = append(diagonal, matrix[i-k][j-k])
		}
		if strings.Join(diagonal, "") == xmas {
			ocurrences++
		}
	}

	// Search for the word XMAS in the diagonal with down direction
	if i + xmasSize < len(matrix) && j + xmasSize < len(matrix[i]) {
		diagonal := make([]string, 0)
		for k := 0; k <= xmasSize; k++ {
			diagonal = append(diagonal, matrix[i+k][j+k])
		}
		if strings.Join(diagonal, "") == xmas {
			ocurrences++
		}
	}

	// Search for the word XMAS in the diagonal with up direction
	if i - xmasSize >= 0 && j + xmasSize < len(matrix[i]) {
		diagonal := make([]string, 0)
		for k := 0; k <= xmasSize; k++ {
			diagonal = append(diagonal, matrix[i-k][j+k])
		}
		if strings.Join(diagonal, "") == xmas {
			ocurrences++
		}
	}

	// Search for the word XMAS in the diagonal with down direction
	if i + xmasSize < len(matrix) && j - xmasSize >= 0 {
		diagonal := make([]string, 0)
		for k := 0; k <= xmasSize; k++ {
			diagonal = append(diagonal, matrix[i+k][j-k])
		}
		if strings.Join(diagonal, "") == xmas {
			ocurrences++
		}
	}

	return ocurrences
}

func parseNumber(number string) int {
	
	n, err := strconv.Atoi(number)

	if err != nil {
		log.Fatalf("An error has ocurred while parsing the number: %v", err)
	}

	return n
}