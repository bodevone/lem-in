package solver

import (
	"fmt"
)

// Room represents room with x and y values
type Room struct {
	x, y int
}

// Link represents links between two rooms
type Link struct {
	room1 string
	room2 string
}

var rooms = make(map[string]Room)
var links []Link
var farm [][]string

// ConnectRooms to connect two rooms for the current state of farm
func ConnectRooms() {
	link := links[0]
	room1 := rooms[link.room1]
	room2 := rooms[link.room2]

	BFS(room1, room2)

	fmt.Println(room1)
	fmt.Println(room2)

}

// MakeFarm to create 2d matrix with rooms
func MakeFarm() {
	var maxX int
	var maxY int

	for _, room := range rooms {
		if room.x > maxX {
			maxX = room.x
		}
		if room.y > maxY {
			maxY = room.y
		}
	}

	maxX++
	maxY++

	farm = make([][]string, maxY)
	for i := 0; i < maxY; i++ {
		farm[i] = make([]string, maxX)
	}

	for name, room := range rooms {
		farm[room.y][room.x] = name
	}

	fmt.Println(farm)

}
