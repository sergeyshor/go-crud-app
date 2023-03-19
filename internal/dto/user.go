package dto

import "time"

type CreateUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserByIdParams struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type GetUserByIdDTO struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DeleteUserByIdParams struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}
