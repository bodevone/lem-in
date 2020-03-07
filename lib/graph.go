package solver

// Graph connects all together
type Graph struct {
	rooms map[string]Room
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
	name string
	x, y int
}

// Link represents links between two rooms
type Link struct {
	room1 string
	room2 string
}

// InitGraph to make map filed in Graph
func InitGraph() {
	graph.rooms = make(map[string]Room)
}
