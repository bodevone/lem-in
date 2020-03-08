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

// PathsComb represents combinations of paths which do not intersect
type PathsComb struct {
	paths  []Path
	weight int
}

// Path represents array of array of connected rooms
type Path struct {
	path   []string
	weight int
}

// Room represents room with x and y values
type Room struct {
	name       string
	x, y       int
	neighbours []string
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
		graph.rooms[link.room1].neighbours = append(graph.rooms[link.room1].neighbours, link.room2)
		graph.rooms[link.room2].neighbours = append(graph.rooms[link.room2].neighbours, link.room1)
	}

	for _, room := range graph.rooms {
		fmt.Print(room.name + ": ")
		for _, r := range room.neighbours {
			fmt.Print(r + " ")
		}
		fmt.Println()
	}

	graph.start = *graph.rooms[graph.start.name]
	graph.end = *graph.rooms[graph.end.name]
}

var visited []string

// FindPaths to find all possible paths
func FindPaths() {

	path := []string{}
	dfs(graph.start.name, path)

	fmt.Print("Paths: ")
	fmt.Println(graph.paths)
	// for _, p := range graph.paths {
	// 	for _, r := range p.path {
	// 		fmt.Print(r)
	// 		fmt.Print(" ")
	// 	}
	// 	fmt.Println()
	// }
	return

}

func dfs(roomStr string, path []string) {
	room := *graph.rooms[roomStr]
	path = append(path, roomStr)
	if roomStr == graph.end.name {
		var newPath []string
		for _, s := range path {
			newPath = append(newPath, s)
		}
		p := Path{newPath, len(newPath) - 1}
		graph.paths = append(graph.paths, p)
		return
	}

	visited = append(visited, roomStr)
	for _, r := range room.neighbours {
		if !InVisited(r) {
			dfs(r, path)
		}
	}

	visited = visited[:len(visited)-1]

	return

}

// InVisited to check if Room exists in visited
func InVisited(room string) bool {
	for _, r := range visited {
		if r == room {
			return true
		}
	}
	return false
}

// FindPathComb to find all combinations of paths
func FindPathComb() {

}
