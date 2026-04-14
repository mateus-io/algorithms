package main

import (
	"fmt"
	"slices"
)

type ShortestPathUnit struct {
	Node         string
	Distance     int
	PreviousNode string
}

type Neighbor struct {
	Node   string
	Weight int
}

func dijkstra(startVertex string, graph map[string][]Neighbor) map[string]*ShortestPathUnit {
	vertex := startVertex
	visited := []string{}
	trackingStructure := make(map[string]*ShortestPathUnit)

	for len(visited) < len(graph) {
		neighborhood := graph[vertex]
		currentMetadata, exists := trackingStructure[vertex]

		if !exists {
			unit := &ShortestPathUnit{Node: vertex, Distance: 0, PreviousNode: ""}
			trackingStructure[vertex] = unit
			currentMetadata = unit
		}

		var nextVertexMetadata *ShortestPathUnit = nil

		for _, neighbor := range neighborhood {
			neighborNode := neighbor.Node
			neighborDistance := neighbor.Weight

			if !slices.Contains(visited, neighborNode) {
				neighborMetadata, found := trackingStructure[neighborNode]

				newDistance := currentMetadata.Distance + neighborDistance

				if !found {
					trackingStructure[neighborNode] = &ShortestPathUnit{neighborNode, newDistance, vertex}
					neighborMetadata = trackingStructure[neighborNode]
				} else {
					if newDistance < neighborMetadata.Distance {
						neighborMetadata.Distance = newDistance
						neighborMetadata.PreviousNode = vertex
					}
				}

				if nextVertexMetadata == nil || neighborMetadata.Distance < nextVertexMetadata.Distance {
					nextVertexMetadata = neighborMetadata
				}
			}
		}

		if !slices.Contains(visited, vertex) {
			visited = append(visited, vertex)
		}

		if nextVertexMetadata == nil {
			if currentMetadata.PreviousNode != "" {
				vertex = currentMetadata.PreviousNode
			}
		} else {
			vertex = nextVertexMetadata.Node
		}
	}

	return trackingStructure
}

func main() {
	graph := map[string][]Neighbor{
		"A": {{Node: "B", Weight: 5}, {Node: "C", Weight: 2}},
		"B": {{Node: "A", Weight: 5}, {Node: "D", Weight: 7}},
		"C": {{Node: "A", Weight: 2}},
		"D": {{Node: "B", Weight: 7}},
	}

	result := dijkstra("B", graph)

	for key, shortestPathUnit := range result {
		fmt.Printf("vertex: %s | short_distance: %d | previous_node: %s\n",
			key, shortestPathUnit.Distance, shortestPathUnit.PreviousNode)
	}
}
