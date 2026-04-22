package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers handles /api/v1/users
func (Server) GetUsers(ctx *gin.Context) {
	userID := "test"
	resp := []User{
		{
			UserId: &userID,
		},
	}
	ctx.JSON(http.StatusOK, resp)
}

// GetUser handles /api/v1/user/:id
func (Server) GetUser(ctx *gin.Context, id string, params GetUserParams) {
	userID := "test"
	ctx.JSON(http.StatusOK, User{UserId: &userID})
}
