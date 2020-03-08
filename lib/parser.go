package solver

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// ReadFile reads .txt file from argument and puts it into string
func ReadFile() string {
	args := os.Args[1:]
	length := 0
	for i := range args {
		length = i + 1
	}

	if length == 0 {
		SetError("File name missing")
		return ""
	}
	if length != 1 {
		SetError("Too many arguments")
		return ""
	}

	fileName := args[0]

	file, err := os.Open(fileName)
	if err != nil {
		SetError(err.Error())
		return ""
	}

	dataByteArr, err := ioutil.ReadAll(file)
	if err != nil {
		SetError(err.Error())
		return ""
	}

	dataStr := string(dataByteArr)
	return dataStr

}

//ParseDataFromFile parsies data from string obtained from .txt file into Structs of Rooms and Links
func ParseDataFromFile(dataStr string) {

	length := len(dataStr)

	i := 0
	temp := ""

	for dataStr[i] != '\n' {
		temp += string(dataStr[i])
		i++
	}
	i++

	num, err := strconv.Atoi(temp)
	if err != nil {
		SetError("Invalid data format")
		return
	}
	if num <= 0 {
		SetError("Invalid number of ants")
		return
	}
	graph.ants = num

	temp = ""

	spaceCount := 0
	var name string
	var a []string
	marker := false
	start := false
	end := false
	for i < length {
		if dataStr[i] == '#' || dataStr[i] == 'L' {
			if dataStr[i+1] == '#' {
				marker = true
				i = i + 2
			} else {
				for i < length && dataStr[i] != '\n' {
					i++
				}
				i++
			}
		} else {
			for i < length && dataStr[i] != '\n' {
				temp += string(dataStr[i])
				if dataStr[i] == ' ' {
					spaceCount++
				}
				i++
			}
			i++
			if marker {
				if temp == "start" {
					start = true
				} else if temp == "end" {
					end = true
				}
				temp = ""
				marker = false
				continue
			}
			if spaceCount == 0 {
				a = strings.Split(temp, "-")
				if len(a) != 2 {
					SetError("Invalid data format")
					return
				}
				graph.links = append(graph.links, Link{a[0], a[1]})
			} else if spaceCount == 2 {
				a = strings.Split(temp, " ")
				name = a[0]
				xInt, err1 := strconv.Atoi(a[1])
				yInt, err2 := strconv.Atoi(a[2])
				if err1 != nil {
					SetError(err1.Error())
					return
				}
				if err2 != nil {
					SetError(err2.Error())
					return
				}
				room := &Room{name: name, x: xInt, y: yInt}
				if start {
					graph.start = *room
					start = false
				} else if end {
					graph.end = *room
					end = false
				}
				graph.rooms[name] = room
			} else {
				SetError("Invalid data format")
				return
			}
			temp = ""
			spaceCount = 0
		}

	}

}
