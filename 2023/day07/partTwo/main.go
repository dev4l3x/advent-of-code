package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readInput(os.Args[1])
	totalWinnings := GetTotalWinningsFromHands(input)
	fmt.Println("The answer is:", totalWinnings)
}

func readInput(fileName string) []string {
	file, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("An error has ocurred while reading input:", err)
		os.Exit(1)
	}

	return strings.Split(string(file), "\n")
}

func GetTotalWinningsFromHands(handsWithBids []string) int {
	ranks := make([]string, len(handsWithBids))
	handBid := make(map[string]int)
	for i, handsWithBid := range handsWithBids {
		hand, bid := splitHandAndBid(handsWithBid)
		handBid[hand] = bid
		ranks[i] = hand
	}

	orderedRanks := orderHandsFromLowestToHighest(ranks)

	totalWinnings := 0

	for index, hand := range orderedRanks {
		rank := index + 1	
		totalWinnings += rank * handBid[hand]
	}

	return totalWinnings	
}

func orderHandsFromLowestToHighest(hands []string) []string {

	if len(hands) == 1 {
		return hands
	}

	halfIndex := len(hands) / 2

	leftOrdered := orderHandsFromLowestToHighest(hands[0:halfIndex])
	rightOrdered := orderHandsFromLowestToHighest(hands[halfIndex:])

	orderedHand := mergeOrderedSlices(leftOrdered, rightOrdered)

	return orderedHand
}

func mergeOrderedSlices(left []string, right []string) []string {
	ordered := make([]string, len(left) + len(right))
	var i, j, k int

	for i < len(left) || j < len(right) {
		if j == len(right) || (i < len(left) && isLower(left[i], right[j])) {
			ordered[k] = left[i]
			i++
			k++
		} else if i == len(left) || (j < len(right) && isLower(right[j], left[i])) {
			ordered[k] = right[j]
			k++
			j++
		}
	}
	return ordered
}

var handType = make(map[string]int)

func isLower(leftHand string, rightHand string) bool {

	leftHandSplitted := strings.Split(leftHand, "")
	rightHandSplitted := strings.Split(rightHand, "")

	leftHandType, leftOk := handType[leftHand]
	rightHandType, rightOk := handType[rightHand]
	
	if !leftOk {
		leftHandType = getHandType(leftHandSplitted)
		handType[leftHand] = leftHandType
	}

	if !rightOk {
		rightHandType := getHandType(rightHandSplitted)
		handType[rightHand] = rightHandType
	}


	leftHandType = getHandType(leftHandSplitted)
	rightHandType = getHandType(rightHandSplitted)

	if leftHandType < rightHandType{
		return true
	} else if rightHandType < leftHandType {
		return false
	}


	for i := 0 ; i < len(leftHandSplitted) && i < len(rightHandSplitted) ; i++ {
		leftCardStrength := getCardStrength(leftHandSplitted[i])	
		rightCardStrength := getCardStrength(rightHandSplitted[i])	
		if leftCardStrength < rightCardStrength {
			return true
		} else if rightCardStrength < leftCardStrength{
			return false
		}
	}

	return false
}

var cardStrength = map[string]int {
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 1,
	"T": 10,
}

func getCardStrength(card string) int {
	strength, ok := cardStrength[card]

	if !ok {
		return parseNumber(card)
	}

	return strength
}

func getHandType(hand []string) int {
	cardsCount := make(map[string]int)
	for _, card := range hand { 
		_, ok := cardsCount[card]		

		if !ok {
			cardsCount[card] = 0
		}

		cardsCount[card]++
	}

	equalCards := map[int] int{
		5: 0,
		4: 0,
		3: 0,
		2: 0,
		1: 0,
	}

	for card, count := range cardsCount {
		if card == "J" {
			continue
		}
		equalCards[count]++
	}

	jokersCount := cardsCount["J"] 
	leftCards := 5 - jokersCount

	if leftCards == 0 || equalCards[leftCards] == 1 {
		fmt.Println("Hand", hand, ": Five of a kind")
		return 7
	} else if equalCards[leftCards - 1] == 1 {
		fmt.Println("Hand", hand, ": Four of a kind")
		return 6
	} else if (equalCards[3] == 1 && equalCards[2] == 1) || (equalCards[2] == 2 && jokersCount == 1) {
		fmt.Println("Hand", hand, ": Full house")
		return 5
	} else if equalCards[leftCards - 2] >= 1 {
		fmt.Println("Hand", hand, ": Three of a Kind")
		return 4
	} else if equalCards[2] >= 2 {
		fmt.Println("Hand", hand, ": Two Pairs")
		return 3
	} else if equalCards[leftCards - 3] >= 1 {
		fmt.Println("Hand", hand, ": One Pair")
		return 2
	} else {
		fmt.Println("Hand", hand, ": High Card")
		return 1
	}
}

func splitHandAndBid(handsWithBids string) (string, int) {
	parsedHand := strings.Split(handsWithBids, " ")
	return parsedHand[0], parseNumber(parsedHand[1])
}

func parseNumber(number string) int {
	n, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println("An error has ocurred while parsing number:", err)
		os.Exit(1)
	}
	return n
}
