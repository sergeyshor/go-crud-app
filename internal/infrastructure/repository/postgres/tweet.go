package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"go-crud-app/internal/dto"
	"go-crud-app/internal/entity"
)

type TweetRepo struct {
	*sql.DB
}

func NewTweetRepo(db *sql.DB) *TweetRepo {
	return &TweetRepo{db}
}

func (tr *TweetRepo) Create(ctx context.Context, dto dto.CreateTweetDTO) (*entity.Tweet, error) {
	query := "INSERT INTO tweets(author_id, content) VALUES($1, $2) RETURNING id, author_id, content, created_at"
	row := tr.QueryRowContext(ctx, query, dto.AuthorID, dto.Content)

	tweet := &entity.Tweet{}

	err := row.Scan(&tweet.ID, &tweet.AuthorID, &tweet.Content, &tweet.CreatedAt)
	if err != nil {
		return nil, err
	}

	return tweet, nil
}

func (tr *TweetRepo) GetByID(ctx context.Context, id int64) (*dto.GetTweetByIdDTO, error) {
	query := "SELECT tweets.id AS id, users.name AS author_name, content, tweets.created_at AS created_at FROM tweets INNER JOIN users ON tweets.author_id = users.id WHERE tweets.id=$1"
	row := tr.QueryRowContext(ctx, query, id)

	tweet := &dto.GetTweetByIdDTO{}

	err := row.Scan(&tweet.ID, &tweet.AuthorName, &tweet.Content, &tweet.CreatedAt)
	if err != nil {
		return nil, err
	}
	return tweet, nil
}

func (tr *TweetRepo) GetAuthorID(ctx context.Context, id int64) (*dto.GetTweetAuthorIdDTO, error) {
	query := "SELECT author_id FROM tweets WHERE id=$1"
	row := tr.QueryRowContext(ctx, query, id)

	tweet := &dto.GetTweetAuthorIdDTO{}

	err := row.Scan(&tweet.AuthorID)
	if err != nil {
		return nil, err
	}
	return tweet, nil
}

func (tr *TweetRepo) ListAllTweets(ctx context.Context, params dto.ListAllTweetsParams) ([]*dto.GetTweetByIdDTO, error) {
	query := "SELECT tweets.id AS id, users.name AS author_name, content, tweets.created_at AS created_at FROM tweets INNER JOIN users ON tweets.author_id = users.id LIMIT $1 OFFSET $2"
	rows, err := tr.QueryContext(ctx, query, params.Limit, params.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tweets := []*dto.GetTweetByIdDTO{}

	for rows.Next() {
		tweet := &dto.GetTweetByIdDTO{}
		if err := rows.Scan(&tweet.ID, &tweet.AuthorName, &tweet.Content, &tweet.CreatedAt); err != nil {
			return nil, err
		}
		tweets = append(tweets, tweet)
	}
	return tweets, nil
}

func (tr *TweetRepo) DeleteByID(ctx context.Context, id int64) error {
	result, err := tr.ExecContext(ctx, "DELETE FROM tweets WHERE id=$1", id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return fmt.Errorf("expected to affect 1 row, affected %d", rows)
	}
	return nil
}
