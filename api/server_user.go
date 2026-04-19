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

func (Server) GetUser(ctx *gin.Context, id string) {
	userID := "test"
	ctx.JSON(http.StatusOK, User{UserId: &userID})
}

func (Server) GetUserProfile(ctx *gin.Context, id string) {
	ctx.JSON(http.StatusOK, UserProfile{})
}
