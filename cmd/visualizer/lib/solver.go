package visualizer

import (
	"encoding/json"
	"net/http"
	"os"
	"text/template"
)

// Graph connects all together
type Graph struct {
	Rooms      map[string]Room
	Links      []Link
	Start      Room
	End        Room
	Ants       int
	Iterations int
	Steps      []map[int]string
}

// Room represents room with x and y values
type Room struct {
	Name string
	X, Y int
}

// Link represents links between two rooms
type Link struct {
	Room1 string
	Room2 string
}

var graph Graph

// InitGraph to make map filed in Graph
func InitGraph() {
	graph.Rooms = make(map[string]Room)
	// graph.steps = make(map[int]string)

}

// WriteHandler to
func WriteHandler(w http.ResponseWriter, r *http.Request) {

	status := http.StatusNotFound

	if r.URL.Path != "/" {
		w.WriteHeader(status)
		if status == http.StatusNotFound {
			http.Error(w, "Error: 404 - Page not found.Please refresh with the valid URL.", http.StatusNotFound)
		}

		return
	}

	f, err := os.Open("cmd/visualizer/templates/index.html")

	if err != nil {
		http.Error(w, "Internal server error - 500.", http.StatusInternalServerError)
		return
	}

	f.Close()
	tmpl := template.Must(template.ParseFiles("cmd/visualizer/templates/index.html"))
	// tmpl := template.Must(template.New("").Parse(templ))

	err = tmpl.Execute(w, graph)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// SendData to send links
func SendData(rw http.ResponseWriter, req *http.Request) {

	json.NewEncoder(rw).Encode(graph)

}
