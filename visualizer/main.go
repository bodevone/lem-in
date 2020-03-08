package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	bytes, _ := ioutil.ReadAll(os.Stdin)
	str := string(bytes)
	fmt.Println(str)
}
