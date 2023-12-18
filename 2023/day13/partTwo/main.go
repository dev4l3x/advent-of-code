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
	steps := SummarizeNotes(input)
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

func SummarizeNotes(notes string) int {
	patterns := strings.Split(notes, "\n\n")

	var summary int
	
	for _, pattern := range patterns {

		fmt.Println("Pattern:")
		fmt.Println(pattern)

		l, r := getVerticalMirrorPoint(pattern)	

		if l != -1 && r != -1 {
			summary += l + 1	
			fmt.Println("Vertical mirror:", l, r)
			continue
		}

		u, d := getHorizontalMirrorPoint(pattern)

		
		if u != -1 && d != -1 {
			fmt.Println("Horizontal mirror:", u, d)
			summary += 100 * (u + 1)
		}
		
		if u == -1 && l == -1 && d == -1 && r == -1 {
			fmt.Println("No pattern found!!")
			os.Exit(1)
		}
	}
	

	return summary
}

func getVerticalMirrorPoint(pattern string) (int, int) {

	patternLines := strings.Split(pattern, "\n")


	for i, j := 0, 1; j < len(patternLines[0]) ; i, j = i+1, j+1 {
		possibleChanges := 1
		if isVerticalMirrorPoint(patternLines, [2]int {i, j}, &possibleChanges) && possibleChanges == 0{
			return i, j
		}
	}

	return -1, -1
}

func isVerticalMirrorPoint(patternLines []string, point [2]int, possibleChanges *int) bool {

	if *possibleChanges == 0 && patternLines[0][point[0]] != patternLines[0][point[1]] {
		// It's not a mirror point, we can continue the loop
		return false	
	}

	if isRowMirror(patternLines[0], point, possibleChanges) {

		if len(patternLines) == 1 {
			return true
		}
		return isVerticalMirrorPoint(patternLines[1:], point, possibleChanges)	
	}

	return false
}

func isRowMirror(patternLine string, point [2]int, possibleChanges *int) bool {

	for left, right := point[0], point[1] ; left >= 0 && right < len(patternLine) ; left, right = left-1, right+1 {
		if *possibleChanges == 0 && patternLine[left] != patternLine[right] {
			return false
		} else if patternLine[left] != patternLine[right]{
			*possibleChanges--
		}
	}

	return true
}

func getHorizontalMirrorPoint(pattern string) (int, int) {

	patternLines := strings.Split(pattern, "\n")


	for i, j := 0, 1; j < len(patternLines) ; i, j = i+1, j+1 {
		possibleChanges := 1
		if isHorizontalMirrorPoint(patternLines, 0, [2]int {i, j}, &possibleChanges) && possibleChanges == 0 {
			return i, j
		}
	}

	return -1, -1
}

func isHorizontalMirrorPoint(patternLines []string, currentCol int, point [2]int, possibleChanges *int) bool {

	if *possibleChanges == 0 && patternLines[point[0]][currentCol] != patternLines[point[1]][currentCol] {
		return false
	}

	if isColMirror(patternLines, currentCol, point, possibleChanges) {
		if currentCol == len(patternLines[0]) - 1 {
			return true
		}
		return isHorizontalMirrorPoint(patternLines, currentCol + 1, point, possibleChanges)
	}

	return false

}

func isColMirror(patternLines []string, col int, point [2]int, possibleChanges *int) bool {

	for u, l := point[0], point[1] ; u >= 0 && l < len(patternLines) ; u, l = u-1, l+1 {
		if *possibleChanges == 0 && patternLines[u][col] != patternLines[l][col] {
			return false
		} else if patternLines[u][col] != patternLines[l][col] {
			(*possibleChanges)--
		}
	}

	return true
}
