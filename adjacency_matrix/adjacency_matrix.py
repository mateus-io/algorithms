edges = [(1, 4), (2, 7)]
nodes = [1, 4, 2, 7]
number_of_nodes = len(nodes)

adjacency_matrix = [[0] * number_of_nodes for _ in range(number_of_nodes)]

print(adjacency_matrix)

for i in range(number_of_nodes):
    line = nodes[i]
    for j in range(number_of_nodes):
        column = nodes[j]
        edge_draft = (line, column)
        if edge_draft in edges:
            adjacency_matrix[i][j] = 1

print(adjacency_matrix)