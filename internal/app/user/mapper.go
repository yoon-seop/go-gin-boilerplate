package user

import (
	"go-gin-boilerplate/entity"
)

func ToV1(u *entity.User) Response {
	return Response{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
	}
}
