package entity

import "time"

type Tweet struct {
	ID        int64     `json:"id"`
	AuthorID  int64     `json:"author_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
