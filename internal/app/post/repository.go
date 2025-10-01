package post

import (
	"context"
	"errors"
	"go-gin-boilerplate/entity"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

var Repo *repository

func InitRepository(db *gorm.DB) {
	Repo = &repository{db: db}
}

func (r *repository) getDB(ctx context.Context, db *gorm.DB) *gorm.DB {
	if db != nil {
		return db.WithContext(ctx)
	}
	return r.db.WithContext(ctx)
}

func (r *repository) Create(ctx context.Context, post *entity.Post, db *gorm.DB) error {
	return r.getDB(ctx, db).Create(post).Error
}

func (r *repository) FindByID(ctx context.Context, id uint64, db *gorm.DB) (*entity.Post, error) {
	var post entity.Post
	err := r.getDB(ctx, db).First(&post, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &post, err
}

func (r *repository) FindByAuthorIDCursor(
	ctx context.Context,
	authorID uint64,
	lastID uint64,
	pageSize int,
	db *gorm.DB,
) ([]entity.Post, error) {
	var posts []entity.Post

	query := r.getDB(ctx, db).
		Where("author_id = ?", authorID).
		Limit(pageSize).
		Order("id DESC")

	if lastID > 0 {
		query = query.Where("id < ?", lastID)
	}

	err := query.Find(&posts).Error
	if err != nil {
		return nil, err
	}

	return posts, nil
}
