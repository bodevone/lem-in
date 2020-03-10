package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	visualizer "./lib"
)

func main() {
	bytes, _ := ioutil.ReadAll(os.Stdin)
	str := string(bytes)
	fmt.Println(str)

	visualizer.InitGraph()

	if !visualizer.ParseDataFromPipe(str) {
		fmt.Println("ERROR: No visualisation today!")
		return
	}

	http.HandleFunc("/", visualizer.WriteHandler)
	http.HandleFunc("/data", visualizer.SendData)

	http.Handle("/cmd/visualizer/public/", http.StripPrefix("/cmd/visualizer/public/", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":8080", nil)

}
