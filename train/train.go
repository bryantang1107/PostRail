package train

import (
	"github.com/bryantang1107/PostRail/node"
)

type Train struct {
	ID            int
	Name          string
	Capacity      int       // in kg
	CurrentWeight int       // in kg
	Current       node.Node // Current node ID
	Previous      node.Node // Previous node ID
	Load          []int
}
