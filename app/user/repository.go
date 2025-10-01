package user

import (
	"go-gin-boilerplate/entity"
	"gorm.io/gorm"
)

type RepositoryInterface interface {
	Create(user *entity.User) error
	FindByID(id uint64) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Update(user *entity.User) error
	Delete(id uint64) error

	WithTx(tx *gorm.DB) RepositoryInterface
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}
func (r *userRepository) FindByID(id uint64) (*entity.User, error) {
	var user entity.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) Update(user *entity.User) error {
	return r.db.Save(user).Error
}
func (r *userRepository) Delete(id uint64) error {
	return r.db.Delete(&entity.User{}, id).Error
}

func (r *userRepository) WithTx(tx *gorm.DB) RepositoryInterface {
	return &userRepository{db: tx}
}

var Repository RepositoryInterface

func RepositoryInitialize(db *gorm.DB) {
	Repository = &userRepository{db: db}
}
