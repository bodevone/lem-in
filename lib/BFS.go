package solver

// Point to represent point in farm where to place tunnel
type Point struct {
	x, y int
}

// Stack implementation
type Stack []Point

var tunnels []Point

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

// BFS to apply bfs algorithm
func BFS(room1, room2 Room) {
	var start = Point{room1.x, room1.y}
	var end = Point{room2.x, room2.y}

	var stack Stack
	stack = stack.Push(start)

	var previousPointMap = make(map[Point]Point)

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {1, -1}, {-1, -1}, {-1, 1}}

	visited := []Point{start}

	//TODO: Remove 2d array and just write array of rooms and tunnels
	for len(stack) > 0 {

		var point Point
		stack, point = stack.Shift()

		if point.x == end.x && point.y == end.y {
			break
		}

		for i := 0; i < len(dirs); i++ {
			newPoint := Point{point.x + dirs[i][0], point.y + dirs[i][1]}
			if !ExiststInVisisted(newPoint, visited) && ((newPoint.x == end.x && newPoint.y == end.y) || !BarrierExists(newPoint.x, newPoint.y)) {
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

// BarrierExists to check if any barrier exists on the way
func BarrierExists(x, y int) bool {
	for _, room := range rooms {
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
