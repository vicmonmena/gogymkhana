package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type Error struct {
	Message string `json:"error"`
}

type IndexedWord struct {
	Word  string `json:"word"`
	Index int    `json:"index"`
}

type MessyText []IndexedWord

func route(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL Path: ", r.URL)
	if strings.Contains(strings.ToLower(r.URL.Path), "baunatal") == true {
		// Muchos de nosotros somos más capaces que algunos de nosotros
		messyText := MessyText{
			IndexedWord{"más", 1},
			IndexedWord{"de", 2},
			IndexedWord{"somos", 3},
			IndexedWord{"nosotros", 4},
			IndexedWord{"Muchos", 5},
			IndexedWord{"capaces", 6},
			IndexedWord{"nosotros", 7},
			IndexedWord{"algunos", 8},
			IndexedWord{"de", 9},
			IndexedWord{"que", 10},
		}
		json.NewEncoder(w).Encode(messyText)
	} else if strings.Contains(strings.ToLower(r.URL.Path), "sevilla") == true {
		// Pero ninguno de nosotros somos tan capaces como todos nosotros
		messyText := MessyText{
			IndexedWord{"somos", 1},
			IndexedWord{"de", 2},
			IndexedWord{"como", 3},
			IndexedWord{"nosotros", 4},
			IndexedWord{"Pero", 5},
			IndexedWord{"ninguno", 6},
			IndexedWord{"capaces", 7},
			IndexedWord{"nosotros", 8},
			IndexedWord{"todos", 9},
			IndexedWord{"tan", 10},
		}
		json.NewEncoder(w).Encode(messyText)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/baunatal", route)
	http.HandleFunc("/sevilla", route)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
