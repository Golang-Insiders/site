package data

import (
	"database/sql"
	"errors"
	"fmt"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Services struct {
	Talks    TalkService
	TimeZone TimeZoneService
}

func NewServices(db *sql.DB) (Services, error) {
	if db == nil {
		return Services{}, fmt.Errorf("DB does not exist")
	}
	return Services{
		Talks:    TalkService{DB: db},
		TimeZone: TimeZoneService{},
	}, nil
}
