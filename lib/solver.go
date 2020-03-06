package solver

import (
	"fmt"
)

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

// ConnectRooms to connect two rooms for the current state of farm
func ConnectRooms() {
	link := links[0]
	room1 := rooms[link.room1]
	room2 := rooms[link.room2]

	BFS(room1, room2)

}

// MakeFarm to create 2d matrix with rooms
func MakeFarm() {

	for _, room := range rooms {
		if room.x > lenX {
			lenX = room.x
		}
		if room.y > lenY {
			lenY = room.y
		}
	}

	lenX++
	lenY++

	farm = make([][]string, lenY)
	for i := 0; i < lenY; i++ {
		farm[i] = make([]string, lenX)
	}

	for name, room := range rooms {
		farm[room.y][room.x] = name
	}

	fmt.Println(farm)

}
