package solver

import "fmt"

// Error in order to handle error
type Error struct {
	message string
	occured bool
}

// Room represents room with x and y values
type Room struct {
	x, y int
}

// Link represents links between two rooms
type Link struct {
	room1 string
	room2 string
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

var rooms = make(map[string]Room)
var links []Link
var farm [][]string

var lenX int
var lenY int

var shiftX int
var shiftY int

// ConnectRooms to connect two rooms for the current state of farm
func ConnectRooms() {
	// link := links[0]
	// room1 := rooms[link.room1]
	// room2 := rooms[link.room2]
	// BFS(room1, room2)

	for _, link := range links {
		room1 := rooms[link.room1]
		room2 := rooms[link.room2]
		BFS(room1, room2)
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

	for _, room := range rooms {
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

	for name, room := range rooms {
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
