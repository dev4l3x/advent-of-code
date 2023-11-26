package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	ROCK = 1
	PAPER = 2
	SCISSORS = 3
)

var winRules = map[int]int {
	ROCK: SCISSORS,
	PAPER: ROCK,
	SCISSORS: PAPER,
}

var loseRules = map[int]int {
	ROCK: PAPER,
	PAPER: SCISSORS,
	SCISSORS: ROCK,
}

var mappings = map[string]int {
	"A": ROCK,
	"B": PAPER,
	"C": SCISSORS,
}

func readInput() string {
	file, error := os.ReadFile("input.txt")

	if error != nil {
		fmt.Printf("An error has ocurred while reading the input: %v", error)
	}

	return string(file)
}

func getChoiceFromPrediction(oponentChoice int, prediction string) int {
	switch prediction {
	case "X":
		return winRules[oponentChoice]
	case "Y":
		return oponentChoice	
	case "Z":
		return loseRules[oponentChoice]
	default:
		return -1
	}
}

func getRoundsFromInput(input string) [][2]int{

	compactedRounds := strings.Split(input, "\n")

	var rounds [][2]int
	
	for index := range compactedRounds {

		if len(compactedRounds[index]) == 0 {
			break
		}
		
		roundSelections := strings.Split(compactedRounds[index], " ")
		oponentChoice := mappings[roundSelections[0]]
		rounds = append(rounds, [2]int{ oponentChoice, getChoiceFromPrediction(oponentChoice, roundSelections[1]) })
	}
	return rounds
}

func getMyScoreForRound(round [2]int) (score int) {
	oponentChoice := round[0]
	myChoice := round[1]

	if winRules[myChoice] == oponentChoice {
		// I WIN!
		score += myChoice + 6

	} else if winRules[oponentChoice] == myChoice {
		// I LOSE!
		score += myChoice
	} else {
		// IT'S A TIE
		score += myChoice + 3
	}
	return score
}

func getMyTotalScore() (totalScore int) {
	input := readInput()
	rounds := getRoundsFromInput(input)
	for _, round := range rounds {
		totalScore += getMyScoreForRound(round)	
	}
	return totalScore
}

func main() {
	fmt.Printf("getMyTotalScore(): %v\n", getMyTotalScore())
}
