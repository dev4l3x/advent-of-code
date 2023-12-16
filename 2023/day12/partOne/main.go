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


		springs, numbers := extractSpringsAndNumbers(line)

		unknownGroupsRegex := regexp.MustCompile("\\?+")
		unknownGroups := unknownGroupsRegex.FindAllString(springs, -1)

		numbersRegex := createRegexForNumbers(numbers)		

		groupsPossibilities := make([][]string, len(unknownGroups))

		for index, group := range unknownGroups {
			groupsPossibilities[index] = getPossibleGroupsForUnknown(len(group))
		}

		counters := make([]int, len(groupsPossibilities))

		var lineSum int

		for counters[0] < len(groupsPossibilities[0]) {

			combinationSprings := springs

			for i := 0; i < len(counters); i++{
				combinationSprings = strings.Replace(combinationSprings, unknownGroups[i], groupsPossibilities[i][counters[i]], 1)
			}

			if numbersRegex.MatchString(combinationSprings) {
				lineSum++
			}

			for i := len(counters) - 1 ; i >= 0 ; i-- {
				if counters[i] + 1 < len(groupsPossibilities[i]) || i == 0 {
					counters[i]++
					break
				} else if i != 0 && counters[i] >= len(groupsPossibilities[i]) - 1 {
					counters[i] = 0	
				}
			}
		}

		fmt.Println("Line:", line, "->", lineSum)
		totalSum += lineSum

	}


	return totalSum
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

func extractSpringsAndNumbers(line string) (string, []int) {
	lineSplitted := strings.Split(line, " ")

	numbersAsText := strings.Split(lineSplitted[1], ",")

	numbers := make([]int, 0)

	for _, number := range numbersAsText {
		numbers = append(numbers, parseNumber(number))
	}

	return lineSplitted[0], numbers
}

func parseNumber(n string) int {
	number, err := strconv.Atoi(n)

	if err != nil {
		fmt.Println("An error has ocurred while parsing number:", err)
		os.Exit(1)
	}

	return number
}
