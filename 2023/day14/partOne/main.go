package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	defer timer()()
	input := readInput(os.Args[1])
	steps := GetTotalLoad(input)
	fmt.Println("The answer is:", steps)
}

func timer() func() {
    start := time.Now()
    return func() {
	fmt.Printf("Execution time: %v\n", time.Since(start))
    }
}

func readInput(fileName string) string {
	file, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("An error has ocurred while reading input:", err)
		os.Exit(1)
	}

	return string(file)
}

func GetTotalLoad(platform string) int {
	
	plat := strings.Split(platform, "\n")

	nextPositions := make([]int, len(plat[0]))

	var totalLoad int

	for i, line := range plat {
		parsedLine := strings.Split(line, "")	

		for j, cell := range parsedLine {
			if cell == "O" {
				if nextPositions[j] == i {
					nextPositions[j] = i + 1		
					totalLoad += len(parsedLine) - i
				} else {
					totalLoad += len(parsedLine) - nextPositions[j]
					nextPositions[j]++
				}
			} else if cell == "#" {
				nextPositions[j] = i + 1		
			}
		}
	}

	return totalLoad
}
