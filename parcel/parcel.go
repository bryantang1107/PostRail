package parcel

import (
	"sync"

	"github.com/bryantang1107/PostRail/node"
)

type Parcel struct {
	ID        int
	Name      string
	Weight    int       // in kg
	StartNode node.Node // start node id
	EndNode   node.Node // end node id
	Delivered bool
	PickedUp  bool
	Mutex     sync.Mutex
}
