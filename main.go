package main

import (
	"fmt"
	"io/ioutil"
	"os"

	solver "./lib"
)

var dataStr string

func main() {

	ReadFile()

	occured, message := solver.GetError()
	if occured {
		fmt.Println("ERROR: " + message)
		return
	}

	solver.ParseDataFromFile(dataStr)

	occured, message = solver.GetError()
	if occured {
		fmt.Println("ERROR: " + message)
		return
	}

	solver.ConnectRooms()

	solver.MakeFarm()

}

// ReadFile reads .txt file from argument and puts it into string
func ReadFile() {
	args := os.Args[1:]
	length := 0
	for i := range args {
		length = i + 1
	}

	if length == 0 {
		solver.SetError("File name missing")
		return
	}
	if length != 1 {
		solver.SetError("Too many arguments")
		return
	}

	fileName := args[0]

	file, err := os.Open(fileName)
	if err != nil {
		solver.SetError(err.Error())
		return
	}

	dataByteArr, err := ioutil.ReadAll(file)
	if err != nil {
		solver.SetError(err.Error())
		return
	}

	dataStr = string(dataByteArr)

}
