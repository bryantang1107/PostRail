package move

type Move struct {
	Time       int
	TrainName  string
	StartNode  string
	EndNode    string
	PickedUp   []string
	DroppedOff []string
}

func NewMove(time int, trainName string, startNode string, endNode string, pickedUp []string, droppedOff []string) *Move {
	return &Move{
		Time:       time,
		TrainName:  trainName,
		StartNode:  startNode,
		EndNode:    endNode,
		PickedUp:   pickedUp,
		DroppedOff: droppedOff,
	}
}
