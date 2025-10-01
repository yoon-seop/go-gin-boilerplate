package handler

import (
	"go-gin-boilerplate/internal/app/post"
	"go-gin-boilerplate/pkg/common"

	"github.com/gin-gonic/gin"
)

func NewPostHandler(r *gin.Engine) {
	group := r.Group("/post")

	group.POST("", postPost)
}

func postPost(c *gin.Context) {
	var req post.CreateRequest
	if ok := common.BindJSONBody(c, &req); !ok {
		return
	}

	data, err := post.CreatePost(c, req)
	if err != nil {
		common.Fail(c, err)
		return
	}

	common.Success(c, post.Response{
		ID:          data.ID,
		AuthorID:    data.AuthorID,
		Title:       data.Title,
		Content:     data.Content,
		IsPublished: data.IsPublished,
		CreatedAt:   data.CreatedAt,
	})
}
