package types

import (
	"slices"
	"strings"
	"testing"
)

var (
	ValidTwitterUsername   = "valid-twitter-username"
	InvalidTwitterUsername = "dne"
	ValidTitle             = strings.Repeat("a", 30)
	InvalidTitle           = strings.Repeat("a", 29)
	ValidSummary           = strings.Repeat("a", 300)
	InvalidSummary         = strings.Repeat("a", 299)
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
	// if err != ErrTwitterUsernameDNE {
	if !slices.Contains(err, ErrTwitterUsernameDNE) {
		t.Error("expected twitter username error: ", err)
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
	// if err != ErrTitleLength {
	if !slices.Contains(err, ErrTitleLength) {
		t.Error("expected title length error: ", err)
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
	// if err != ErrSummaryLength {
	if !slices.Contains(err, ErrSummaryLength) {
		t.Error("expected summary length error: ", err)
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
