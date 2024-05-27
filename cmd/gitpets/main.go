package main

import (
	"log"
	"net/http"

	"github.com/cassiusfive/gitpets/internal/card"
	"github.com/cassiusfive/gitpets/internal/pet"
)

func main() {
	http.Handle("/api", http.HandlerFunc(api))
	err := http.ListenAndServe(":3030", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func api(w http.ResponseWriter, req *http.Request) {
	username := req.URL.Query().Get("username")
	petname := req.URL.Query().Get("petname")
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
	pet, err := pet.Create(username, petname)
	if err != nil {
		return
	}
	card.Generate(w, pet, card.CardStyles{})
}
