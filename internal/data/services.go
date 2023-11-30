package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Services struct {
	Talks    TalkService
	TimeZone TimeZoneService
}

func NewServices(db *sql.DB) Services {
	return Services{
		Talks:    TalkService{DB: db},
		TimeZone: TimeZoneService{},
	}
}
