package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	cards := readInput(os.Args[1])

	points := GetPointsFromCards(cards)

	fmt.Println("The total number of points is:", points)
}

func readInput(fileName string) []string {
	file, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("An error has occurred:", err)
		os.Exit(1)
	}

	return strings.Split(string(file), "\n")
}

func GetPointsFromCards(cards []string) int {

	var totalPoints int
	
	for _, card := range cards {
		game := strings.Split(strings.Split(card, ": ")[1], " | ")
		textWinningNumbers := strings.Split(game[0], " ")
		textNumbers := strings.Split(game[1], " ")

		var winningNumbers map[int] bool = make(map[int]bool)

		for _, textWinningNumber := range textWinningNumbers {

			winningNumber, isNumber := parseNumber(textWinningNumber)

			if !isNumber {
				continue
			}

			winningNumbers[winningNumber] = true

		}

		var gamePoints int
		for _, textNumber := range textNumbers {
			number, isNumber := parseNumber(textNumber)	
			
			if !isNumber {
				continue
			}

			isWinningNumber := winningNumbers[number]

			if isWinningNumber && gamePoints > 0 {
				gamePoints *= 2	
			} else if isWinningNumber && gamePoints == 0 {
				gamePoints++
			}
		}

		totalPoints += gamePoints
	}

	return totalPoints
}

func parseNumber(number string) (int, bool) {
	n, err := strconv.Atoi(number)

	if err != nil {
		return 0, false
	}

	return n, true
}
