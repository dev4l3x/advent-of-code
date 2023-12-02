package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
	
	var sumOfMinimumCubesProduct int

	for _, game := range games {
		r, g, b := GetMinimumNumberOfCubes(strings.Split(game, ": ")[1])
		sumOfMinimumCubesProduct += r*g*b 
	}

	fmt.Println("The products sum is:", sumOfMinimumCubesProduct)
}

func GetMinimumNumberOfCubes(game string) (r int, g int, b int) {

	gameSets := strings.Split(game, "; ")

	var maxCubesPresentedInGame = map[string]int {
		"red": 0,
		"green": 0,
		"blue": 0,
	}
	
	for _, gameSet := range gameSets {
		cubes := strings.Split(gameSet, ", ")
		for _, cube := range cubes {
			parsedCube := strings.Split(cube, " ")
			cubesNumber := parseNumber(parsedCube[0])
			color := parsedCube[1]
			if cubesNumber > maxCubesPresentedInGame[color] {
				maxCubesPresentedInGame[color] = cubesNumber
			}
		}
	}

	r, g, b = maxCubesPresentedInGame["red"], maxCubesPresentedInGame["green"], maxCubesPresentedInGame["blue"]

	return r, g, b
}

func parseNumber(number string) int {
	n, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println("An error has occurred while parsing the number", number)
		os.Exit(1)
	}
	return n
}


