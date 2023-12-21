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

const (
	UP = "UP"
	DOWN = "DOWN"
	LEFT = "LEFT"
	RIGHT = "RIGHT"
)

var directions = map[string][2]int {
	UP: {1, 0},
	DOWN: {-1, 0},
	LEFT: {0, -1},
	RIGHT: {0, 1},
}

var cycleCache = make(map[string]cachedCycle)

type cachedCycle struct {
	platform [][]string
	load int
	cycleNumber int
}

func GetTotalLoad(platform string) int {
	
	plat := strings.Split(platform, "\n")

	parsedPlatform := make([][]string, len(plat))

	for i, line := range plat {
		parsedLine := strings.Split(line, "")
		parsedPlatform[i] = parsedLine
	}

	totalLoad := 0
	cycles := 1000000000
	foundCycle := false
	currentPlatform := parsedPlatform

	for i := 0; i < cycles ; i++ {

		key := encodeCacheKey(currentPlatform)
		cached, ok := cycleCache[key]

		if ok && !foundCycle{

			foundCycle= true

			totalLoad = cached.load
			currentPlatform = duplicate(cached.platform)

			remainingCycles := cycles - i
			cycleLength := i - cached.cycleNumber

			i = cycles - (remainingCycles % cycleLength)

			fmt.Println("Jumping to cycle:", i)

			continue	
		}


		upPlatform, _ := tiltPlatformVertical(currentPlatform, UP)

		leftPlatform, _ := tiltPlatformLeft(upPlatform, UP)

		downPlatform, _ := tiltPlatformDown(leftPlatform, UP)

		currentPlatform, totalLoad = tiltPlatformRight(downPlatform, UP)

		cycleCache[key] = cachedCycle{duplicate(currentPlatform), totalLoad, i}	

	}
	

	return totalLoad
}

func duplicate(matrix [][]string) [][]string {
	
	duplicatedMatrix := make([][]string, len(matrix))

	for i, line := range matrix {
		duplicatedLine := make([]string, len(line))
		copy(duplicatedLine, line)
		duplicatedMatrix[i] = duplicatedLine
	}
	return duplicatedMatrix
}

func encodeCacheKey(platform [][]string) string {
	var key string
	for _, line := range platform {
		encodedLine := strings.Join(line, "")	
		key += encodedLine
	}
	return key
}

func tiltPlatformVertical(platform [][]string, direction string) ([][]string, int) {

	newPlatform := duplicate(platform) 

	nextPositions := make([]int, len(platform[0]))
	var totalLoad int

	for i, line := range platform {
		for j, cell := range line {
			if cell == "O" {
				if nextPositions[j] == i {
					nextPositions[j] = i + 1		
					totalLoad += len(platform) - i
				} else {
					newPlatform[nextPositions[j]][j] = cell
					newPlatform[i][j] = "." 
					totalLoad += len(platform) - nextPositions[j]
					nextPositions[j]++
				}
			} else if cell == "#" {
				nextPositions[j] = i + 1		
			}
		}
	}

	return newPlatform, totalLoad
}

func tiltPlatformDown(platform [][]string, direction string) ([][]string, int) {
	newPlatform := duplicate(platform)

	nextPositions := make([]int, len(platform[0]))

	for i := 0; i < len(nextPositions) ; i++ {
		nextPositions[i] = len(platform) - 1	
	}

	var totalLoad int

	for i := len(platform) - 1 ; i >= 0 ; i-- {
		for j, cell := range platform[i] {

			nextPosition := nextPositions[j]

			if cell == "O" {
				if nextPosition == i {
					nextPositions[j] = i - 1		
					totalLoad += len(platform) - i
				} else {
					newPlatform[nextPosition][j] = cell
					newPlatform[i][j] = "." 
					totalLoad += len(platform) - nextPositions[j]
					nextPositions[j]--
				}
			} else if cell == "#" {
				nextPositions[j] = i - 1		
			}
		}
	}

	return newPlatform, totalLoad
}

func tiltPlatformRight(platform [][]string, direction string) ([][]string, int) {
	newPlatform := duplicate(platform) 

	nextPositions := make([]int, len(platform))

	for i := 0 ; i < len(nextPositions) ; i++{
		nextPositions[i] = len(platform[i]) - 1	
	}

	var totalLoad int

	for i, line := range platform {
		for j := len(line) - 1 ; j >= 0; j-- {
			cell := platform[i][j]
			if cell == "O" {
				if nextPositions[i] == j {
					nextPositions[i] = j - 1		
					totalLoad += len(platform) - i
				} else {
					newPlatform[i][nextPositions[i]] = cell
					newPlatform[i][j] = "." 
					totalLoad += len(platform) - i
					nextPositions[i]--
				}
			} else if cell == "#" {
				nextPositions[i] = j - 1		
			}
		}
	}

	return newPlatform, totalLoad
}

func tiltPlatformLeft(platform [][]string, direction string) ([][]string, int) {
	newPlatform := duplicate(platform)

	nextPositions := make([]int, len(platform))

	var totalLoad int

	for i, line := range platform {
		for j, cell := range line {
			if cell == "O" {
				if nextPositions[i] == j {
					nextPositions[i] = j + 1		
					totalLoad += len(platform) - i
				} else {
					newPlatform[i][nextPositions[i]] = cell
					newPlatform[i][j] = "." 
					totalLoad += len(platform) - i
					nextPositions[i]++
				}
			} else if cell == "#" {
				nextPositions[i] = j + 1		
			}
		}
	}

	return newPlatform, totalLoad
}

func printMatrix(matrix [][]string) {
	for _, row := range matrix {
		fmt.Println(row)
	}
	fmt.Println()
}
