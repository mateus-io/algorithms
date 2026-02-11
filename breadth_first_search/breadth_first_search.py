graph_example = {
    "N": ["S", "NO", "NE"],
    "S": ["N", "SO", "SE"],
    "O": ["L", "NO", "SO"],
    "L": ["O", "NE", "SE"],
    "NO": ["N", "O"],
    "NE": ["N", "L"],
    "SO": ["S", "O"],
    "SE": ["S", "L"],
}


from collections import deque

def bfs(vertex: str, queue: deque, visited: list, graph: dict) -> None:
    if vertex not in visited:
        visited.append(vertex)
    
    neighborhood = graph[vertex]

    for neighbor in neighborhood:
        if neighbor not in visited:
            queue.append(neighbor)
    
    if queue:
        next_vertex = queue.popleft()
        if next_vertex:
            bfs(next_vertex, queue, visited, graph)

my_queue = deque()
result = list()

bfs("N", my_queue, result, graph_example)

print(result)