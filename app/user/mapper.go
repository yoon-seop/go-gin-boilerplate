package user

import (
	"go-gin-boilerplate/entity"
)

func ToV1Response(u *entity.User) Response {
	return Response{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
	}
}
