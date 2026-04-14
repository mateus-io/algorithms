package main

import "fmt"

type ListNode struct {
	Val  string
	Next *ListNode
}

func buildAdjacencyList(edges [][2]string, nodes []string) []*ListNode {
	edgeMap := make(map[[2]string]bool)
	for _, edge := range edges {
		edgeMap[edge] = true
	}

	var adjacencyList []*ListNode

	for _, nodeVal := range nodes {
		head := &ListNode{Val: nodeVal}
		adjacencyList = append(adjacencyList, head)
		current := head

		for _, target := range nodes {

			if edgeMap[[2]string{nodeVal, target}] {
				current.Next = &ListNode{Val: target}
				current = current.Next
			}
		}
	}

	return adjacencyList
}

func main() {
	edges := [][2]string{
		{"A", "L"}, {"A", "N"}, {"N", "B"},
		{"B", "N"}, {"L", "A"}, {"L", "B"},
	}
	nodes := []string{"A", "N", "B", "L"}

	adjacencyList := buildAdjacencyList(edges, nodes)

	for _, head := range adjacencyList {
		current := head
		for current != nil {
			if current.Next != nil {
				fmt.Printf("%s -> ", current.Val)
			} else {
				fmt.Printf("%s\n", current.Val)
			}
			current = current.Next
		}
	}
}
