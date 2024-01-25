package main

import (
	"fmt"
	"os"
	"testing"
)

func TestGetComputerDraw(t *testing.T) {
	draws := make([]int, len(RPS))

	for i := 0; i < len(RPS)*1000; i++ {
		if draw := getComputerDraw(); draw < 1 || draw > len(RPS) {
			t.Errorf("Excpected computer to draw number between 1 and %d, but recieved %d.", len(RPS), draw)
		} else {
			draws[draw-1]++
		}
	}

	for i, d := range draws {
		if d == 0 {
			t.Errorf("Excpected that computer would succeed to draw the number '%d' (%s) out of %d attempts.", d, RPS[i+1], len(RPS)*1000)
		} else {
			fmt.Printf("Computer drew the number '%d' (%s) %d times (%.1f%%)\n", i+1, RPS[i+1], d, float64(d)/float64(len(RPS)*1000)*100)
		}
	}
}

func TestDidPlayerWon(t *testing.T) {
	if didPlayerWon(1, 2) {
		t.Errorf("Expected Player to lose when player draw 1 (Rock) and computer draw 2 (Paper).")
	}

	if didPlayerWon(1, 5) {
		t.Errorf("Expected Player to lose when player draw 1 (Rock) and computer draw 5 (Spock).")
	}

	if didPlayerWon(2, 3) {
		t.Errorf("Expected Player to lose when player draw 2 (Paper) and computer draw 3 (Scissors).")
	}

	if didPlayerWon(2, 4) {
		t.Errorf("Expected Player to lose when player draw 2 (Paper) and computer draw 4 (Lizard).")
	}

	if didPlayerWon(3, 1) {
		t.Errorf("Expected Player to lose when player draw 3 (Scissors) and computer draw 1 (Rock).")
	}

	if didPlayerWon(3, 5) {
		t.Errorf("Expected Player to lose when player draw 3 (Scissors) and computer draw 5 (Spock).")
	}

	if didPlayerWon(4, 1) {
		t.Errorf("Expected Player to lose when player draw 4 (Lizard) and computer draw 1 (Rock).")
	}

	if didPlayerWon(4, 3) {
		t.Errorf("Expected Player to lose when player draw 4 (Lizard) and computer draw 3 (Scissors).")
	}

	if didPlayerWon(5, 2) {
		t.Errorf("Expected Player to lose when player draw 5 (Spock) and computer draw 2 (Paper).")
	}

	if didPlayerWon(5, 4) {
		t.Error("Expected Player to lose when player draw 5 (Spock) and computer draw 4 (Lizard).")
	}
}

func TestSaveToFileAndLoadFromFile(t *testing.T) {
	defer func() {
		os.Remove("_statstesting.txt")
	}()

	os.Remove("_statstesting.txt")
	var pSave = Player{
		Stats: stats{123, 456, 789},
	}
	var pLoad = Player{
		Stats: stats{0, 0, 0},
	}

	pSave.SaveToFile("_statstesting.txt")
	pLoad.LoadFromFile("_statstesting.txt")

	if pSave.Stats.won != pLoad.Stats.won || pSave.Stats.lost != pLoad.Stats.lost || pSave.Stats.draw != pLoad.Stats.draw {
		t.Error("Expected to load the same stats that were saved to file.")
	}
	os.Remove("_statstesting.txt")
}
