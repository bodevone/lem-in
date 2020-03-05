package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Room represents room with row, col and name values
type Room struct {
	name     string
	row, col int
}

// Link represents links between two rooms
type Link struct {
	room1 string
	room2 string
}

var dataStr string
var rooms []Room
var links []Link

func main() {

	ReadFile()

	ParseDataFromFile()

	fmt.Println(rooms)
	fmt.Println(links)

}

// ReadFile reads .txt file from argument and puts it into string
func ReadFile() {
	args := os.Args[1:]
	length := 0
	for i := range args {
		length = i + 1
	}

	if length == 0 {
		fmt.Println("File name missing")
		return
	}
	if length != 1 {
		fmt.Println("Too many arguments")
		return
	}

	fileName := args[0]

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	dataByteArr, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	dataStr = string(dataByteArr)

}

//ParseDataFromFile parsies data from string obtained from .txt file into Structs of Rooms and Links
func ParseDataFromFile() {

	length := len(dataStr)

	i := 0
	temp := ""
	spaceCount := 0
	xInt, yInt := 0, 0
	var name string
	var a []string
	for i < length {
		if dataStr[i] == '#' {
			for dataStr[i] != '\n' {
				i++
			}
			i++
		} else {
			for i < length && dataStr[i] != '\n' {
				temp += string(dataStr[i])
				if dataStr[i] == ' ' {
					spaceCount++
				}
				i++
			}
			i++
			if spaceCount == 0 {
				a = strings.Split(temp, "-")
				links = append(links, Link{a[0], a[1]})
			} else if spaceCount == 2 {
				a = strings.Split(temp, " ")
				name = a[0]
				xInt, _ = strconv.Atoi(a[1])
				yInt, _ = strconv.Atoi(a[2])
				rooms = append(rooms, Room{name, xInt, yInt})
			}
			temp = ""
			spaceCount = 0
		}

	}

}
