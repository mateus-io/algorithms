package main

import (
	"fmt"
	"slices"
)

func dfs(vertex string, visited *[]string, graph map[string][]string) {
	*visited = append(*visited, vertex)

	neighborhood := graph[vertex]
	for _, neighbor := range neighborhood {
		if !slices.Contains(*visited, neighbor) {
			dfs(neighbor, visited, graph)
		}
	}
}

func main() {
	graph := map[string][]string{
		"1":  {"10"},
		"10": {"1", "25", "30"},
		"25": {"10"},
		"30": {"10", "40", "45"},
		"40": {"30"},
		"45": {"30"},
	}

	result := []string{}

	dfs("1", &result, graph)

	fmt.Println(result)
}
