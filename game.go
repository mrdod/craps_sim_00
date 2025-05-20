package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Bet struct {
	Type     string  `json:"type"`
	Amount   float64 `json:"amount"`
	Odds     float64 `json:"odds"`
	IsActive bool    `json:"isActive"`
}

type GameState struct {
	Point        int     `json:"point"`
	IsPointSet   bool    `json:"isPointSet"`
	IsGameActive bool    `json:"isGameActive"`
	LastRoll     int     `json:"lastRoll"`
	Message      string  `json:"message"`
	Bankroll     float64 `json:"bankroll"`
	Bets         []Bet   `json:"bets"`
	Die1         int     `json:"die1"`
	Die2         int     `json:"die2"`
}

func NewGame() *GameState {
	return &GameState{
		IsGameActive: true,
		Message:      "Place your bet and roll the dice!",
		Bankroll:     1000.00,
		Bets:         make([]Bet, 0),
	}
}

func (g *GameState) PlaceBet(betType string, amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("bet amount must be greater than 0")
	}
	if amount > g.Bankroll {
		return fmt.Errorf("insufficient funds")
	}

	var odds float64
	switch betType {
	case "pass":
		odds = 1.0
	case "dont_pass":
		odds = 1.0
	case "come":
		odds = 1.0
	case "dont_come":
		odds = 1.0
	case "field":
		odds = 1.0
	case "any_seven":
		odds = 4.0
	case "any_craps":
		odds = 7.0
	case "hard_4":
		odds = 7.0
	case "hard_6":
		odds = 9.0
	case "hard_8":
		odds = 9.0
	case "hard_10":
		odds = 7.0
	default:
		return fmt.Errorf("invalid bet type")
	}

	g.Bankroll -= amount
	g.Bets = append(g.Bets, Bet{
		Type:     betType,
		Amount:   amount,
		Odds:     odds,
		IsActive: true,
	})
	return nil
}

func (g *GameState) RollDice() {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// Roll two dice
	g.Die1 = rand.Intn(6) + 1
	g.Die2 = rand.Intn(6) + 1
	g.LastRoll = g.Die1 + g.Die2

	// Process bets
	g.processBets()

	if !g.IsPointSet {
		// First roll
		switch g.LastRoll {
		case 7, 11:
			g.Message = "Natural! You win!"
			g.IsGameActive = false
		case 2, 3, 12:
			g.Message = "Craps! You lose!"
			g.IsGameActive = false
		default:
			g.Point = g.LastRoll
			g.IsPointSet = true
			g.Message = fmt.Sprintf("Point is set to %d. Roll again!", g.Point)
		}
	} else {
		// Subsequent rolls
		if g.LastRoll == g.Point {
			g.Message = "You hit the point! You win!"
			g.IsGameActive = false
		} else if g.LastRoll == 7 {
			g.Message = "Seven out! You lose!"
			g.IsGameActive = false
		} else {
			g.Message = "Roll again!"
		}
	}
}

func (g *GameState) processBets() {
	for i := range g.Bets {
		if !g.Bets[i].IsActive {
			continue
		}

		switch g.Bets[i].Type {
		case "pass":
			if !g.IsPointSet {
				if g.LastRoll == 7 || g.LastRoll == 11 {
					g.Bankroll += g.Bets[i].Amount * (1 + g.Bets[i].Odds)
					g.Bets[i].IsActive = false
				} else if g.LastRoll == 2 || g.LastRoll == 3 || g.LastRoll == 12 {
					g.Bets[i].IsActive = false
				}
			} else if g.LastRoll == g.Point {
				g.Bankroll += g.Bets[i].Amount * (1 + g.Bets[i].Odds)
				g.Bets[i].IsActive = false
			} else if g.LastRoll == 7 {
				g.Bets[i].IsActive = false
			}
		case "dont_pass":
			if !g.IsPointSet {
				if g.LastRoll == 2 || g.LastRoll == 3 {
					g.Bankroll += g.Bets[i].Amount * (1 + g.Bets[i].Odds)
					g.Bets[i].IsActive = false
				} else if g.LastRoll == 7 || g.LastRoll == 11 {
					g.Bets[i].IsActive = false
				}
			} else if g.LastRoll == 7 {
				g.Bankroll += g.Bets[i].Amount * (1 + g.Bets[i].Odds)
				g.Bets[i].IsActive = false
			} else if g.LastRoll == g.Point {
				g.Bets[i].IsActive = false
			}
		case "field":
			if g.LastRoll == 2 {
				g.Bankroll += g.Bets[i].Amount*2 + g.Bets[i].Amount
				g.Bets[i].IsActive = false
			} else if g.LastRoll == 12 {
				g.Bankroll += g.Bets[i].Amount*3 + g.Bets[i].Amount
				g.Bets[i].IsActive = false
			} else if g.LastRoll == 3 || g.LastRoll == 4 || g.LastRoll == 9 || g.LastRoll == 10 || g.LastRoll == 11 {
				g.Bankroll += g.Bets[i].Amount + g.Bets[i].Amount
				g.Bets[i].IsActive = false
			} else {
				g.Bets[i].IsActive = false
			}
		case "any_seven":
			if g.LastRoll == 7 {
				g.Bankroll += g.Bets[i].Amount * g.Bets[i].Odds
			}
			g.Bets[i].IsActive = false
		case "any_craps":
			if g.LastRoll == 2 || g.LastRoll == 3 || g.LastRoll == 12 {
				g.Bankroll += g.Bets[i].Amount * g.Bets[i].Odds
			}
			g.Bets[i].IsActive = false
		case "hard_4":
			if g.LastRoll == 4 && g.Die1 == g.Die2 {
				g.Bankroll += g.Bets[i].Amount * g.Bets[i].Odds
			}
			g.Bets[i].IsActive = false
		case "hard_6":
			if g.LastRoll == 6 && g.Die1 == g.Die2 {
				g.Bankroll += g.Bets[i].Amount * g.Bets[i].Odds
			}
			g.Bets[i].IsActive = false
		case "hard_8":
			if g.LastRoll == 8 && g.Die1 == g.Die2 {
				g.Bankroll += g.Bets[i].Amount * g.Bets[i].Odds
			}
			g.Bets[i].IsActive = false
		case "hard_10":
			if g.LastRoll == 10 && g.Die1 == g.Die2 {
				g.Bankroll += g.Bets[i].Amount * g.Bets[i].Odds
			}
			g.Bets[i].IsActive = false
		}
	}
}

func (g *GameState) Reset() {
	g.Point = 0
	g.IsPointSet = false
	g.IsGameActive = true
	g.LastRoll = 0
	g.Message = "Place your bet and roll the dice!"
	g.Bets = make([]Bet, 0)
	g.Die1 = 0
	g.Die2 = 0
}
