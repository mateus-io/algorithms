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

def dijkstra(start_vertex: str, graph: dict) -> dict:
    vertex = start_vertex
    visited = []
    tracking_structure = {}

    while len(visited) < len(graph.keys()):

        neighborhood = graph[vertex]

        current_vertex_metadata = tracking_structure.get(vertex)

        if not current_vertex_metadata:
            empty_shortest_path_unit = Shortest_Path_Unit(vertex, 0)
            tracking_structure[vertex] = empty_shortest_path_unit
            current_vertex_metadata = empty_shortest_path_unit

        next_vertex_metadata: Shortest_Path_Unit = None
        
        for neighbor in neighborhood:
            neighbor_node = neighbor['node']
            neighbor_distance = neighbor['weight']

            if neighbor_node not in visited:
                neighbor_metadata = tracking_structure.get(neighbor_node)

                if not neighbor_metadata:
                    tracking_structure[neighbor_node] = Shortest_Path_Unit(neighbor_node, current_vertex_metadata.distance + neighbor_distance, vertex)
                else:
                    if (current_vertex_metadata.distance + neighbor_distance) < neighbor_metadata.distance:
                        tracking_structure[neighbor_node] = Shortest_Path_Unit(neighbor_node, current_vertex_metadata.distance + neighbor_distance, vertex)
                
                if not next_vertex_metadata:
                    next_vertex_metadata = tracking_structure.get(neighbor_node)
                else:
                    neighbor_node_metadata = tracking_structure.get(neighbor_node)
                    if neighbor_node_metadata.distance < next_vertex_metadata.distance:
                        next_vertex_metadata = neighbor_node_metadata
        
        if vertex not in visited:
            visited.append(vertex)
        
        if not next_vertex_metadata:
            if current_vertex_metadata.previous_node:
                vertex = current_vertex_metadata.previous_node
        else:
            vertex = next_vertex_metadata.node
    return tracking_structure

result = dijkstra("B", graph)

result_keys = result.keys()

for key in result_keys:
    shortest_path_unit = result[key]
    print(f"vertex: {shortest_path_unit.node} | short_distance: {shortest_path_unit.distance} | previous_node: {shortest_path_unit.previous_node}")