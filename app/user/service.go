package user

import (
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"go-gin-boilerplate/common"
	"go-gin-boilerplate/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(req CreateRequest) (*entity.User, error) {
	if req.Name == "" || req.Email == "" || req.Password == "" {
		return nil, common.RequiredFieldMissingError
	}

	existing, err := Repository.FindByEmail(req.Email)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("find by email failed: %w", err)
		}
	}
	if existing != nil {
		return nil, common.EmailAlreadyExistsError
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("password hash failed: %w", err)
	}

	user := &entity.User{
		Email:    req.Email,
		Password: string(hash),
		Name:     req.Name,
	}

	if err = Repository.Create(user); err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return nil, common.EmailAlreadyExistsError
		}
		return nil, fmt.Errorf("create user failed: %w", err)
	}

	return user, nil
}

func GetUserByID(id uint64) (*entity.User, error) {
	user, err := Repository.FindByID(id)
	if err != nil {
		return nil, common.UserNotFoundError
	}

	return user, nil
}
