package main

import(
	"fmt"
	"net/http"
	"github.com/emmaunel/MagicGo/doors"
)

func main() {
	//Call the good function
	doors.goodguy()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world!")
	})
	http.ListenAndServe("127.0.0.1:8080")
}