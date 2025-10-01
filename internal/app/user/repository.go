package user

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

func (r *repository) Create(ctx context.Context, u *entity.User, db *gorm.DB) error {
	return r.getDB(ctx, db).Create(u).Error
}

func (r *repository) FindByID(ctx context.Context, id uint64, db *gorm.DB) (*entity.User, error) {
	var u entity.User
	err := r.getDB(ctx, db).First(&u, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &u, err
}

func (r *repository) FindByEmail(ctx context.Context, email string, db *gorm.DB) (*entity.User, error) {
	var u entity.User
	err := r.getDB(ctx, db).Where("email = ?", email).First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &u, err
}

func (r *repository) Update(ctx context.Context, u *entity.User, db *gorm.DB) error {
	return r.getDB(ctx, db).Save(u).Error
}

func (r *repository) Delete(ctx context.Context, id uint64, db *gorm.DB) error {
	return r.getDB(ctx, db).Delete(&entity.User{}, id).Error
}
