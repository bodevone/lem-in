package visualizer

import (
	"strconv"
	"strings"
)

//ParseDataFromPipe parses data from string obtained from terminal pipe into Structs of Rooms and Links
func ParseDataFromPipe(dataStr string) bool {

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
		return false
	}
	temp = ""
	graph.Ants = num

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
		} else if dataStr[i] == '\n' {
			i++
			for i < length-1 {
				for i < length-1 && dataStr[i] != '\n' {
					temp += string(dataStr[i])
					i++
				}
				arr := strings.Split(temp, " ")
				steps := make(map[int]string)
				for _, a := range arr {
					arrNew := strings.Split(a[1:], "-")
					ant, room := arrNew[0], arrNew[1]
					antNum, _ := strconv.Atoi(ant)
					steps[antNum] = room
				}
				temp = ""
				i++
				graph.Steps = append(graph.Steps, steps)
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
				graph.Links = append(graph.Links, Link{a[0], a[1]})
			} else if spaceCount == 2 {
				a = strings.Split(temp, " ")
				name = a[0]
				xInt, _ := strconv.Atoi(a[1])
				yInt, _ := strconv.Atoi(a[2])
				room := Room{name, xInt, yInt}
				if start {
					graph.Start = room
					start = false
				} else if end {
					graph.End = room
					end = false
				}
				graph.Rooms[name] = room
			}
			temp = ""
			spaceCount = 0

		}

	}

	return true

}
