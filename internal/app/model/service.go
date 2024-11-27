package model

import "time"

type Service struct {
	ID          int
	CompanyID   int // тут будет id из company
	ServiceType ServiceType //тут я буду получать id типа из таблицы service_types
	Text        string
	Price       float32
	//Reviews     []Review
	//Tags        []Tag
	CreatedAt   time.Time
	UpdatedAt	time.Time
}
