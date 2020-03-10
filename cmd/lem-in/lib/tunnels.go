package solver

import "fmt"

var tunnels []Point
var farm [][]string

var lenX int
var lenY int

var shiftX int
var shiftY int

// Point to represent point in farm where to place tunnel
type Point struct {
	x, y int
}

// Stack implementation
type Stack []Point

// Push to push new Point into stack and return stack
func (s Stack) Push(v Point) Stack {
	return append(s, v)
}

// Pop to remove last Point from stack and return it
func (s Stack) Pop() (Stack, Point) {
	// TODO: EMPTY STACK
	l := len(s)
	return s[:l-1], s[l-1]
}

// Shift to remove first Point from stack and return it
func (s Stack) Shift() (Stack, Point) {
	// TODO: EMPTY STACK
	return s[1:], s[0]
}

// ConnectRooms to connect two rooms for the current state of farm
func ConnectRooms() {

	for _, link := range graph.links {
		room1 := graph.rooms[link.room1]
		room2 := graph.rooms[link.room2]
		BFS(*room1, *room2)
	}

}

// MakeFarm to create 2d matrix with rooms
func MakeFarm() {

	var maxX int
	var maxY int
	var minX int
	var minY int

	shiftX := 0
	shiftY := 0

	for _, room := range graph.rooms {
		if room.x > maxX {
			maxX = room.x
		}
		if room.y > maxY {
			maxY = room.y
		}
	}

	for _, tunnel := range tunnels {
		if tunnel.x > maxX {
			maxX = tunnel.x
		}
		if tunnel.x < minX {
			minX = tunnel.x
		}
		if tunnel.y > maxY {
			maxY = tunnel.y
		}
		if tunnel.y < minY {
			minY = tunnel.y
		}
	}

	if minX < 0 {
		shiftX = -minX
	}
	if minY < 0 {
		shiftY = -minY
	}

	lenX = maxX - minX + 1
	lenY = maxY - minY + 1

	farm = make([][]string, lenY)
	for i := 0; i < lenY; i++ {
		farm[i] = make([]string, lenX)
	}

	for name, room := range graph.rooms {
		farm[room.y+shiftY][room.x+shiftX] = name
	}

	for _, tunnel := range tunnels {
		farm[tunnel.y+shiftY][tunnel.x+shiftX] = "X"
	}

	for i := range farm {
		for j := range farm[i] {
			if farm[i][j] != "" {
				fmt.Print(farm[i][j])
			} else {
				fmt.Print(".")

			}
		}
		fmt.Println()
	}

}

// BFS to apply bfs algorithm
func BFS(room1, room2 Room) {
	var start = Point{room1.x, room1.y}
	var end = Point{room2.x, room2.y}

	var stack Stack
	stack = stack.Push(start)

	var previousPointMap = make(map[Point]Point)

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {1, -1}, {-1, -1}, {-1, 1}}

	visited := []Point{start}

	for len(stack) > 0 {

		var point Point
		stack, point = stack.Shift()

		if point.x == end.x && point.y == end.y {
			break
		}

		for i := 0; i < len(dirs); i++ {
			newPoint := Point{point.x + dirs[i][0], point.y + dirs[i][1]}

			// for j := 0; j < len(dirs); j++ {
			// 	newPoint2 := Point{newPoint.x + dirs[j][0], newPoint.y + dirs[j][1]}
			// 	if IsValid(newPoint, visited, end) && IsValid(newPoint2, visited, end) {
			// 		stack = stack.Push(newPoint)
			// 		visited = append(visited, newPoint)
			// 		previousPointMap[newPoint] = point
			// 	}
			// }

			if IsValid(newPoint, visited, end) {
				stack = stack.Push(newPoint)
				visited = append(visited, newPoint)
				previousPointMap[newPoint] = point
			}

		}
	}

	var shortestPath []Point

	for ValueExists(previousPointMap, end) {
		point := previousPointMap[end]
		if point.x == start.x && point.y == start.y {
			break
		}
		shortestPath = append(shortestPath, point)
		end = point
	}

	tunnels = append(tunnels, shortestPath...)

}

// IsValid to check if given Point is valid
func IsValid(point Point, visited []Point, end Point) bool {
	return !ExiststInVisisted(point, visited) && ((point.x == end.x && point.y == end.y) || !BarrierExists(point.x, point.y))
}

// BarrierExists to check if any barrier exists on the way
func BarrierExists(x, y int) bool {
	for _, room := range graph.rooms {
		if room.x == x && room.y == y {
			return true
		}
	}

	for _, tunnel := range tunnels {
		if tunnel.x == x && tunnel.y == y {
			return true
		}
	}

	return false
}

// ValueExists to check if given Point exists in map
func ValueExists(parentNode map[Point]Point, point Point) bool {
	for pointInMap := range parentNode {
		if pointInMap.x == point.x && pointInMap.y == point.y {
			return true
		}
	}
	return false
}

// ExiststInVisisted to check if Point exists in visisted array of Points
func ExiststInVisisted(point Point, visited []Point) bool {
	for _, pointInVis := range visited {
		if point.x == pointInVis.x && point.y == pointInVis.y {
			return true
		}
	}
	return false

}
