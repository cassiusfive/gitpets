package main

import (
	"fmt"
	"log"
	"net/http"

	svg "github.com/ajstarks/svgo"
	"github.com/cassiusfive/gitpets/internal/gitstats"
)

func main() {
	fmt.Println(gitstats.GetStats("cassiusfive"))
	http.Handle("/api", http.HandlerFunc(pet))
	err := http.ListenAndServe(":3030", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func pet(w http.ResponseWriter, req *http.Request) {
	// username := req.URL.Query().Get("username")
	// petname := req.URL.Query().Get("petname")
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(500, 500)
	s.Circle(250, 250, 125, "fill:none;stroke:black")
	s.End()
}
