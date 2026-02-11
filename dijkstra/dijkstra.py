graph = {
    'A': [
        {
            'node': 'B',
            'weight': 5
        },
        {
            'node': 'C',
            'weight': 2
        },
    ],
    'B': [
        {
            'node': 'A',
            'weight': 5
        },
        {
            'node': 'D',
            'weight': 7
        },
    ],
    'C': [
        {
            'node': 'A',
            'weight': 2
        },
    ],
    'D': [
        {
            'node': 'B',
            'weight': 7
        },
    ],
}

POSITIVE_INF = float('inf')

class Shortest_Path_Unit:
    def __init__(self, node: str, distance: float = POSITIVE_INF, previous_node: str = None):
        self.node = node
        self.distance = distance
        self.previous_node = previous_node

def dijkstra(vertex: str, visited: list[str], tracking_structure: dict, graph: dict):
    neighborhood = graph[vertex]

    current_vertex = tracking_structure.get(vertex)

    if not current_vertex:
        tracking_structure[vertex] = Shortest_Path_Unit(vertex, 0)
        current_vertex = tracking_structure[vertex]

    next_vertex = None
    
    for neighbor in neighborhood:
        neighbor_node = neighbor['node']
        neighbor_distance = neighbor['weight']

        if neighbor_node not in visited:
            neighbor_metadata = tracking_structure.get(neighbor_node)

            if not neighbor_metadata:
                tracking_structure[neighbor_node] = Shortest_Path_Unit(neighbor_node, current_vertex.distance + neighbor_distance, vertex)
            else:
                if (current_vertex.distance + neighbor_distance) < neighbor_metadata.distance:
                    tracking_structure[neighbor_node] = Shortest_Path_Unit(neighbor_node, current_vertex.distance + neighbor_distance, vertex)
            
            if not next_vertex:
                next_vertex = tracking_structure.get(neighbor_node)
            else:
                if tracking_structure.get(neighbor_node).distance < next_vertex.distance:
                    next_vertex = tracking_structure.get(neighbor_node)
    
    if vertex not in visited:
        visited.append(vertex)

    if len(visited) == len(graph.keys()):
        return

    if not next_vertex:
        if current_vertex.previous_node:
            dijkstra(current_vertex.previous_node, visited, tracking_structure, graph)
    else:
        dijkstra(next_vertex.node, visited, tracking_structure, graph)


tracking_structure = {}

dijkstra("B", [], tracking_structure, graph)

tracking_structure_keys = tracking_structure.keys()

for key in tracking_structure_keys:
    shortest_path_unit = tracking_structure[key]
    print(f"vertex: {shortest_path_unit.node} | short_distance: {shortest_path_unit.distance} | previous_node: {shortest_path_unit.previous_node}")