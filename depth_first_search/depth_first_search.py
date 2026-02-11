graph_life_of_ivan_ilitch = {
    "1": ["10"],
    "10": ["1", "25", "30"],
    "25": ["10"],
    "30": ["10", "40", "45"],
    "40": ["30"],
    "45": ["30"],
}

def dfs(vertex: str, visited: list, graph: dict)-> None:
    neighborhood = graph[vertex]
    visited.append(vertex)
    for neighbor in neighborhood:
        if neighbor not in visited:
            dfs(neighbor, visited, graph)

result = list()
dfs("1", result, graph_life_of_ivan_ilitch)
print(result)
