package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value string
	left, right *Node
}

func main() {
	input := readInput(os.Args[1])
	steps := GetSumExtrapolatedValues(input)
	fmt.Println("The answer is:", steps)
}

func readInput(fileName string) []string {
	file, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("An error has ocurred while reading input:", err)
		os.Exit(1)
	}

	return strings.Split(string(file), "\n")
}

func GetSumExtrapolatedValues(report []string) int {
	totalExtrapolated := 0
	for _, reportSerie := range report {
		extrapolatedValue := getExtrapolatedValue(reportSerie)	
		fmt.Println(reportSerie, "->", extrapolatedValue)
		totalExtrapolated += extrapolatedValue
	}
	return totalExtrapolated
}

func getExtrapolatedValue(reportSerie string) int {
	serieAsText := strings.Split(reportSerie, " ")
	serie := make([]int, len(serieAsText))

	for i := 0 ; i < len(serieAsText) ; i++ {
		serie[i] = parseNumber(serieAsText[i])
	}

	firstElementSequence := []int {}

	isAllZeroes := false

	nextSerie := serie
	for !isAllZeroes {
		nextSerieAux := make([]int, 0)
		firstElementSequence = append(firstElementSequence, nextSerie[0])
		for i := 0 ; i < len(nextSerie) - 1; i++ {
			difference := nextSerie[i+1] - nextSerie[i]
			nextSerieAux = append(nextSerieAux, difference)

			if (difference == 0) {
				isAllZeroes = true
			} else {
				isAllZeroes = false
			}
		}
		nextSerie = nextSerieAux
	}


	extrapolatedValue := firstElementSequence[len(firstElementSequence) - 1]
	for i := len(firstElementSequence) - 2; i >= 0; i-- {
		extrapolatedValue = firstElementSequence[i] - extrapolatedValue
	}

	return extrapolatedValue
}

func parseNumber(n string) int {

	number, err := strconv.Atoi(n)

	if err != nil {
		fmt.Println("An error has ocurred while parsing number:", err)
		os.Exit(1)
	}

	return number
}

