package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := readInput(os.Args[1])
	steps := GetAreaFromMap(input)
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

var pipes = map[string]bool{
	"|": true,
	"-": true, 
	"L": true,
	"J": true,
	"7": true,
	"F": true,
}

var changeDirectionMappings = map[int]map[string]int {
	UP: {
		"7": LEFT,
		"F": RIGHT,
	},
	LEFT: {
		"L": UP,
		"F": DOWN,
	},
	DOWN: {
		"L": RIGHT,
		"J": LEFT,
	},
	RIGHT: {
		"7": DOWN,
		"J": UP,
	},
}

const (
	UP = 0
	DOWN = 1
	LEFT = 2
	RIGHT = 3
)

func GetAreaFromMap(pipeMap string) int {

	startingIndexes, parsedPipeMap := getMatrixFromMap(pipeMap)

	var currentTile[2]int

	if startingIndexes[0] - 1 >= 0 && isPipe(parsedPipeMap[startingIndexes[0] - 1][startingIndexes[1]], UP) {
		currentTile[0] = startingIndexes[0] - 1	
		currentTile[1] = startingIndexes[1]
	} else if startingIndexes[0] + 1 < len(parsedPipeMap) && isPipe(parsedPipeMap[startingIndexes[0] + 1][startingIndexes[1]], DOWN) {
		currentTile[0] = startingIndexes[0] + 1
		currentTile[1] = startingIndexes[1]
	} else if startingIndexes[1] - 1 < len(parsedPipeMap[startingIndexes[0]]) && isPipe(parsedPipeMap[startingIndexes[0]][startingIndexes[1] - 1], LEFT) {
		currentTile[0] = startingIndexes[0]
		currentTile[1] = startingIndexes[1] - 1
	} else {
		currentTile[0] = startingIndexes[0]
		currentTile[1] = startingIndexes[1] + 1
	}

	

	hasReachedEndOfLoop := false
	totalLoopPipes := 1
	previousTile := startingIndexes

	direction := getDirection(previousTile, currentTile)

	pathPoints := map[int]map[int]bool {
		startingIndexes[0]: {
			startingIndexes[1]: true,
		},
	}

	for !hasReachedEndOfLoop {

		currentTileValue :=  parsedPipeMap[currentTile[0]][currentTile[1]]

		if currentTileValue == "S" {
			hasReachedEndOfLoop = true
			break
		}

		_, existsMap := pathPoints[currentTile[0]]
		if !existsMap {
			pathPoints[currentTile[0]] = make(map[int]bool) 
		}

		pathPoints[currentTile[0]][currentTile[1]] = true

		newDirection, ok := changeDirectionMappings[direction][currentTileValue]

		if ok {
			direction = newDirection
		}

		//parsedPipeMap[currentTile[0]][currentTile[1]] = "*"

		switch direction {
			case UP: 
				currentTile[0]--
			case DOWN:
				currentTile[0]++
			case LEFT:
				currentTile[1]--
			case RIGHT:
				currentTile[1]++
		}
		totalLoopPipes++
	}

	//parsedPipeMap[startingIndexes[0]][startingIndexes[1]] = "*"

	fmt.Println(parsedPipeMap)

	return getTilesInsideLoop(parsedPipeMap, pathPoints)
}

func getTilesInsideLoop(starMap [][]string, pathPoints map[int]map[int]bool) int {
	
	var tilesOutside int

	for i, row := range starMap {

		for j := range row {

			_, isPathTile := pathPoints[i][j]

			if isPathTile {
				continue
			}

			isTileInside := isTileInsideLoop([2]int {i, j}, pathPoints, starMap)	
			if isTileInside {
				tilesOutside++
			}
		}

		fmt.Println(row, "->", tilesOutside)
	}
	return tilesOutside
}

func isTileInsideLoop(tileIndexes [2]int, pathPoints map[int]map[int]bool, starMap [][]string) bool {

	pathTilesAfterTile := 0

	for i, j := tileIndexes[0], tileIndexes[1]; i < len(starMap) && j < len(starMap[i]); i, j = i + 1, j + 1 {
		_, ok := pathPoints[i][j]

		pathTile := starMap[i][j]

		if ok && (pathTile == "L" || pathTile == "7") {
			pathTilesAfterTile+=2
		} else if ok {
			pathTilesAfterTile++
		}
	}

	fmt.Println(starMap[tileIndexes[0]][tileIndexes[1]], pathTilesAfterTile)

	return pathTilesAfterTile % 2 != 0
}

func getDirection(previousTile [2]int, currentTile [2]int) int {
	row, col := currentTile[0] - previousTile[0], currentTile[1] - previousTile[1]

	if row == 1 && col == 0 {
		return DOWN
	} else if row == -1 && col == 0 {
		return UP
	} else if row == 0 && col == 1 {
		return RIGHT
	} else {
		return LEFT
	}
}

func isPipe(tile string, direction int) bool {
	_, existsPipe := pipes[tile]

	if !existsPipe {
		return false
	}

	if direction == UP {
		return tile == "|" || tile == "7" || tile == "F"
	} else if direction == DOWN {
		return tile == "|" || tile == "L" || tile == "J"
	} else if direction == LEFT {
		return tile == "-" || tile == "L" || tile == "F"
	} else {
		return tile == "-" || tile == "7" || tile == "J"
	}
}

func getMatrixFromMap(pipeMap string) (startingIndexes [2]int, processedMap[][]string) {
	filesSplitted := strings.Split(pipeMap, "\n")
	processedMap = make([][]string, 0)

	foundIndexes := false

	for i := 0 ; i < len(filesSplitted) ; i++ {
		processedMap = append(processedMap, strings.Split(filesSplitted[i], ""))
		for j := 0; !foundIndexes && j < len(processedMap[i]) ; j++ {
			if processedMap[i][j] == "S" {
				startingIndexes = [2]int {i, j}
				foundIndexes = true
			}
		}
	}
	return startingIndexes, processedMap
}
