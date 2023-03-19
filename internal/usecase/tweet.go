package usecase

import (
	"context"

	"go-crud-app/internal/dto"
	"go-crud-app/internal/entity"
)

type TweetRepo interface {
	Create(ctx context.Context, dto dto.CreateTweetDTO) (*entity.Tweet, error)
	GetByID(ctx context.Context, id int64) (*dto.GetTweetByIdDTO, error)
	GetAuthorID(ctx context.Context, id int64) (*dto.GetTweetAuthorIdDTO, error)
	ListAllTweets(ctx context.Context, params dto.ListAllTweetsParams) ([]*dto.GetTweetByIdDTO, error)
	DeleteByID(ctx context.Context, id int64) error
}

type TweetUseCase struct {
	repo TweetRepo
}

func NewTweetUseCase(r TweetRepo) *TweetUseCase {
	return &TweetUseCase{r}
}

func (uc *TweetUseCase) Create(ctx context.Context, dto dto.CreateTweetDTO) (*entity.Tweet, error) {
	tweet, err := uc.repo.Create(ctx, dto)
	if err != nil {
		return nil, err
	}
	return tweet, nil
}

func (uc *TweetUseCase) GetByID(ctx context.Context, id int64) (*dto.GetTweetByIdDTO, error) {
	tweet, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return tweet, nil
}

func (uc *TweetUseCase) GetAuthorID(ctx context.Context, id int64) (*dto.GetTweetAuthorIdDTO, error) {
	dto, err := uc.repo.GetAuthorID(ctx, id)
	if err != nil {
		return nil, err
	}
	return dto, nil
}

func (uc *TweetUseCase) ListAllTweets(ctx context.Context, dto dto.ListAllTweetsParams) ([]*dto.GetTweetByIdDTO, error) {
	tweets, err := uc.repo.ListAllTweets(ctx, dto)
	if err != nil {
		return nil, err
	}
	return tweets, nil
}

func (uc *TweetUseCase) DeleteByID(ctx context.Context, id int64) error {
	err := uc.repo.DeleteByID(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
