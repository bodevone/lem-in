package main

import (
	"fmt"
	"io/ioutil"
	"os"

	solver "./lib"
)

// Error in order to handle error
type Error struct {
	message string
	occured bool
}

var dataStr string

// var rooms []Room
var error Error

func main() {

	error.occured = false

	ReadFile()

	if error.occured {
		fmt.Println(error.message)
		return
	}

	solver.ParseDataFromFile(dataStr)

	solver.MakeFarm()

	solver.ConnectRooms()

}

// ReadFile reads .txt file from argument and puts it into string
func ReadFile() {
	args := os.Args[1:]
	length := 0
	for i := range args {
		length = i + 1
	}

	if length == 0 {
		error.occured = true
		error.message = "File name missing"
		return
	}
	if length != 1 {
		error.occured = true
		error.message = "Too many arguments"
		return
	}

	fileName := args[0]

	file, err := os.Open(fileName)
	if err != nil {
		error.occured = true
		error.message = err.Error()
		return
	}

	dataByteArr, err := ioutil.ReadAll(file)
	if err != nil {
		error.occured = true
		error.message = err.Error()
		return
	}

	dataStr = string(dataByteArr)

}
