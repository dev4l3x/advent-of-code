package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const (
	UP = 1 
	DOWN = -1 
)

func main() {

	file, error := os.ReadFile("input.txt")	

	if error != nil {
		fmt.Printf("An error has ocurred while processing the input: %v", error)
		os.Exit(1)
	}

	input := string(file)

	// Remove deactivated muls

	deactivtedRegex := `don't\(\)(.|\n)*?do\(\)`
	dre := regexp.MustCompile(deactivtedRegex)
	activatedExpressions := dre.ReplaceAllString(input, "")

	dre2 := regexp.MustCompile(`don't\(\)(.|\n)*`)
	activatedExpressions = dre2.ReplaceAllString(activatedExpressions, "")

	fmt.Printf("Activated expressions: %v\n", activatedExpressions)


	// Calculate activated muls

	regex := `mul\(([0-9]{1,3}),([0-9]{1,3})\)`

	regexp := regexp.MustCompile(regex)
	regexMatches := regexp.FindAllStringSubmatch(activatedExpressions, -1)

	multiplicationsSum := 0

	for _, match := range regexMatches {

		firstNumber := parseNumber(match[1])
		secondNumber := parseNumber(match[2])

		multiplication := firstNumber * secondNumber

		multiplicationsSum += multiplication

	}

	fmt.Printf("The sum of all multiplications is: %v\n", multiplicationsSum)

}

func parseNumber(number string) int {
	
	n, err := strconv.Atoi(number)

	if err != nil {
		log.Fatalf("An error has ocurred while parsing the number: %v", err)
	}

	return n
}