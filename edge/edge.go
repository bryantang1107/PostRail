package edge

import (
	"github.com/bryantang1107/PostRail/node"
)

type Edge struct {
	Name             string
	FromNode, ToNode node.Node
	JourneyTime      int // in minutes
}

func NewEdge(name string, fromNode node.Node, toNode node.Node, journeyTime int) *Edge {
	return &Edge{
		Name:        name,
		FromNode:    fromNode,
		ToNode:      toNode,
		JourneyTime: journeyTime,
	}
}
