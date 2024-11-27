package store

import "github.com/VitalyCone/account/internal/app/model"

type UsersRepository struct {
	store *Store
}

func (r *UsersRepository) CreateUser(m *model.User) error {
	if err := r.store.db.QueryRow(
		"INSERT INTO users (username, password_hash, first_name, second_name) "+
			"VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at",
		m.Username, m.PasswordHash, m.FirstName, m.SecondName).Scan(
		&m.ID, &m.CreatedAt, &m.UpdatedAt); err != nil {
		return err
	}

	return nil
}

func (r *UsersRepository) FindUserByUsername(username string) (model.User, error) {
	var m model.User

	if err := r.store.db.QueryRow(
		"SELECT * FROM users WHERE username = $1",
		username).Scan(
		&m.ID, &m.Username, &m.PasswordHash, &m.FirstName, &m.SecondName, &m.CreatedAt, &m.UpdatedAt, &m.Avatar); err != nil {
		return model.User{}, err
	}
	return m, nil
}

func (r *UsersRepository) DeleteUserByUsername(username string) error {
	if _, err := r.store.db.Exec(
		"DELETE FROM users WHERE username = $1",
		username); err != nil {
		return err
	}

	return nil
}

func (r *UsersRepository) ModifyUser(oldUsername string, m *model.User) error {
	if _, err := r.store.db.Exec(
		"UPDATE users "+
			"SET username = $2, password_hash = $3, first_name = $4, second_name = $5, avatar = $6 "+
			"WHERE username = $1",
		oldUsername, m.Username, m.PasswordHash, m.FirstName, m.SecondName, m.Avatar); err != nil {
		return err
	}
	return nil
}
