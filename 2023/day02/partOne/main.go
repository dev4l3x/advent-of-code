package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var maxCubesNumber = map[string]int {
	"red": 12,
	"green": 13,
	"blue": 14,
}

func readGames(fileName string) []string {
	file, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("An error has occurred while reading the input:", err)
		os.Exit(1)
	}

	return strings.Split(string(file), "\n")
}

func main() {

	games := readGames(os.Args[1])
	
	var gameIdsSum int

	for _, game := range games {
		gameId := getGameId(game)
		isPossible := IsGamePossible(strings.Split(game, ": ")[1])
		if isPossible {
			gameIdsSum += gameId
		}
	}

	fmt.Println("The ids sum is:", gameIdsSum)
}

func getGameId(game string) int {
	gameNumber := strings.Split(game, " ")[1]
	return parseNumber(gameNumber[0:len(gameNumber) - 1])
}

func IsGamePossible(game string) (isPossible bool) {

	gameSets := strings.Split(game, "; ")
	
	for _, gameSet := range gameSets {
		cubes := strings.Split(gameSet, ", ")
		for _, cube := range cubes {
			parsedCube := strings.Split(cube, " ")
			cubesNumber := parseNumber(parsedCube[0])
			color := parsedCube[1]
			if cubesNumber > maxCubesNumber[color] {
				return false
			}
		}
	}


	return true
}

func parseNumber(number string) int {
	n, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println("An error has occurred while parsing the number", number)
		os.Exit(1)
	}
	return n
}


