package main

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"
)

const statsFileName = "stats.txt"

var p Player

func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func play(w http.ResponseWriter, r *http.Request) {
	result := PlayRound(&p, 1)
	out, err := json.MarshalIndent(result, "", "   ")
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func save(w http.ResponseWriter, r *http.Request) {
	if err := p.SaveToFile(statsFileName); err == nil {
		log.Println("Saving stats to file: Success.")
	} else {
		log.Println("Saving stats to file: Failure.", err)
	}
}

func reset(w http.ResponseWriter, r *http.Request) {
	p.ResetStats()
}
