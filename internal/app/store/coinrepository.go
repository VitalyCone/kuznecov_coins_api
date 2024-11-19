package store

import (
	"time"

	"github.com/VitalyCone/kuznecov_coins_api/internal/app/model"
)

type CoinRepository struct {
	store *Store
}

func (r *CoinRepository) Create(m *model.Coin) error {
	if err := r.store.db.QueryRow(
		"INSERT INTO user_coins (username, coins) "+
		"VALUES ($1, $2) RETURNING id, last_update",
	m.Username, m.Coins).Scan(&m.ID, &m.LastUpdate); err != nil{
		return err
	}

	return nil
}

func (r *CoinRepository) FindByUsername(username string) (model.Coin, error) {
	var m model.Coin
	if err := r.store.db.QueryRow(
		"SELECT id, coins, last_update FROM user_coins WHERE username = $1",
	username).Scan(&m.ID, &m.Coins, &m.LastUpdate); err != nil{
		return m,err
	}

	m.Username = username

	return m,nil
}

func (r *CoinRepository) FindById(id int) (model.Coin, error) {
	var m model.Coin
	if err := r.store.db.QueryRow(
		"SELECT username, coins, last_update FROM user_coins WHERE id = $1",
	id).Scan(&m.Username, &m.Coins, &m.LastUpdate); err != nil{
		return m,err
	}

	m.ID = id

	return m,nil
}

func (r *CoinRepository) DeleteById(id int) error {
	if _, err := r.store.db.Query(
		"DELETE FROM user_coins WHERE id = $1",
	id); err!= nil{
		return err
	}

	return nil
}
func (r *CoinRepository) DeleteByUsername(username string) error {
	if _, err := r.store.db.Query(
		"DELETE FROM user_coins WHERE username = $1",
	username); err!= nil{
		return err
	}

	return nil
}

func (r *CoinRepository) UpdateCoinsByUsername(username string, coins int) (model.Coin, error) {
	var model model.Coin
	timeNow := time.Now()
	if err := r.store.db.QueryRow(
		"UPDATE user_coins SET coins = $1, last_update = $2 WHERE username = $3 "+
		"RETURNING id",
		coins, timeNow, username).Scan(&model.ID); err != nil {
		return model, err
	}
	model.Username = username
	model.Coins = coins
	model.LastUpdate = timeNow

	return model, nil
}

func (r *CoinRepository) UpdateCoinsById(id int, coins int) error {
	if _, err := r.store.db.Exec(
		"UPDATE user_coins SET coins = $1, last_update = $2 WHERE id = $3",
		coins, time.Now(), id); err != nil {
		return err
	}
	return nil
}