package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func sumItemsOfElf(elfItems string) int {

	var totalKcal int

	items := strings.Split(elfItems, "\n")

	for _, itemKcal := range items {
		parsedCalories, _ := strconv.Atoi(itemKcal)
		totalKcal += parsedCalories
	}

	return totalKcal 
}

func getItemsByElf() []string {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Printf("An error has ocurred: %v", err)
		return []string{}
	}

	input := string(data)

	return strings.Split(input, "\n\n")
}

func getSumOfElfsWithMostKcal() int {

	itemsByElf := getItemsByElf()	

	totalKcalByElf := []int {sumItemsOfElf(itemsByElf[0])}

	for _, items := range itemsByElf[1:] {
		
		elfTotalKcal := sumItemsOfElf(items)

		for index, value := range totalKcalByElf{

			if value < elfTotalKcal{
				totalKcalByElf = slices.Insert(totalKcalByElf, index, elfTotalKcal)
				break
			}	

		}

	}

	return totalKcalByElf[0] + totalKcalByElf[1] + totalKcalByElf[2];
}

func main() {

	fmt.Println(getSumOfElfsWithMostKcal())

}
