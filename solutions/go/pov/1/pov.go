package pov

// Tree represents a node in a tree.
type Tree struct {
	value    string
	children []*Tree
}

// New creates and returns a new Tree with the given root value and children.
func New(value string, children ...*Tree) *Tree {
	return &Tree{
		value:    value,
		children: children,
	}
}

// Value returns the value at the root of a tree.
func (tr *Tree) Value() string {
	if tr == nil {
		return ""
	}
	return tr.value
}

// Children returns a slice containing the children of a tree.
// There is no need to sort the elements in the result slice,
// they can be in any order.
func (tr *Tree) Children() []*Tree {
	if tr == nil {
		return nil
	}
	return tr.children
}

// String describes a tree in a compact S-expression format.
// This helps to make test outputs more readable.
func (tr *Tree) String() string {
	if tr == nil {
		return "nil"
	}
	result := tr.Value()
	if len(tr.Children()) == 0 {
		return result
	}
	for _, ch := range tr.Children() {
		result += " " + ch.String()
	}
	return "(" + result + ")"
}

// FindNode searches for a node with the given value in the tree.
func (tr *Tree) FindNode(value string) *Tree {
	if tr == nil {
		return nil
	}
	if tr.value == value {
		return tr
	}
	for _, child := range tr.children {
		if res := child.FindNode(value); res != nil {
			return res
		}
	}
	return nil
}

// buildGraph creates an undirected adjacency list representation of the tree.
// This allows us to traverse freely in any direction (up to parents or down to children).
func (tr *Tree) buildGraph() map[string][]string {
	graph := make(map[string][]string)

	var dfs func(*Tree)
	dfs = func(node *Tree) {
		if node == nil {
			return
		}
		for _, child := range node.children {
			if child != nil {
				graph[node.value] = append(graph[node.value], child.value)
				graph[child.value] = append(graph[child.value], node.value)
				dfs(child)
			}
		}
	}

	dfs(tr)
	return graph
}

// FromPov returns the pov from the node specified in the argument.
// It creates a new tree re-oriented with the specified node as the root.
func (tr *Tree) FromPov(from string) *Tree {
	if tr.FindNode(from) == nil {
		return nil
	}

	graph := tr.buildGraph()

	var buildTree func(current, parent string) *Tree
	buildTree = func(current, parent string) *Tree {
		node := &Tree{value: current}
		for _, neighbor := range graph[current] {
			if neighbor != parent {
				node.children = append(node.children, buildTree(neighbor, current))
			}
		}
		return node
	}

	return buildTree(from, "")
}

// PathTo returns the shortest path between two nodes in the tree.
func (tr *Tree) PathTo(from, to string) []string {
	if tr.FindNode(from) == nil || tr.FindNode(to) == nil {
		return nil
	}

	graph := tr.buildGraph()

	// BFS to find the shortest path
	queue := [][]string{{from}}
	visited := make(map[string]bool)
	visited[from] = true

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]

		current := path[len(path)-1]
		if current == to {
			return path
		}

		for _, neighbor := range graph[current] {
			if !visited[neighbor] {
				visited[neighbor] = true
				// Create a new path slice to avoid mutating the original path in the queue
				newPath := append([]string{}, path...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}

	return nil
}
