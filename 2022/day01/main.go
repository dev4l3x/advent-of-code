package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func sumElfTotal(elfBag string) int {

	var total int
	calories := strings.Split(elfBag, "\n")

	for _, calorie := range calories {
		parsedCalories, _ := strconv.Atoi(calorie)
		total += parsedCalories
	}

	return total
}

func main() {

	data, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Printf("An error has ocurred: %v", err)
		return
	}

	input := string(data)

	elfs := strings.Split(input, "\n\n")

	elfsTotals := []int {sumElfTotal(elfs[0])}

	for _, elf := range elfs[1:] {
		
		elfTotal := sumElfTotal(elf)

		for index, value := range elfsTotals {

			if value < elfTotal {
				elfsTotals = slices.Insert(elfsTotals, index, elfTotal)
				break
			}	

		}

	}

	fmt.Println(elfsTotals[0] + elfsTotals[1] + elfsTotals[2])
}
