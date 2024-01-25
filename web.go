package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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
	playerChoice, err := strconv.Atoi(r.URL.Query().Get("c"))

	if err != nil || playerChoice < 1 || playerChoice > 5 {
		log.Println("Expected integer paramater between 1 to 5 in URL, but recieved c=", r.URL.Query().Get("c"))
	} else {
		result := PlayRound(&p, playerChoice)
		out, err := json.MarshalIndent(result, "", "   ")
		if err != nil {
			log.Println(err)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write(out)
		}
	}
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
	log.Println("Reseting stats.")
}
