package usecase

import (
	"context"

	"go-crud-app/internal/dto"
	"go-crud-app/internal/entity"
)

type UserRepo interface {
	Create(ctx context.Context, dto dto.CreateUserDTO) error
	GetByID(ctx context.Context, id int64) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	DeleteByID(ctx context.Context, id int64) error
}

type UserUseCase struct {
	repo UserRepo
}

func NewUserUseCase(r UserRepo) *UserUseCase {
	return &UserUseCase{repo: r}
}

func (uc *UserUseCase) Create(ctx context.Context, dto dto.CreateUserDTO) error {
	err := uc.repo.Create(ctx, dto)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UserUseCase) GetByID(ctx context.Context, id int64) (*entity.User, error) {
	user, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *UserUseCase) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	user, err := uc.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *UserUseCase) DeleteByID(ctx context.Context, id int64) error {
	err := uc.repo.DeleteByID(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
