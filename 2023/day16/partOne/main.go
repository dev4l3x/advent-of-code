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
	steps := CountEnergized(input)
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

func CountEnergized(contraption string) int {

	lineContraption := strings.Split(contraption, "\n")
	
	parsedContraption := make([][]string, len(lineContraption))
	beamMap := make([][]int, len(lineContraption))

	for i, line := range lineContraption {
		parsedContraption[i] = split(line)
		beamMap[i] = make([]int, len(line))
	}

	castBeam(parsedContraption, &beamMap)

	return countEnergized(beamMap)
}

func split(s string) []string {
	var newS []string
	for _, s := range s {
		newS = append(newS, string(s))
	}
	return newS
}

const (
	UP int = 1
	DOWN = 2
	LEFT = 3
	RIGHT = 4
)

var splitters = map[int]map[string]bool {
	UP: {
		"-": true,
	},
	DOWN: {
		"-": true,
	},
	LEFT: {
		"|": true,
	},
	RIGHT: {
		"|": true,
	},
}

var directions = map[int]map[string]int{
	UP: {
		"\\": LEFT,
		"/": RIGHT,
	},
	DOWN: {
		"\\": RIGHT,
		"/": LEFT,
	},
	LEFT: {
		"\\": UP,
		"/": DOWN,
	},
	RIGHT: {
		"\\": DOWN,
		"/": UP,
	},
}

type beam struct {
	row, col, direction int
}

func castBeam(contraption [][]string, beamMap *[][]int) {

	beams := []beam { {0, 0, RIGHT} }

	visistedTiles := make([][]map[int]bool, len(contraption))

	for len(beams) != 0 {

		fmt.Println(beams)
		printMatrix(contraption)
		printMatrix(*beamMap)


		currentBeam := beams[0]
		beams = beams[1:]

		row, col, direction := currentBeam.row, currentBeam.col, currentBeam.direction

		if row >= len(*beamMap) || row < 0 || col >= len((*beamMap)[0]) || col < 0 {
			continue
		}

		if visistedTiles[row] == nil {
			visistedTiles[row] = make([]map[int]bool, len(contraption[row]))
		}

		if visistedTiles[row][col] == nil {
			visistedTiles[row][col] = make(map[int]bool)
		}

		_, isVisited := visistedTiles[row][col][direction]

		if row >= len(*beamMap) || row < 0 || col >= len((*beamMap)[0]) || col < 0 || isVisited {
			fmt.Println("Already visited!. Skipping..")
			continue
		}


		tile := contraption[row][col]
		visistedTiles[row][col][direction] = true

		(*beamMap)[row][col]++

		newDirection, isChangeDirection := directions[direction][tile]

		if isChangeDirection {
			var nextTile [2]int
			switch newDirection {
				case UP: nextTile = [2]int{row-1, col}
				case DOWN: nextTile = [2]int{row+1, col}
				case LEFT: nextTile = [2]int{row, col-1}
				case RIGHT: nextTile = [2]int{row, col+1}
			}
			beams = append(beams, beam{nextTile[0], nextTile[1], newDirection})
			continue
		}
		
		_, isSplit := splitters[direction][tile]

		if isSplit {
			if direction == UP || direction == DOWN {
				beams = append(beams, beam{row, col - 1, LEFT})
				beams = append(beams, beam{row, col + 1, RIGHT})
			} else {
				beams = append(beams, beam{row-1, col, UP})
				beams = append(beams, beam{row+1, col, DOWN})
			}
			continue
		}

		var nextTile [2]int
		switch direction{
			case UP: nextTile = [2]int{row-1, col}
			case DOWN: nextTile = [2]int{row+1, col}
			case LEFT: nextTile = [2]int{row, col-1}
			case RIGHT: nextTile = [2]int{row, col+1}
		}

		beams = append(beams, beam{nextTile[0], nextTile[1], direction})

	}
}

func countEnergized(beamMap [][]int) int {
	var sum int
	for _, row := range beamMap {
		for _, cell := range row {
			if cell > 0 {
				sum++		
			}	
		}
	}
	return sum
}

func printMatrix[T any](m [][]T) {
	for _, row := range m {
		fmt.Println(row)
	}
}
