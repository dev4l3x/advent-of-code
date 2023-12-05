package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readInput(os.Args[1])
	lowestLocation := GetLowestLocation(input)
	fmt.Println("The lowest location is:", lowestLocation)
}

func readInput(fileName string) []string {
	file, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("An error has ocurred while reading input:", err)
		os.Exit(1)
	}

	return strings.Split(string(file), "\n\n")
}

func GetLowestLocation(textAlmanac []string) int {

	seeds := getSeeds(textAlmanac[0])
	almanac := processAlmanac(textAlmanac[1:])	

	var debug string

	var lowestLocation = getLocationFromValue(seeds[0], almanac, 0, &debug)

	fmt.Println("Seed", seeds[0], ":", debug)

	for _, seed := range seeds[1:] {
		debug = ""
		location := getLocationFromValue(seed, almanac, 0, &debug)
		fmt.Println("Seed", seed, ":", debug)
		if location < lowestLocation {
			lowestLocation = location
		}
	}

	return lowestLocation
}

func getLocationFromValue(sourceValue int, almanac [][][3]int, category int, debug *string) int {

	*debug += " -> " + fmt.Sprint(sourceValue)

	if (len(almanac) == category) {
		return sourceValue
	}


	currentCategory := almanac[category]

	for _, mapping := range currentCategory {
		sourceStartRange := mapping[1]
		rangeLenght := mapping[2]
		hasMapping := sourceStartRange <= sourceValue && sourceValue <= (sourceStartRange + rangeLenght - 1)

		if hasMapping {
			destinationStartRange := mapping[0]
			sourceValueForNextCategory := destinationStartRange + (sourceValue - sourceStartRange)
			return getLocationFromValue(sourceValueForNextCategory, almanac, category + 1, debug)
		}
	}

	return getLocationFromValue(sourceValue, almanac, category + 1, debug)
}

func getSeeds(seeds string) []int {
	var parsedSeeds []int = make([]int, 0)
	textSeeds := strings.Split(seeds, " ")
	for _, seed := range textSeeds[1:] {
		parsedSeeds = append(parsedSeeds, parseNumber(seed))
	}
	return parsedSeeds
}

func processAlmanac(almanac []string) [][][3]int {

	var processedAlmanac [][][3]int = make([][][3]int, 0)

	for _, almanacCategory := range almanac	{
		mappings := strings.Split(almanacCategory, "\n")[1:]
		var category [][3]int = make([][3]int, 0)
		for _, mapping := range mappings {
			values := strings.Split(mapping, " ")
			if len(values) != 3 {
				fmt.Println("Expected mappings to have 3 values: ", values)
				os.Exit(1)
			}
			numbers := [3]int {parseNumber(values[0]), parseNumber(values[1]), parseNumber(values[2])}
			category = append(category, numbers)
		}
		processedAlmanac = append(processedAlmanac, category)
	}

	return processedAlmanac
}

func parseNumber(number string) int {
	n, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println("An error has ocurred while parsing number:", err)
		os.Exit(1)
	}
	return n
}
