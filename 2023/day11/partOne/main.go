package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	input := readInput(os.Args[1])
	steps := GetSumMinStepsBetweenGalaxies(input)
	fmt.Println("The answer is:", steps)
}

func readInput(fileName string) string {
	file, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("An error has ocurred while reading input:", err)
		os.Exit(1)
	}

	return string(file)
}

func GetSumMinStepsBetweenGalaxies(space string) int {

	galaxyPositions := extractGalaxiesPositions(space)

	var sum int

	for i := 0; i < len(galaxyPositions) ; i++ {
		for j := i ; j < len(galaxyPositions) ; j++ {
			sum += getDistanceBetweenGalaxies(galaxyPositions[i], galaxyPositions[j])		
		}
	}

	return sum
}

func getDistanceBetweenGalaxies(x [2]int, y [2]int) int {
	return int(math.Abs(float64(x[0] - y[0]))) + int(math.Abs(float64(x[1] - y[1])))
}

func extractGalaxiesPositions(space string) [][2]int {

	lineSplittedSpace := strings.Split(space, "\n")
	cellSplittedSpace := make([][]string, len(lineSplittedSpace))

	columnGalaxyCount := make([]int, len(lineSplittedSpace[0])) 

	for row, spaceLine := range lineSplittedSpace {
		cells := strings.Split(spaceLine, "")
		cellSplittedSpace[row] = make([]string, len(cells))
		for col, cell := range cells {
			if cell == "#" {
				columnGalaxyCount[col]++
			}
			cellSplittedSpace[row][col] = cell
		}
	}

	galaxyPositions := make([][2]int, 0)

	offsetRowGalaxy := 0
	
	for row, spaceLine := range cellSplittedSpace {
		rowHasGalaxies := false

		expandedSpaceRow := make([]string, 0)

		offsetColGalaxy := 0

		for col, cell := range spaceLine {
			isGalaxy := cell == "#"

			if isGalaxy {
				rowHasGalaxies = true
				galaxyPositions = append(galaxyPositions, [2]int {row + offsetRowGalaxy, col + offsetColGalaxy})
			}
			
			expandedSpaceRow = append(expandedSpaceRow, cell)
			
			if columnGalaxyCount[col] == 0 {
				expandedSpaceRow = append(expandedSpaceRow, ".")
				offsetColGalaxy++
			}
		}


		if !rowHasGalaxies {
			offsetRowGalaxy++
		}

		fmt.Println(galaxyPositions)
	}

	return galaxyPositions
}
