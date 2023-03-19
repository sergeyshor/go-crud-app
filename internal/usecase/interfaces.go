package usecase

import (
	"context"

	"go-crud-app/internal/dto"
	"go-crud-app/internal/entity"
)

type Tweet interface {
	Create(ctx context.Context, dto dto.CreateTweetDTO) (*entity.Tweet, error)
	GetByID(ctx context.Context, id int64) (*dto.GetTweetByIdDTO, error)
	GetAuthorID(ctx context.Context, id int64) (*dto.GetTweetAuthorIdDTO, error)
	ListAllTweets(ctx context.Context, params dto.ListAllTweetsParams) ([]*dto.GetTweetByIdDTO, error)
	DeleteByID(ctx context.Context, id int64) error
}

type User interface {
	Create(ctx context.Context, dto dto.CreateUserDTO) error
	GetByID(ctx context.Context, id int64) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	DeleteByID(ctx context.Context, id int64) error
}
