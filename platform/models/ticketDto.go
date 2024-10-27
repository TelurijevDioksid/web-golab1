package models

import "time"

type TicketDto struct {
	Oib       string    `json:"oib"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewTicketDto(oib, fn, ln string, cat time.Time) *TicketDto {
	return &TicketDto{
		Oib:       oib,
		FirstName: fn,
		LastName:  ln,
		CreatedAt: cat,
	}
}
