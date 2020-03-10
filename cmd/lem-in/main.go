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

	solver.CheckError()
	occured, message = solver.GetError()
	if occured {
		fmt.Println("ERROR: " + message)
		return
	}

	solver.AddNeighbours()
	occured, message = solver.GetError()
	if occured {
		fmt.Println("ERROR: " + message)
		return
	}

	solver.PrintAll()

	if solver.StartEndConnected() {
		solver.PrintStartEnd()
		return
	}

	solver.FindPaths()
	if occured {
		fmt.Println("ERROR: " + message)
		return
	}

	solver.FindPathsCombn()

	solver.FindSolution()

	solver.GetSolution()

	// solver.ConnectRooms()
	// solver.MakeFarm()

}
