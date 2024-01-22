package main

import (
	"log"
	"net/http"
)

// var p Player declared in web.go

func main() {
	if err := p.LoadFromFile(statsFileName); err == nil {
		log.Println("Loading stats from file: Success.")
	} else {
		log.Println("Loading stats from file: Failure.", err)
	}

	defer func() {
		if err := p.SaveToFile(statsFileName); err == nil {
			log.Println("Saving stats to file: Success.")
		} else {
			log.Println("Saving stats to file: Failure.", err)
		}
	}()

	http.HandleFunc("/save", save)
	http.HandleFunc("/reset", reset)
	http.HandleFunc("/play", play)
	http.HandleFunc("/", homePage)

	log.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", nil)
}
