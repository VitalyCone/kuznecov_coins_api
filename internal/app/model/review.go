package model

import "time"

type Review struct {
	ID          int
	ReviewType  ReviewType
	TypeID     int
	Rating      int
	CreatorUser User
	Header      string
	Text        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
