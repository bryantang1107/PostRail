package edge

import (
	"github.com/bryantang1107/PostRail/node"
)

type Edge struct {
	Name             string
	FromNode, ToNode node.Node
	JourneyTime      int // in minutes
}
