package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/bryantang1107/PostRail/edge"
	"github.com/bryantang1107/PostRail/move"
	"github.com/bryantang1107/PostRail/node"
	"github.com/bryantang1107/PostRail/parcel"
	"github.com/bryantang1107/PostRail/train"
	"github.com/bryantang1107/PostRail/util"
)

var (
	trains          []train.Train
	parcels         []parcel.Parcel
	graph           map[int][]edge.Edge // Adjacency list of graph edges
	moves           []move.Move
	parcelOwnership = make(map[int]int) // parcelId --> trainId
	ownershipLock   sync.Mutex
)

func initGraph() {
	nodeA := node.NewNode(1, "A")
	nodeB := node.NewNode(2, "B")
	nodeC := node.NewNode(3, "C")

	graph = map[int][]edge.Edge{
		1: {
			{Name: "Edge 1-2", FromNode: *nodeA, ToNode: *nodeB, JourneyTime: 30},
		},
		2: {
			{Name: "Edge 2-3", FromNode: *nodeB, ToNode: *nodeC, JourneyTime: 10},
			{Name: "Edge 2-1", FromNode: *nodeB, ToNode: *nodeA, JourneyTime: 30},
		},
		3: {
			{Name: "Edge 3-2", FromNode: *nodeC, ToNode: *nodeB, JourneyTime: 10},
		},
	}

	// Initialize parcels
	parcels = []parcel.Parcel{
		{ID: 0, Name: "K1", StartNode: *nodeA, EndNode: *nodeC, Weight: 5, Delivered: false, PickedUp: false},
		{ID: 1, Name: "K2", StartNode: *nodeB, EndNode: *nodeA, Weight: 100, Delivered: false, PickedUp: false},
	}

	// Initialize trains
	trains = []train.Train{
		{ID: 0, Name: "Q1", Capacity: 6, CurrentWeight: 0, Load: []int{}, Current: *nodeB},
		{ID: 1, Name: "Q2", Capacity: 150, CurrentWeight: 0, Load: []int{}, Current: *nodeA},
	}

	for i := range parcels {
		fmt.Printf("Parcel: %s Located at Station: %s\n", parcels[i].Name, parcels[i].StartNode.Name)
		parcelOwnership[parcels[i].ID] = -1
	}
}

func main() {
	initGraph()

	// Start delivery
	startDelivery()
}

func startDelivery() {
	var wg sync.WaitGroup
	allDelivered := false
	totalElapsed := 0

	for !allDelivered {
		allDelivered = true

		// Launch each train in its own goroutine
		for i := range trains {
			wg.Add(1)                     // add goroutine count
			go func(train *train.Train) { // use pointer if state will mutate
				defer wg.Done() // decrement goroutine count

				var pickedUp, droppedOff []string
				fromNode := train.Current

				for i := range parcels {
					p := &parcels[i]

					ownershipLock.Lock()
					isUnclaimed := parcelOwnership[p.ID] == -1
					isAssignedToThisTrain := parcelOwnership[p.ID] == train.ID
					ownershipLock.Unlock()

					if isUnclaimed || isAssignedToThisTrain {
						// Lock package
						p.Mutex.Lock()

						// If train at same station as package
						if !p.Delivered && !p.PickedUp && p.StartNode == train.Current {
							// if train did not exceed capacity
							if train.CurrentWeight+p.Weight <= train.Capacity {
								// assign ownership - Good to have to keep track of which train owns the parcel
								ownershipLock.Lock()
								parcelOwnership[p.ID] = train.ID
								ownershipLock.Unlock()

								train.CurrentWeight += p.Weight
								p.PickedUp = true
								train.Load = append(train.Load, p.ID)
								pickedUp = append(pickedUp, p.Name)
								fmt.Printf("Train Name: %s picked up Parcel: %s at Station: %s\n", train.Name, p.Name, p.StartNode.Name)
							} else {
								fmt.Printf("Train Name: %s unable to pick up Parcel: %s at Station: %s\n", train.Name, p.Name, p.StartNode.Name)
							}
						}
						p.Mutex.Unlock()
					}
				}

				// Move train to next node
				nextNode := getNextNode(train.Current.ID, train.Previous.ID, train.Load)
				fmt.Printf("Train : %s, Current Station: %s, Next Station: %s\n", train.Name, train.Current.Name, nextNode.Name)
				if nextNode != nil {
					e := findEdge(train.Current.ID, nextNode.ID)
					if e != nil {
						// Move to next node without locking edges
						train.Previous = train.Current
						train.Current = *nextNode

						// Drop off parcels
						for j := 0; j < len(train.Load); j++ {
							pID := train.Load[j]
							p := &parcels[pID]

							p.Mutex.Lock()
							if train.Current == p.EndNode {
								fmt.Printf("Train Name: %s dropped off Parcel: %s at Station: %s \n", train.Name, p.Name, p.EndNode.Name)
								p.Delivered = true
								train.CurrentWeight -= p.Weight
								train.Load = append(train.Load[:j], train.Load[j+1:]...)
								j-- // Adjust index after removing from slice

								// Release ownership after delivery
								ownershipLock.Lock()
								delete(parcelOwnership, p.ID)
								ownershipLock.Unlock()

								droppedOff = append(droppedOff, p.Name)
							}

							p.Mutex.Unlock()
						}

						moves = append(moves, *move.NewMove(totalElapsed, train.Name, fromNode.Name, e.ToNode.Name, pickedUp, droppedOff))
						totalElapsed += e.JourneyTime
					}
				}
			}(&trains[i])
		}

		// Wait for all trains to finish their actions
		wg.Wait()

		// Check delivery status
		for _, p := range parcels {
			if !p.Delivered {
				allDelivered = false
				break
			}
		}

		if allDelivered {
			fmt.Printf("All Parcels delivered in %d minutes.\n", totalElapsed)

			// Log all moves
			for _, move := range moves {
				util.PrintMoves(move.Time, move.TrainName, move.StartNode, move.PickedUp, move.EndNode, move.DroppedOff)
			}
		} else {
			time.Sleep(3 * time.Second) // Simulate time between actions
		}
	}
}

func getNextNode(current, previous int, load []int) *node.Node {
	// Implement logic to determine the next node (e.g., Dijkstra)
	edges := graph[current]
	for _, edge := range edges {
		// Avoid backtracking unless necessary
		if edge.ToNode.ID != previous {
			return &edge.ToNode
		}
	}

	// If no alternative path, backtrack
	for _, edge := range edges {
		if edge.ToNode.ID == previous {
			return &edge.ToNode
		}
	}

	// No path available
	return nil
}

// Find an edge between two nodes
func findEdge(fromNodeID, toNodeID int) *edge.Edge {
	for _, edge := range graph[fromNodeID] {
		if edge.ToNode.ID == toNodeID {
			return &edge
		}
	}
	return nil
}
