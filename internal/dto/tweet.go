package dto

import "time"

type CreateTweetParams struct {
	Content string `json:"content"`
}

type CreateTweetDTO struct {
	AuthorID int64  `json:"author_id"`
	Content  string `json:"content"`
}

type GetTweetByIdParams struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type GetTweetAuthorIdDTO struct {
	AuthorID int64 `json:"author_id"`
}

type GetTweetByIdDTO struct {
	ID         int64     `json:"id"`
	AuthorName string    `json:"author_name"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
}

type ListAllTweetsParams struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type ListTweetsRequest struct {
	PageID   int `form:"page_id" binding:"required,min=1"`
	PageSize int `form:"page_size" binding:"required,min=5,max=10"`
}

type DeleteTweetByIdParams struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}
