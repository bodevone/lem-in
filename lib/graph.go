package solver

import "fmt"

// Graph connects all together
type Graph struct {
	rooms map[string]*Room
	links []Link
	start Room
	end   Room
	paths []Path
	ants  int
}

// Path represents array of array of connected rooms
type Path struct {
	path   []Room
	weight int
}

// Room represents room with x and y values
type Room struct {
	name       string
	x, y       int
	neighbours []Room
}

// Link represents links between two rooms
type Link struct {
	room1 string
	room2 string
}

// InitGraph to make map filed in Graph
func InitGraph() {
	graph.rooms = make(map[string]*Room)
}

// AddNeighbours adds adjacent Rooms to the given Room
func AddNeighbours() {
	for _, link := range graph.links {
		graph.rooms[link.room1].neighbours = append(graph.rooms[link.room1].neighbours, *graph.rooms[link.room2])
		graph.rooms[link.room2].neighbours = append(graph.rooms[link.room2].neighbours, *graph.rooms[link.room1])
	}

	for _, room := range graph.rooms {
		fmt.Print(room.name + ": ")
		for _, r := range room.neighbours {
			fmt.Print(r.name + " ")
		}
		fmt.Println()
	}
}
