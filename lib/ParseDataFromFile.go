package solver

import (
	"fmt"
	"strconv"
	"strings"
)

//ParseDataFromFile parsies data from string obtained from .txt file into Structs of Rooms and Links
func ParseDataFromFile(dataStr string) {

	length := len(dataStr)

	i := 0
	temp := ""
	spaceCount := 0
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
				links = append(links, Link{room1: a[0], room2: a[1]})
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
				rooms[name] = Room{xInt, yInt}
			}
			temp = ""
			spaceCount = 0
		}

	}

	fmt.Println(rooms)
	fmt.Println(links)

}
