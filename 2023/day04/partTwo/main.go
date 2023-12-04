package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	cards := readInput(os.Args[1])

	points := GetSumOfScratchCards(cards)

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

func GetSumOfScratchCards(cards []string) int {

	var cardMatches map[int]int = make(map[int]int)

	var cardsToProccess []int
	
	for _, card := range cards {
		parsedCard := strings.Split(card, ": ")
		cardNumber := getCardNumber(parsedCard[0])
		game := strings.Split(parsedCard[1], " | ")
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

		var matches int
		for _, textNumber := range textNumbers {
			number, isNumber := parseNumber(textNumber)	
			
			if !isNumber {
				continue
			}

			isWinningNumber := winningNumbers[number]

			if isWinningNumber {
				matches++
			}
		}

		cardMatches[cardNumber] = matches
		cardsToProccess = append(cardsToProccess, cardNumber)
	}

	totalScratchCards := len(cardsToProccess)
	for len(cardsToProccess) > 0 {
		card := cardsToProccess[0]
		cardsToProccess = cardsToProccess[1:]
		
		matches := cardMatches[card]

		totalScratchCards += matches

		for i := 1 ; i <= matches ; i++ {
			cardsToProccess = append(cardsToProccess, card + i)
		}
	}

	return totalScratchCards 
}

func getCardNumber(card string) int {
	parsedCardTitle := strings.Split(card, " ")
	number := parsedCardTitle[len(parsedCardTitle) - 1]
	n, ok := parseNumber(number)
	if !ok {
		fmt.Println("An error has ocurred while parsing card number:", card)
		os.Exit(1)
	}
	return n
}

func parseNumber(number string) (int, bool) {
	n, err := strconv.Atoi(number)

	if err != nil {
		return 0, false
	}

	return n, true
}
