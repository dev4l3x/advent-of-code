package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	defer timer()()
	input := readInput(os.Args[1])
	steps := CalculateVerificationNumber(input)
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

func CalculateVerificationNumber(initializationSequence string) int {
	
	steps := strings.Split(initializationSequence, ",")

	verificationNumber := 0

	for _, step := range steps {
		encodedStep := hash(step)
		verificationNumber += encodedStep
	}

	return verificationNumber
}

func hash(step string) int {
	asciiStep := []byte(step)
	currentValue := 0

	for _, character := range asciiStep {
		code := int(character)
		currentValue += code
		currentValue *= 17
		currentValue %= 256
	}

	return currentValue
}


