package solver

import (
	"fmt"
	"sort"
)

// Graph connects all together
type Graph struct {
	rooms      map[string]*Room
	links      []Link
	start      Room
	end        Room
	paths      []Path
	pathsCombs []PathsComb
	ants       int
}

// PathsComb represents combinations of paths which do not intersect
type PathsComb struct {
	paths  []Path
	weight int
}

// Path represents array of array of connected rooms
type Path struct {
	index  int
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

var index int

func dfs(roomName string, path []string) {
	room := *graph.rooms[roomName]
	path = append(path, roomName)
	if roomName == graph.end.name {
		var newPath []string
		for _, s := range path {
			newPath = append(newPath, s)
		}
		p := Path{index, newPath, len(newPath) - 1}
		index++
		graph.paths = append(graph.paths, p)
		return
	}

	visited = append(visited, roomName)
	for _, r := range room.neighbours {
		if !InVisited(r) {
			dfs(r, path)
		}
	}

	visited = visited[:len(visited)-1]

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

// FindPathsCombn to find all combinations of paths
func FindPathsCombn() {
	allPaths := graph.paths
	for i, p1 := range allPaths {
		comb := []Path{p1}
		for j, p2 := range allPaths {
			if i == j {
				continue
			}
			same := false

			for _, p3 := range comb {
				found := make(map[string]bool)
				for _, roomName := range p2.path[1 : len(p2.path)-1] {
					found[roomName] = true
				}
				for _, roomName := range p3.path[1 : len(p3.path)-1] {
					if found[roomName] {
						same = true
						break
					}
				}
			}

			if !same {
				comb = append(comb, p2)
			}

		}
		pathsComb := PathsComb{paths: comb}
		SortComb(pathsComb)
		// graph.pathsCombs = append(graph.pathsCombs, pathsComb)

		if !InComb(pathsComb) {
			graph.pathsCombs = append(graph.pathsCombs, pathsComb)
		}
	}
	for _, comb := range graph.pathsCombs {
		fmt.Println(comb.paths)
	}
}

// SortComb to sort paths in combination by index and weight
func SortComb(comb PathsComb) {
	sort.Slice(comb.paths, func(i, j int) bool {
		return comb.paths[i].weight < comb.paths[j].weight
	})

	sort.Slice(comb.paths, func(i, j int) bool {
		return comb.paths[i].index < comb.paths[j].index
	})

}

//InComb to check if given combination is already in graph combinations
func InComb(comb1 PathsComb) bool {
	found := false
	for _, comb2 := range graph.pathsCombs {
		if len(comb1.paths) == len(comb2.paths) {
			for i := range comb1.paths {
				if comb1.paths[i].index == comb2.paths[i].index {
					found = true
				} else {
					found = false
				}
			}
			if found {
				return true
			}
		}
	}
	return false
}
