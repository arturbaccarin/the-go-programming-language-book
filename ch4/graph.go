package ch4

// The value typ e of a map can its elf be a composite typ e, such as a map or slice

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}

	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
