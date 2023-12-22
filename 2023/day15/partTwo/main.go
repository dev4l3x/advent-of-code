package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	defer timer()()
	input := readInput(os.Args[1])
	steps := CalculateFocusingPower(input)
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

func CalculateFocusingPower(initializationSequence string) int {
	
	steps := strings.Split(initializationSequence, ",")

	boxes := make([][]focalLen, 256)

	for _, step := range steps {

		label, op, count := parseStep(step)
		boxIndex := hash(label)
		box := &boxes[boxIndex]

		if box == nil {
			*box = make([]focalLen, 0)
		}

		fmt.Println("Processing:", label, op, count)

		if op == "-" {
			removeLensWithLabel(box, label)	
		} else {
			addLenToBox(box, focalLen{label, count})	
		}
		printMatrix(boxes)
	}

	focusingPower := 0

	for boxIndex, box := range boxes {
		for lenIndex, l := range box {
			lenPower := (1 + boxIndex) * (1 + lenIndex) * l.number
			focusingPower += lenPower
		}
	}

	return focusingPower
}

func printMatrix(matrix [][]focalLen) {
	for i, row := range matrix {
		if len(row) > 0 {
			fmt.Println(i, "->", row)
		}
	}
}

func addLenToBox(box *[]focalLen, l focalLen) {

	lenIndex := indexOf(*box, l.label)

	if lenIndex == -1 {
		*box = append(*box, l)
		return
	}

	(*box)[lenIndex] = l
}

func removeLensWithLabel(box *[]focalLen, label string) {
	lenIndex := indexOf(*box, label)
	if lenIndex != -1 {
		left := (*box)[:lenIndex]
		right := (*box)[lenIndex+1:]
		*box = append(left, right...)

	}
}

func indexOf(box []focalLen, label string) int {
	for i, l := range box {
		if l.label == label {
			return i
		}
	}	
	return -1
}


func parseStep(step string) (label string, op string, count int) {

	label, _, found := strings.Cut(step, "-")

	if found {
		return label, "-", -1 
	}

	label, countText, _ := strings.Cut(step, "=")

	return label, "=", parseNumber(countText)
}

func parseNumber(n string) int {
	number, err := strconv.Atoi(n)

	if err != nil {
		os.Exit(1)
		fmt.Println("Error parsing number")
	}

	return number
}

type focalLen struct {
	label string
	number int
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


