package main

import (
	"fmt"

	solver "./lib"
)

func main() {

	dataStr := solver.ReadFile()

	occured, message := solver.GetError()
	if occured {
		fmt.Println("ERROR: " + message)
		return
	}

	solver.InitGraph()

	solver.ParseDataFromFile(dataStr)

	occured, message = solver.GetError()
	if occured {
		fmt.Println("ERROR: " + message)
		return
	}

	solver.AddNeighbours()

	solver.FindPaths()

	solver.FindPathsCombn()

	solver.FindSolution()

	solver.GetIters()

	// solver.ConnectRooms()

	// solver.MakeFarm()

}
