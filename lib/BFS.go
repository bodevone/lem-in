package solver

import "fmt"

// Stack implementation
type Stack []Room

// Push to push new Room into stack and return stack
func (s Stack) Push(v Room) Stack {
	return append(s, v)
}

// Pop to remove last Room from stack and return it
func (s Stack) Pop() (Stack, Room) {
	// TODO: EMPTY STACK
	l := len(s)
	return s[:l-1], s[l-1]
}

// Shift to remove first Room from stack and return it
func (s Stack) Shift() (Stack, Room) {
	// TODO: EMPTY STACK
	return s[1:], s[0]
}

// BFS to apply bfs algorithm
func BFS(room1, room2 Room) {
	var stack Stack
	stack = stack.Push(room1)

	fmt.Println(stack)
}
