package main

import (
	"fmt"
	"math"
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

var seedLocation map[int]int = make(map[int]int)


func GetLowestLocation(textAlmanac []string) int {

	seeds := getSeeds(textAlmanac[0])
	fmt.Println("Built seeds")
	almanac := processAlmanac(textAlmanac[1:])	

	var lowestLocation = getLocationFromValue(seeds[0], almanac, 0)

	fmt.Println("Seeds:", seeds)

	for _, seed := range seeds[1:] {

		var location int

		location = getLocationFromValue(seed, almanac, 0)

		if location < lowestLocation {
			lowestLocation = location
		}
	}

	return lowestLocation
}

func getLocationFromValue(seedRange[2]int, almanac [][][3]int, category int) int {

	currentCategory := almanac[category]

	if category == len(almanac) - 1 {
		lastCategoryMappings := getMappingRangesInCategory(seedRange, currentCategory)
		lowestLocation := math.MaxInt
		for _, mapping := range lastCategoryMappings {
			location := mapping[0]
			if location < lowestLocation {
				lowestLocation = location
			}
		}
		return lowestLocation
	}

	currentCategoryMappings := getMappingRangesInCategory(seedRange, currentCategory)
	
	lowestLocation := math.MaxInt
	for _, mapping := range currentCategoryMappings {
		location := getLocationFromValue(mapping, almanac, category + 1)
		if location < lowestLocation {
			lowestLocation = location
		}
	}
	return lowestLocation
}

func getMappingRangesInCategory(seedRange [2]int, category [][3]int) [][2]int {
	newSeedRanges := [][2]int {}
	remainingRanges := true
	for _, mapping := range category {
		startMappingSource := mapping[1]
		endMappingSource := mapping[1] + mapping[2] - 1

		if startMappingSource <= seedRange[0]  && seedRange[1] <= endMappingSource {
			targetStartMapping := mapping[0] + (seedRange[0] - startMappingSource)
			targetEndMapping := mapping[0] + (seedRange[1] - startMappingSource)
			newSeedRanges = append(newSeedRanges, [2]int {targetStartMapping, targetEndMapping})
			remainingRanges = false
			break
		} else if startMappingSource <= seedRange[0] && seedRange[0] <= endMappingSource && endMappingSource < seedRange[1]{
			targetStartMapping := mapping[0] + (seedRange[0] - startMappingSource)
			targetEndMapping := mapping[0] + mapping[2] - 1
			newSeedRanges = append(newSeedRanges, [2]int {targetStartMapping, targetEndMapping})
			seedRange[0] = endMappingSource + 1
		} else if seedRange[0] < startMappingSource && seedRange[1] <= endMappingSource && seedRange[1] >= startMappingSource {
			targetStartMapping := mapping[0]
			targetEndMapping := mapping[0] + (seedRange[1] - startMappingSource)
			newSeedRanges = append(newSeedRanges, [2]int {targetStartMapping, targetEndMapping})
			seedRange[1] = startMappingSource - 1
		} 
	}
	if remainingRanges {
		newSeedRanges = append(newSeedRanges, seedRange)
	}
	return newSeedRanges
}

func getSeeds(seeds string) [][2]int {
	textSeedsRanges := strings.Split(seeds, " ")[1:]
	ranges := [][2]int{}
	for i := 0 ; i < len(textSeedsRanges) ; i+=2 {
		startSeed := parseNumber(textSeedsRanges[i])
		seedRangeLength := parseNumber(textSeedsRanges[i+1])
		endSeed := startSeed + seedRangeLength - 1
		ranges = append(ranges, [2]int {startSeed, endSeed})
	}

	fmt.Println(ranges)

	return ranges
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

	fmt.Println("Built almanac")

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
