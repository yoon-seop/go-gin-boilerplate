package user

import (
	"context"
	"errors"
	"fmt"
	"go-gin-boilerplate/entity"
	"go-gin-boilerplate/pkg/common"

	"github.com/go-sql-driver/mysql"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(ctx context.Context, req CreateRequest) (*entity.User, error) {
	if req.Name == "" || req.Email == "" || req.Password == "" {
		return nil, common.RequiredFieldMissingError
	}

	existing, err := Repo.FindByEmail(ctx, req.Email, nil)
	if err != nil {
		return nil, fmt.Errorf("find by email failed: %w", err)
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

	if err = Repo.Create(ctx, user, nil); err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return nil, common.EmailAlreadyExistsError
		}
		return nil, fmt.Errorf("create user failed: %w", err)
	}

	return user, nil
}

func GetUserByID(ctx context.Context, id uint64) (*entity.User, error) {
	user, err := Repo.FindByID(ctx, id, nil)
	if err != nil {
		return nil, fmt.Errorf("find user by ID failed: %w", err)
	}

	if user == nil {
		return nil, common.UserNotFoundError
	}

	return user, nil
}
