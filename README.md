PostRail: A Simulation of an Autonomous Train Delivery System

PostRail is a Go-based simulation project that models an autonomous train delivery system. Designed to demonstrate the use of concurrency in Go, the system orchestrates trains to deliver parcels across a network of nodes connected by edges. The project integrates essential concepts like synchronization, mutex locks, and goroutines to simulate real-time decision-making and multi-train operations.

Key Features:
Concurrency and Synchronization:

Utilizes goroutines for simulating multiple trains operating simultaneously.
Ensures thread-safe operations using mutex locks for parcel ownership and train actions.
Dynamic Parcel Management:

Parcels are dynamically assigned to trains based on location and capacity.
Avoids duplication or conflicts in parcel handling through ownership tracking.
Graph-Based Routing:

The rail network is represented as an adjacency list of nodes (stations) and edges (tracks).
Trains autonomously select their next destination using algorithms that minimize backtracking while prioritizing pending deliveries.
Train Load Balancing:

Trains have weight and capacity constraints, influencing which parcels they can carry.
Supports multiple trains working in parallel without interference.
Simulated Real-Time Operations:

Includes time delays for train movements and parcel handling to mimic real-world scenarios.
Logs every train's actions, including parcel pickup, drop-off, and movement between nodes.
Failure Handling:

Prevents deadlocks where trains might stall due to capacity or routing conflicts.
Parcels are re-evaluated for assignment if a train is unable to handle them.
Goals of PostRail:
Concurrency Demonstration: Showcase the power of Go's lightweight threading and synchronization mechanisms.
Scalability: Allow for easy addition of nodes, edges, trains, and parcels to simulate larger, more complex networks.
Real-Time Simulation: Mimic a live autonomous system with decision-making based on current states and priorities.
Education: Serve as a learning tool for developers exploring concurrent programming, Go, and graph-based problem-solving.
Potential Use Cases:
Research on autonomous logistics and delivery systems.
Benchmarking concurrency models in distributed simulations.
Teaching Go concepts like goroutines, mutex locks, and concurrent data handling.
PostRail bridges simulation with real-world problem-solving, making it a powerful tool for understanding both theoretical and practical aspects of autonomous systems in distributed environments.
