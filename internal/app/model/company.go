package model

import "time"

type Company struct {
	ID          int
	Avatar      []byte
	Name        string
	//Services    []Service
	Description string
	Tags        []Tag
	//Member      []User
	//Moderators  []User
	Reviews     []Review
	CreatedAt   time.Time
	UpdatedAt	time.Time
}