package post

import "time"

type CreateRequest struct {
	AuthorID    uint64 `json:"authorId,omitempty"`
	Title       string `json:"title,omitempty"`
	Content     string `json:"content,omitempty"`
	IsPublished bool   `json:"isPublished,omitempty"`
}

type Response struct {
	ID          uint64    `json:"id,omitempty"`
	AuthorID    uint64    `json:"authorId,omitempty"`
	Title       string    `json:"title,omitempty"`
	Content     string    `json:"content,omitempty"`
	IsPublished bool      `json:"isPublished,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
}
