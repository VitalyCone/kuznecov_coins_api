package dto

import (
	"time"

	"github.com/VitalyCone/kuznecov_coins_api/internal/app/model"
)

type CoinDetailsDto struct {
	Coins      int       `json:"coins"`
	LastUpdate time.Time `json:"last_update"`
}

func CoinModelToCoinDetailsDto(model model.Coin) CoinDetailsDto {
	return CoinDetailsDto{
		Coins:      model.Coins,
		LastUpdate: model.LastUpdate,
	}
}
