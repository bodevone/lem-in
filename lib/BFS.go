package solver

import (
	"fmt"
)

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

// BFS to apply bfs algorithm
func BFS(room1, room2 Room) {
	var start = Point{room1.x, room1.y}
	var end = Point{room2.x, room2.y}

	var stack Stack
	stack = stack.Push(start)

	var previousPointMap = make(map[Point]Point)

	dirs := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}, {-1, 1}, {-1, -1}, {1, 1}, {1, -1}}
	//dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	visited := []Point{start}

	//TODO: Remove 2d array and just write array of rooms and tunnels
	for len(stack) > 0 {
		// fmt.Println(stack)
		var point Point
		stack, point = stack.Shift()
		// fmt.Println(point)
		if point.x == end.x && point.y == end.y {
			break
		}
		for i := 0; i < len(dirs); i++ {
			newPoint := Point{point.x + dirs[i][0], point.y + dirs[i][1]}
			if newPoint.x >= 0 && newPoint.x < lenX && newPoint.y >= 0 && newPoint.y < lenY && !ExiststInVisisted(newPoint, visited) {
				if (newPoint.x == end.x && newPoint.y == end.y) || farm[newPoint.y][newPoint.x] == "" {
					stack = stack.Push(newPoint)
					visited = append(visited, newPoint)
					previousPointMap[newPoint] = point
					fmt.Println(previousPointMap)
				}
			}
		}
	}

	var shortestPath []Point

	for ValueExists(previousPointMap, end) {
		point := previousPointMap[end]
		shortestPath = append(shortestPath, point)
		end = point
	}

	for _, point := range shortestPath[:len(shortestPath)-1] {
		farm[point.y][point.x] = "X"
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
