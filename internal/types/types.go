package types

import "time"

type Talk struct {
	TwitterUsername string
	Title           string
	Summary         string
	Timezone        string
	ID              int64
	CreatedAt       time.Time
}

func (tf *Talk) ValidateTalk() (bool, string) {
	if tf.TwitterUsername == "dne" {
		return false, "Twitter username must exist"
	}
	if len(tf.Title) < 30 {
		return false, "Title must be at least 30 characters long"
	}
	if len(tf.Summary) < 300 {
		return false, "Summary must be at least 300 characters long"
	}
	return true, ""
}
