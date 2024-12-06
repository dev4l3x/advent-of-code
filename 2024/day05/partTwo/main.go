package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, error := os.ReadFile("input.txt")	

	if error != nil {
		fmt.Printf("An error has ocurred while processing the input: %v", error)
		os.Exit(1)
	}

	input := string(file)

	printings := strings.Split(input, "\n\n")

	rulesConcatenated := strings.Split(printings[0], "\n")
	printingUpdatesConcatenated := strings.Split(printings[1], "\n")

	rules := parseRules(rulesConcatenated)

	totalSum := 0

	for _, printingUpdateConcat := range printingUpdatesConcatenated {
		printingUpdate := strings.Split(printingUpdateConcat, ",")
		fixedPrinting := make([]string, 0)

		for i := 0 ; i < len(printingUpdate) ; i++ {
			checkingPage := printingUpdate[i]

			j := 0
			for (j < len(fixedPrinting)) {
				if _, ok := rules[checkingPage][1][fixedPrinting[j]]; ok {
					break
				}
				j++
			}

			fixedPrinting = insertAtIndex(fixedPrinting, j, checkingPage)	
		}

		isValidUpdate := isSameSlice(fixedPrinting, printingUpdate) 

		if isValidUpdate {
			continue
		}

		fmt.Printf("Original printing: %v", printingUpdate)
		fmt.Printf(" - Fixed printing: %v\n", fixedPrinting)

		middleIndex := int(math.Round(float64(len(fixedPrinting) / 2)))
		middleUpdate := parseNumber(fixedPrinting[middleIndex]) 
		totalSum += middleUpdate	
	}
	

	fmt.Printf("The sum of all ocurrences is: %v\n", totalSum)

}

func isSameSlice(firstSlice []string, secondSlice []string) bool {
	for i := 0; i < len(firstSlice); i++ {
		if firstSlice[i] != secondSlice[i] {
			return false 
		}
	}

	return true
}

func insertAtIndex(array []string, index int, elem string) []string {
	newArray := make([]string, 0)	

	newArray = append(newArray, array[:index]...)
	newArray = append(newArray, elem)
	newArray = append(newArray, array[index:]...)

	return newArray
}

func parseRules(rulesConcatenated []string) map[string][2]map[string]bool {
	// Each map will have a a number as a key, and the value will be all the numbers that should go before in the 
	// index 0 and all the numbers that should be after in the index 1
	rules := make(map[string][2]map[string]bool, 0)

	for _, ruleConcat := range rulesConcatenated {
		rule := strings.Split(ruleConcat, "|")

		// Initialize map if rules does not exist yet
		if _, ok := rules[rule[0]]; !ok {
			rules[rule[0]] = [2]map[string]bool {make(map[string]bool), make(map[string]bool)}
		}

		if _, ok := rules[rule[1]]; !ok {
			rules[rule[1]] = [2]map[string]bool {make(map[string]bool), make(map[string]bool)}
		}

		// Adds to the number of the left side of the rule, the number that should go after
		rules[rule[0]][1][rule[1]] = true

		// Adds to the number of the right side of the rule, the number that should go before 
		rules[rule[1]][0][rule[0]] = true
	}

	return rules
}

func parseNumber(number string) int {
	
	n, err := strconv.Atoi(number)

	if err != nil {
		log.Fatalf("An error has ocurred while parsing the number: %v", err)
	}

	return n
}