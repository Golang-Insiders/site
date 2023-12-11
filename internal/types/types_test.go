package types

import (
	"strings"
	"testing"
)

// Generated a string of the provided length
func generateString(length int) string {
	if length <= 0 {
		return ""
	}

	pattern := "abc"
	result := strings.Repeat(pattern, (length/len(pattern))+1)[:length]

	return result
}

var (
	ValidTwitterUsername   = "valid-twitter-username"
	InvalidTwitterUsername = "dne"
	ValidTitle             = generateString(30)
	InvalidTitle           = generateString(29)
	ValidSummary           = generateString(300)
	InvalidSummary         = generateString(299)
)

func TestValidateTalk(t *testing.T) {

	// Test empty talk
	err := ValidateTalk(Talk{})
	if err == nil {
		t.Error("expected an error")
	}

	// Test invalid twitter username check
	err = ValidateTalk(Talk{
		TwitterUsername: InvalidTwitterUsername,
		Title:           ValidTitle,
		Summary:         ValidSummary,
		Timezone:        "UTC",
	})
	if err == nil {
		t.Error("expected an error: ", err)
	}
	if err != nil {
		if err != ErrTwitterUsernameDNE {
			t.Error("expected twitter username error: ", err)
		}
	}

	// Test invalid title check
	err = ValidateTalk(Talk{
		TwitterUsername: ValidTwitterUsername,
		Title:           InvalidTitle,
		Summary:         ValidSummary,
		Timezone:        "UTC",
	})
	if err == nil {
		t.Error("expected an error: ", err)
	}
	if err != nil {
		if err != ErrTitleLength {
			t.Error("expected title length error: ", err)
		}
	}

	// Test invalid summary check
	err = ValidateTalk(Talk{
		TwitterUsername: ValidTwitterUsername,
		Title:           ValidTitle,
		Summary:         InvalidSummary,
		Timezone:        "UTC",
	})
	if err == nil {
		t.Error("expected an error: ", err)
	}
	if err != nil {
		if err != ErrSummaryLength {
			t.Error("expected summary length error: ", err)
		}
	}

	// Test valid talk check
	err = ValidateTalk(Talk{
		TwitterUsername: ValidTwitterUsername,
		Title:           ValidTitle,
		Summary:         ValidSummary,
		Timezone:        "UTC",
	})
	if err != nil {
		t.Error("should have passed: ", err)
	}
}
