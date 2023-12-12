package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/golang-insiders/site/internal/types"
)

type TalkService struct {
	DB *sql.DB
}

func (tm TalkService) Insert(ctx context.Context, t *types.Talk) error {
	query := `
		INSERT INTO talks (twitter_username, title, summary, timezone)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return tm.DB.QueryRowContext(ctx, query,
		t.TwitterUsername, t.Title, t.Summary, t.Timezone,
	).Scan(&t.ID, &t.CreatedAt)
}

func (tm TalkService) GetByID(ctx context.Context, id int) (*types.Talk, error) {
	query := `
		SELECT id, twitter_username, title, summary, timezone
		FROM talks
		WHERE id=$1;
	`
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var t types.Talk

	err := tm.DB.QueryRowContext(ctx, query, id).Scan(
		&t.ID,
		&t.TwitterUsername,
		&t.Title,
		&t.Summary,
		&t.Timezone,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &t, nil
}
