package solver

import (
	"fmt"
	"sort"
	"strconv"
)

// Graph connects all together
type Graph struct {
	rooms      map[string]*Room
	links      []Link
	start      Room
	end        Room
	paths      []Path
	pathsCombs []PathsComb
	chosenComb PathsComb
	ants       int
	mapPaths   map[int]Path
	decision   map[int]int
	iterations int
}

// PathsComb represents combinations of paths which do not intersect
type PathsComb struct {
	paths      []Path
	weight     int
	iterations int
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

// Error in order to handle error
type Error struct {
	message string
	occured bool
}

var error Error

// SetError to define found Error
func SetError(message string) {
	error.occured = true
	error.message = message
}

// GetError to check if Error exists
func GetError() (bool, string) {
	return error.occured, error.message
}

var graph Graph

// InitGraph to make map filed in Graph
func InitGraph() {
	graph.rooms = make(map[string]*Room)
	maxIters = 1000
}

// CheckError to check for errors
func CheckError() {
	if graph.start.name == "" || graph.end.name == "" {
		SetError("No start or end defined")
		return
	}

	for _, link := range graph.links {
		if InRooms(link.room1) && InRooms(link.room2) {
			if link.room1 == link.room2 {
				SetError("Room links to itself")
				return
			}
		} else {
			SetError("Link to unknown room")
			return
		}
	}

}

// AddNeighbours adds adjacent Rooms to the given Room
func AddNeighbours() {
	for _, link := range graph.links {
		graph.rooms[link.room1].neighbours = append(graph.rooms[link.room1].neighbours, link.room2)
		graph.rooms[link.room2].neighbours = append(graph.rooms[link.room2].neighbours, link.room1)
	}

	// for _, room := range graph.rooms {
	// 	fmt.Print(room.name + ": ")
	// 	for _, r := range room.neighbours {
	// 		fmt.Print(r + " ")
	// 	}
	// 	fmt.Println()
	// }

	graph.start = *graph.rooms[graph.start.name]
	graph.end = *graph.rooms[graph.end.name]

}

// StartEndConnected to check if start and end rooms are connected
func StartEndConnected() bool {
	for _, link := range graph.links {
		if (link.room1 == graph.start.name && link.room2 == graph.end.name) || (link.room2 == graph.start.name && link.room1 == graph.end.name) {
			return true
		}
	}
	return false
}

// PrintStartEnd to print result if start and end connected
func PrintStartEnd() {
	for i := 1; i <= graph.ants; i++ {
		if i > 1 {
			fmt.Print(" ")
		}
		fmt.Print("L" + strconv.Itoa(i) + "-" + graph.end.name)
	}
	fmt.Println()
}

var visited []string

// FindPaths to find all possible paths
func FindPaths() {

	path := []string{}
	dfs(graph.start.name, path)

	// fmt.Print("Paths: ")
	// fmt.Println(graph.paths)
	// for _, p := range graph.paths {
	// 	for _, r := range p.path {
	// 		fmt.Print(r)
	// 		fmt.Print(" ")
	// 	}
	// 	fmt.Println()
	// }

	graph.mapPaths = make(map[int]Path)
	for _, p := range graph.paths {
		graph.mapPaths[p.index] = p
	}

	if len(graph.paths) == 0 {
		SetError("No connections between start and end rooms")
	}

}

var index int
var maxIters int
var iters int

func dfs(roomName string, path []string) {
	// if iters >= maxIters {
	// 	SetError("Invalid data format")
	// 	return
	// }
	// iters++

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

//PrintAll to print number of ants, rooms and links
func PrintAll() {

	fmt.Println(graph.ants)

	fmt.Println("##start")
	fmt.Print(graph.start.name)
	fmt.Print(" ")
	fmt.Print(graph.start.x)
	fmt.Print(" ")
	fmt.Print(graph.start.y)
	fmt.Println()

	for r := range graph.rooms {
		if r != graph.start.name || r != graph.end.name {
			fmt.Print(graph.rooms[r].name)
			fmt.Print(" ")
			fmt.Print(graph.rooms[r].x)
			fmt.Print(" ")
			fmt.Print(graph.rooms[r].y)
			fmt.Println()
		}
	}

	fmt.Println("##end")
	fmt.Print(graph.end.name)
	fmt.Print(" ")
	fmt.Print(graph.end.x)
	fmt.Print(" ")
	fmt.Print(graph.end.y)
	fmt.Println()

	for _, link := range graph.links {
		fmt.Print(link.room1)
		fmt.Print("-")
		fmt.Print(link.room2)
		fmt.Println()
	}

	fmt.Println()

}

// FindSolution to find solution
func FindSolution() {
	var finalDecision map[int]int
	minIterations := 9999999
	var finalComb PathsComb

	for _, comb := range graph.pathsCombs {
		decision := make(map[int]int)

		for i := 1; i <= graph.ants; i++ {
			minn := 99999999
			minnIndex := 0
			for _, p := range comb.paths {
				if minn == 0 {
					minn = p.weight + decision[p.index]
				}

				if p.weight+decision[p.index] < minn {
					minn = p.weight + decision[p.index]
					minnIndex = p.index
				}

			}
			decision[minnIndex]++
		}

		maxx := 0
		for index, ants := range decision {
			if graph.mapPaths[index].weight+ants > maxx {
				maxx = graph.mapPaths[index].weight + ants - 1
			}
		}
		if maxx < minIterations {
			minIterations = maxx
			finalDecision = decision
			finalComb = comb
		}

	}

	graph.chosenComb = finalComb
	graph.decision = finalDecision
	graph.iterations = minIterations

	// fmt.Println(graph.iterations)
}

// GetSolution to print output in required format
func GetSolution() {

	answer := make([]string, graph.iterations)
	newDecision := make(map[int][]int)

	decision := graph.decision
	ant := 1

	maxx := 0
	for i, length := range decision {
		if decision[i] > 0 {
			if length > maxx {
				maxx = length
			}
		}
	}

	for i := 0; i < maxx; i++ {
		for pathIndex := range decision {
			if decision[pathIndex] > i {
				newDecision[pathIndex] = append(newDecision[pathIndex], ant)
				ant++
			}
		}
	}

	ants := make([]int, graph.ants)
	antsPath := make([]int, graph.ants)
	movingAnts := []int{}

	for i := 0; i < graph.iterations; i++ {
		for pathIndex, arr := range newDecision {
			if len(newDecision[pathIndex]) > 0 {
				if i < len(arr) {
					movingAnts = append(movingAnts, arr[i])
					antsPath[arr[i]-1] = pathIndex

				}
			}

		}

		movingAntsNext := []int{}
		for _, ant := range movingAnts {
			ants[ant-1]++

			if ants[ant-1] == graph.mapPaths[antsPath[ant-1]].weight {
				continue
			} else {
				movingAntsNext = append(movingAntsNext, ant)
			}
		}

		notFirst := false
		for _, ant := range movingAnts {
			if notFirst {
				answer[i] += " "
			}
			answer[i] += "L" + strconv.Itoa(ant) + "-" + string(graph.mapPaths[antsPath[ant-1]].path[ants[ant-1]])
			notFirst = true
		}

		movingAnts = movingAntsNext
	}
	for _, ans := range answer {
		fmt.Println(ans)

	}

}
