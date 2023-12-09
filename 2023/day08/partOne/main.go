package main

import (
	"fmt"
	"os"
	"strings"
)

type Node struct {
	value string
	left, right *Node
}

func main() {
	input := readInput(os.Args[1])
	steps := GetStepsFromNetwork(input)
	fmt.Println("The answer is:", steps)
}

func readInput(fileName string) string {
	file, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("An error has ocurred while reading input:", err)
		os.Exit(1)
	}

	return string(file)
}

func GetStepsFromNetwork(network string) int {

	instructions, rootNode := createGraphFromNetwork(network)

	steps := countStepsToFinal(rootNode, 0, instructions, 0)

	return steps
}

func countStepsToFinal(node *Node, steps int, instructions []string, currentInstruction int) int {

	if (*node).value == "ZZZ" {
		return steps
	}

	direction := instructions[currentInstruction]

	currentInstruction++
	if currentInstruction >= len(instructions) {
		currentInstruction = 0
	}

	var nextNode *Node
	if direction == "R" {
		nextNode = (*node).right
	} else {
		nextNode = (*node).left
	}

	return countStepsToFinal(nextNode, steps + 1, instructions, currentInstruction)		
}


func createGraphFromNetwork(network string) ([]string, *Node) {
	networkParsed := strings.Split(network, "\n\n")
	instructions := strings.Split(networkParsed[0], "")
	nodes := strings.Split(networkParsed[1], "\n")
	existingNodes := make(map[string]*Node)

	var rootNode *Node

	for _, node := range nodes {
		parsedNode := strings.Split(node, " = ")
		current := parsedNode[0]
		connections := strings.Split(parsedNode[1], ", ")

		currentNode, exists := existingNodes[current]

		if !exists {
			newNode := Node{current, nil, nil}
			existingNodes[current] = &newNode
			currentNode = &newNode
		}

		if current == "ZZZ" {
			continue
		}

		if current == "AAA" {
			rootNode = currentNode
		}

		leftConnection := connections[0][1:]
		rightConnection := connections[1][:len(connections[1]) - 1]

		leftNode, leftOk := existingNodes[leftConnection]
		if !leftOk {
			newNode := Node{leftConnection, nil, nil}
			existingNodes[leftConnection] = &newNode
			leftNode = &newNode
		}

		rightNode, rightOk := existingNodes[rightConnection]
		if !rightOk{
			newNode := Node{rightConnection, nil, nil}
			existingNodes[rightConnection] = &newNode
			rightNode = &newNode
		}

		(*currentNode).left = leftNode
		(*currentNode).right = rightNode 
	}
	return instructions, rootNode
}
