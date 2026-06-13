package deps

type Node struct {
	Name string
	Path string
}

type Graph struct {
	Nodes map[string]*Node
	Edges map[string]map[string]bool
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*Node),
		Edges: make(map[string]map[string]bool),
	}
}

func (g *Graph) AddNode(name, path string) {
	if _, exists := g.Nodes[name]; !exists {
		g.Nodes[name] = &Node{Name: name, Path: path}
	}
}

func (g *Graph) AddDependency(from, to string) {
	if _, exists := g.Edges[from]; !exists {
		g.Edges[from] = make(map[string]bool)
	}
	g.Edges[from][to] = true
}

func (g *Graph) DetectCycles() [][]string {
	visited := make(map[string]bool)
	recStack := make(map[string]bool)
	var cycles [][]string

	var dfs func(node string, path []string)
	dfs = func(node string, path []string) {
            if recStack[node] {
        start := -1
        for i, n := range path {
            if n == node {
                start = i
                break
            }
        }
        if start != -1 {
            cycle := append([]string(nil), path[start:]...)
            cycle = append(cycle, node)
            cycles = append(cycles, cycle)
        }
        return
    }
		if visited[node] {
			return
		}
		visited[node] = true
		recStack[node] = true
		path = append(path, node)

		for neighbor := range g.Edges[node] {
			dfs(neighbor, path)
		}
		recStack[node] = false
	}

	for node := range g.Nodes {
		if !visited[node] {
			dfs(node, []string{})
		}
	}
	return cycles
}

func (g *Graph) HasCycles() bool {
	return len(g.DetectCycles()) > 0
}