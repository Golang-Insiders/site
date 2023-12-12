package types

import (
	"errors"
	"time"
)

var (
	// Placeholder until check for valid twitter user function is implemented
	// TODO: implement twitter username check (api? scraping?)
	INVALID_TWITTER_USER = "dne"
	MAX_TITLE_LEN        = 30
	MAX_SUMMARY_LEN      = 300

	// Declare errors for talk validation
	ErrTwitterUsernameDNE = errors.New("twitter username must exist")
	ErrTitleLength        = errors.New("title must be at least 30 characters long")
	ErrSummaryLength      = errors.New("summary must be at least 300 characters long")
	ErrTimeZone           = errors.New("timezone must be provided")
)

type Talk struct {
	TwitterUsername string
	Title           string
	Summary         string
	Timezone        string
	ID              int64
	CreatedAt       time.Time
}

func ValidateTalk(t Talk) []error {
	var errors []error
	if t.TwitterUsername == INVALID_TWITTER_USER {
		errors = append(errors, ErrTwitterUsernameDNE)
	}
	if len(t.Title) < MAX_TITLE_LEN {
		errors = append(errors, ErrTitleLength)
	}
	if len(t.Summary) < MAX_SUMMARY_LEN {
		errors = append(errors, ErrSummaryLength)
	}
	if t.Timezone == "" {
		errors = append(errors, ErrTimeZone)
	}
	return errors
}
