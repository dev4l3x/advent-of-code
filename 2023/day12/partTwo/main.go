package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	defer timer()()
	input := readInput(os.Args[1])
	steps := GetSumPossibleArrangements(input)
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

func GetSumPossibleArrangements(space string) int {

	lines := strings.Split(space, "\n")

	var totalSum int

	for _, line := range lines {


		springs, numbersAsText := extractSpringsAndNumbers(line)

		unfoldedSprings, unfoldedTextNumbers := unfoldLine(springs, numbersAsText)

		unfoldedNumbers := parseNumbers(unfoldedTextNumbers)

		//r := regexp.MustCompile("\\.+")

		lineSum := getPossibleCombinations(unfoldedSprings, unfoldedNumbers, 0)

		fmt.Println("Line:", line, "->",lineSum)
		totalSum += lineSum
	}


	return totalSum
}

var cache = make(map[string]int)

func getPossibleCombinations(springs string, groups []int, currentGroup int) int {

	fmt.Println(springs, groups, currentGroup)

	key := getKey(springs, groups, currentGroup)

	if springs == "" {
		if (len(groups) == 1 && currentGroup == groups[0]) || len(groups) == 0 {
			return 1
		} else {
			return 0
		}
	}

	currentChar := springs[0:1]

	isGroupInProgress := currentGroup > 0

	if len(groups) == 0 {
		if strings.Contains(springs, "#") {
			return 0
		} else {
			return 1
		}
	} else if currentGroup > groups[0] {
		return 0
	} 

	solution, isCached := cache[key]

	if isCached {
		fmt.Println("Using cache:", springs, groups, currentGroup, " --> ", solution)
		return solution
	}


	// If we are processing a group (> 0) and find a '.' we must pass to the next group 
	if currentChar == "." {
		var c int
		if isGroupInProgress && currentGroup == groups[0] {
			c = getPossibleCombinations(springs[1:], groups[1:], 0)
		} else if isGroupInProgress {
			return 0
		} else {
			c = getPossibleCombinations(springs[1:], groups, currentGroup)
		}
		cache[key] = c 
		return c
	} else if currentChar == "#" {
		c := getPossibleCombinations(springs[1:], groups, currentGroup + 1)
		cache[key] = c
		return c
	} else {
		if currentChar != "?" {

			//fmt.Println("Found ?")
			os.Exit(1)
		}
		if isGroupInProgress && currentGroup == groups[0] {
			// It has to be a .
			c := getPossibleCombinations(springs[1:], groups[1:], 0)
			cache[key] = c
			return c
		} else if isGroupInProgress && currentGroup < groups[0] {
			// It has to be a #
			c := getPossibleCombinations(springs[1:], groups, currentGroup + 1)
			cache[key] = c
			return c
		} else {
			// It could be both
			c := getPossibleCombinations(springs[1:], groups, currentGroup + 1) + getPossibleCombinations(springs[1:], groups, currentGroup)
			cache[key] = c
			return c
		}
	}

}

func getKey(springs string, groups []int, currentGroup int) string {
	key := springs + "-"

	for _, group := range groups {
		key += fmt.Sprintf("%v,", group)
	}

	return key + "-" + fmt.Sprintf("%v", currentGroup)
}

func unfoldLine(springs string, numbers string) (string, string) {
	
	unfoldedSprings := springs
	unfoldedNumbers := numbers

	for i := 0 ; i < 4 ; i++ {
		unfoldedSprings += fmt.Sprintf("?%s", springs)
		unfoldedNumbers += fmt.Sprintf(",%s", numbers)
	}

	return unfoldedSprings, unfoldedNumbers
}

var cachedCombinations = make(map[int][]string)

func getPossibleGroupsForUnknown(unknownNumber int) []string {
	
	if unknownNumber == 1 {
		return []string {".", "#"}
	}

	cachedCombs, isCached := cachedCombinations[unknownNumber]

	if isCached {
		return cachedCombs 
	}

	possibleComibations := getPossibleGroupsForUnknown(unknownNumber - 1)

	combinations := make([]string, 0)

	for _, combination := range possibleComibations {
		combinations = append(combinations, combination + "#")	
		combinations = append(combinations, combination + ".")	
	}

	cachedCombinations[unknownNumber] = combinations

	return combinations
}

func createRegexForNumbers(numbers []int) regexp.Regexp {

	numbersRegexText := "^\\.*"

	for index, number := range numbers {
		if index == len(numbers) - 1{
			numbersRegexText += fmt.Sprintf("#{%d}\\.*$", number)
			continue
		}
		numbersRegexText += fmt.Sprintf("#{%d}\\.+", number)
	}

	return *regexp.MustCompile(numbersRegexText)
}

func extractSpringsAndNumbers(line string) (string, string) {
	lineSplitted := strings.Split(line, " ")

	return lineSplitted[0], lineSplitted[1]
}

func parseNumbers(numbers string) []int {
	numbersAsText := strings.Split(numbers, ",")

	parsedNumbers := make([]int, 0)

	for _, number := range numbersAsText {
		parsedNumbers = append(parsedNumbers, parseNumber(number))
	}

	return parsedNumbers 
}

func parseNumber(n string) int {
	number, err := strconv.Atoi(n)

	if err != nil {
		fmt.Println("An error has ocurred while parsing number:", err)
		os.Exit(1)
	}

	return number
}
