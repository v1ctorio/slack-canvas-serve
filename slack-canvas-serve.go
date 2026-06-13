package main
import (
	"fmt"
	"os"
	"net/http"
	)


// TODO remove all these constants and load from cli flags/ a config file
const Port = 4555
const Authentication = false



func main() {
	mux := &http.ServeMux{}
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/canvas/{id}", canvasHandler)
	fmt.Println("Hello gang")

	http.ListenAndServe(":4555", mux)
	os.Exit(69)
}


func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from http. I'm here to server your canvas")
}
func canvasHandler(w http.ResponseWriter, r *http.Request) {
	canvas_id := r.PathValue("id")
	format := r.FormValue("format")

	if format == "" {
		format = "md"
	}

	fmt.Fprintf(w, "requested canvas: `%s` with format: `%s`", canvas_id, format)
}
