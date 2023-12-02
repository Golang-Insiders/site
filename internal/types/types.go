package types

import "time"

var (
	// Placeholder until check for valid twitter user function is implemented
	INVALID_TWITTER_USER = "dne"
	MAX_TITLE_LEN        = 30
	MAX_SUMMARY_LEN      = 300
)

type Talk struct {
	TwitterUsername string
	Title           string
	Summary         string
	Timezone        string
	ID              int64
	CreatedAt       time.Time
}

func (tf *Talk) ValidateTalk() (bool, string) {
	if tf.TwitterUsername == INVALID_TWITTER_USER {
		return false, "Twitter username must exist"
	}
	if len(tf.Title) < MAX_TITLE_LEN {
		return false, "Title must be at least 30 characters long"
	}
	if len(tf.Summary) < MAX_SUMMARY_LEN {
		return false, "Summary must be at least 300 characters long"
	}
	return true, ""
}
