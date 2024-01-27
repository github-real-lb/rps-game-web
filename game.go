package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// RPS is a map of game's draw options and their int values.
var RPS = map[int]string{
	1: "Rock",
	2: "Paper",
	3: "Scissors",
	4: "Lizard",
	5: "Spock",
}

// stats is a struct to record player's stat during the game.
type Stats struct {
	Won  int `json:"stats_won"`
	Lost int `json:"stats_lost"`
	Draw int `json:"stats_draw"`
}

type Player struct {
	Name  string
	Stats Stats
}

type Round struct {
	PlayerChoice   string `json:"player_choice"`
	ComputerChoice string `json:"computer_choice"`
	Result         string `json:"result"`
}

// PlayRound receives player's draw (between 1 to 5)
// It returns Round type of round results.
func PlayRound(p *Player, playerDraw int) Round {
	computerDraw := getComputerDraw()

	var r Round
	r.PlayerChoice = fmt.Sprintf("You chose %v.", RPS[playerDraw])
	r.ComputerChoice = fmt.Sprintf("Computer chose %v.", RPS[computerDraw])

	if playerDraw == computerDraw {
		p.Stats.Draw++
		r.Result = "It's a draw!"
	} else if didPlayerWon(playerDraw, computerDraw) {
		p.Stats.Won++
		r.Result = "You've won!"
	} else {
		p.Stats.Lost++
		r.Result = "You've lost!"
	}

	return r
}

// getComputerDraw gets computer random draw and returns it's value.
func getComputerDraw() int {
	return rand.Intn(len(RPS)*20)%len(RPS) + 1
}

// didPlayerWon receives player's and computer's draw and returns true if player has won.
func didPlayerWon(playerDraw, computerDraw int) bool {
	if playerDraw == computerDraw {
		return false
	} else {
		switch playerDraw {
		case 1:
			return computerDraw != 2 && computerDraw != 5
		case 2:
			return computerDraw != 3 && computerDraw != 4
		case 3:
			return computerDraw != 5 && computerDraw != 1
		case 4:
			return computerDraw != 3 && computerDraw != 1
		case 5:
			return computerDraw != 4 && computerDraw != 2
		default:
			return false
		}
	}
}

// ResetStats reset player's stats to 0.
func (p *Player) ResetStats() {
	p.Stats = Stats{0, 0, 0}
}

// SaveToFile save the players stats to a file.
func (p *Player) SaveToFile(fileName string) error {
	str := fmt.Sprintf("%d,%d,%d", p.Stats.Won, p.Stats.Lost, p.Stats.Draw)
	err := os.WriteFile(fileName, []byte(str), 0666)
	return err
}

// LoadFromFile loads the players stats from a file.
func (p *Player) LoadFromFile(fileName string) error {
	p.Stats = Stats{0, 0, 0}
	bs, err := os.ReadFile(fileName)

	if err != nil {
		return err
	}

	ss := strings.Split(string(bs), ",")

	if len(ss) != 3 {
		return errors.New("file isn't in the correct format")
	}

	if p.Stats.Won, err = strconv.Atoi(ss[0]); err != nil {
		return errors.New("file isn't in the correct format")
	}

	if p.Stats.Lost, err = strconv.Atoi(ss[1]); err != nil {
		return errors.New("file isn't in the correct format")
	}

	if p.Stats.Draw, err = strconv.Atoi(ss[2]); err != nil {
		return errors.New("file isn't in the correct format")
	}

	return nil
}
