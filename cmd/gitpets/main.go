package main

import (
	"log"
	"net/http"

	"github.com/cassiusfive/gitpets/internal/card"
	"github.com/cassiusfive/gitpets/internal/pet"
)

func main() {
	http.Handle("/api", http.HandlerFunc(api))
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func api(w http.ResponseWriter, req *http.Request) {
	username := req.URL.Query().Get("username")
	petname := req.URL.Query().Get("petname")
	species := req.URL.Query().Get("species")
	theme := req.URL.Query().Get("theme")
	if username == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Missing param: username"))
		return
	}
	if petname == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing param: petname"))
		return
	}
	if species == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing param: species"))
		return
	}
	styles := card.CardStyles{}
	if theme == "light" {
		styles.Text = "black"
	} else {
		styles.Text = "white"
	}
	pet, err := pet.Create(username, petname, species)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	card.Generate(w, pet, styles)
}
