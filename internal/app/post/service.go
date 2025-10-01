package post

import (
	"context"
	"fmt"
	"go-gin-boilerplate/entity"
)

func CreatePost(ctx context.Context, req CreateRequest) (*entity.Post, error) {
	post := &entity.Post{
		AuthorID:    req.AuthorID,
		Title:       req.Title,
		Content:     req.Content,
		IsPublished: req.IsPublished,
	}

	if err := Repo.Create(ctx, post, nil); err != nil {
		return nil, fmt.Errorf("create post failed: %w", err)
	}

	return post, nil
}

func GetPostByAuthorID(ctx context.Context, authorID uint64, lastID uint64, pageSize int) ([]entity.Post, error) {
	posts, err := Repo.FindByAuthorIDCursor(ctx, authorID, lastID, pageSize, nil)
	if err != nil {
		return nil, fmt.Errorf("create post failed: %w", err)
	}

	return posts, nil
}
