package main

import (
	"fmt"
	"slices"
)

func bfs(vertex string, queue *[]string, visited *[]string, graph map[string][]string) {
	if !slices.Contains(*visited, vertex) {
		*visited = append(*visited, vertex)
	}

	neighborhood := graph[vertex]

	for _, neighbor := range neighborhood {
		if !slices.Contains(*visited, neighbor) && !slices.Contains(*queue, neighbor) {
			*queue = append(*queue, neighbor)
		}
	}

	if len(*queue) > 0 {
		nextVertex := (*queue)[0]
		*queue = (*queue)[1:]

		bfs(nextVertex, queue, visited, graph)
	}
}

func main() {
	graphExample := map[string][]string{
		"N":  {"S", "NO", "NE"},
		"S":  {"N", "SO", "SE"},
		"O":  {"L", "NO", "SO"},
		"L":  {"O", "NE", "SE"},
		"NO": {"N", "O"},
		"NE": {"N", "L"},
		"SO": {"S", "O"},
		"SE": {"S", "L"},
	}

	myQueue := []string{}
	result := []string{}

	bfs("N", &myQueue, &result, graphExample)

	fmt.Println(result)
}
