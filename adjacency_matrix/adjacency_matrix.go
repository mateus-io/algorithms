package main

import (
	"fmt"
)

type Edge struct {
	From int
	To   int
}

func main() {
	edges := []Edge{{1, 4}, {2, 7}}
	nodes := []int{1, 4, 2, 7}
	numberOfNodes := len(nodes)

	adjacencyMatrix := make([][]int, numberOfNodes)
	for i := range adjacencyMatrix {
		adjacencyMatrix[i] = make([]int, numberOfNodes)
	}

	fmt.Println("Matriz inicial:", adjacencyMatrix)

	for i := range numberOfNodes {
		line := nodes[i]
		for j := range numberOfNodes {
			column := nodes[j]

			for _, edge := range edges {
				if edge.From == line && edge.To == column {
					adjacencyMatrix[i][j] = 1
				}
			}
		}
	}

	fmt.Println("Matriz final:", adjacencyMatrix)
}
