package main

import (
	"log"
	"net/http"

	svg "github.com/ajstarks/svgo"
)

func main() {
	http.Handle("/api", http.HandlerFunc(pet))
	err := http.ListenAndServe(":3030", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func pet(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(500, 500)
	s.Circle(250, 250, 125, "fill:none;stroke:black")
	s.End()
}
