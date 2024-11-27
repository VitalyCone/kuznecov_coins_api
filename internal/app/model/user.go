package model

import "time"

type User struct {
	ID           int    `json:"id"`
	Avatar       []byte `json:"avatar"`
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
	FirstName    string `json:"first_name"`
	SecondName   string `json:"second_name"`
	// MemberIn     []Company
	// ModeratorIn  []Company
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
