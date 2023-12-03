package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	input := readInput(os.Args[1])

	fmt.Println("The sum of the part numbers is:", GetPartNumbersSumFromEngine(input))
}

func readInput(fileName string) string {
	file, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("An error has ocurred while reading file:", fileName)
		os.Exit(1)
	}

	return string(file)
}

func GetPartNumbersSumFromEngine(engine string) int {
	lines := strings.Split(engine, "\n")

	var sum int

	for lineNumber, line := range lines {
		fmt.Println(line)
		chars := strings.Split(line, "")
		var isProcessingNumber bool
		var startNumber int
		for index, char := range chars {
			if !isProcessingNumber && isNumber(char) {
				isProcessingNumber = true
				startNumber = index
			} 
			if isProcessingNumber && isNumber(char) && (index == len(chars) - 1 || !isNumber(chars[index + 1])) {
				isProcessingNumber = false	
				var number [3]int
				number[0] = parseNumber(line[startNumber:index+1])
				if hasSymbolAdjacent(lineNumber, startNumber, index, lines) {
					sum += number[0]
					fmt.Println("Number", number[0], "is adjacent")
				} else {
					fmt.Println("Number", number[0], "is not adjacent")
				}
			} 
		}

	}

	return sum;	
}

func hasSymbolAdjacent(line int, start int, end int, lines []string) bool {
	for i := max(0, line - 1) ; i <= line + 1 && i < len(lines) ; i++ {
		for j := max(0, start - 1) ; j <= end + 1 && j < len(lines[i]) ; j++ {
			if isSymbol(lines[i][j:j+1]) {
				return true
			}				
		}
	}
	return false
}

func max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func isSymbol(symbol string) bool {
	regex, _ := regexp.Compile("[0-9.]")
	return !regex.MatchString(symbol)
}

func parseNumber(number string) int {
	n, err := strconv.Atoi(number)

	if  err != nil {
		fmt.Println("An error has ocurred: ", err)
		os.Exit(1)
	}

	return n
}

func isNumber(number string) bool {
	_, err := strconv.Atoi(number)
	return err == nil
}
