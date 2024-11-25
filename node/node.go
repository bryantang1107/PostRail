package node

// Station
type Node struct {
	ID   int
	Name string
}

func NewNode(id int, name string) *Node {
	return &Node{
		ID:   id,
		Name: name,
	}
}
