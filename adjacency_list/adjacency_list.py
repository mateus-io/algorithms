class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
    
edges = [('A', 'L'), ('A', 'N'), ('N', 'B'), ('B', 'N'), ('L', 'A'), ('L', 'B')]
nodes = ['A', 'N', 'B', 'L']

def build_adjacency_list(edges, nodes):
    adjacency_list = []

    for node in nodes:
        head = ListNode(node)
        adjacency_list.append(head)
        current = head
        for target in nodes:
            if (head.val, target) in edges:
                current.next = ListNode(target)
                current = current.next
    
    return adjacency_list

adjacency_list = build_adjacency_list(edges, nodes)

for list_node in adjacency_list:
    current = list_node
    while current:
        print(current.val, end=" -> " if current.next else "\n")
        current = current.next