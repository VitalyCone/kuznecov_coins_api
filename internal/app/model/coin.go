package model

import "time"

type Coin struct {
	ID         int       `json:"id"`
	Username   string    `json:"username"`
	Coins      int       `json:"coins"`
	LastUpdate time.Time `json:"last_update"`
}
