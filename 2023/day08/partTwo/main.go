package main

import (
	"fmt"
	"math"
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

	instructions, rootNodes := createGraphFromNetwork(network)

	totalSteps := 1 
	for _, node := range rootNodes {
			
		steps := countStepsToFinal(node, instructions)

		if totalSteps % steps != 0 {
			totalSteps = lcm(totalSteps, steps)
		}
	}

	return totalSteps
}

func lcm(a int, b int) int {
	return a * (b / gcd(a, b));
}

func gcd(a int, b int) int {

	greatest := int(math.Max(float64(a), float64(b)))
	lowest := int(math.Min(float64(a), float64(b)))

	for greatest != lowest {
		newNumber := greatest - lowest
		greatest = int(math.Max(float64(newNumber), float64(lowest)))
		lowest = int(math.Min(float64(newNumber), float64(lowest)))
	}

	return greatest
}

func countStepsToFinal(node *Node, instructions []string) int {

	currentNode := node
	steps := 0
	currentInstruction := 0

	for !areAllNodesWithZ(currentNode) {

		fmt.Println("Current NOdes:", currentNode)

		direction := instructions[currentInstruction]

		currentNode = getNodesByDirection(direction, currentNode)

		currentInstruction++
		if currentInstruction >= len(instructions) {
			currentInstruction = 0
		}

		steps++
	}

	return steps
}

func getNodesByDirection(direction string, node *Node) *Node {
	if direction == "R" {
		return (*node).right
	}
	return (*node).left
}

func areAllNodesWithZ(node *Node) bool {
	if !strings.HasSuffix((*node).value, "Z") {
		return false
	}

	return true
}

func createGraphFromNetwork(network string) ([]string, []*Node) {
	networkParsed := strings.Split(network, "\n\n")
	instructions := strings.Split(networkParsed[0], "")
	nodes := strings.Split(networkParsed[1], "\n")
	existingNodes := make(map[string]*Node)

	var rootNodes []*Node = make([]*Node, 0)

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

		if strings.HasSuffix(current, "A") {
			rootNodes = append(rootNodes, currentNode)
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
	return instructions, rootNodes
}
