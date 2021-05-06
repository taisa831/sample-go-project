package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taisa831/sample-go-project/usecase"
)

type UserHandler struct {
	u *usecase.UserUsecase
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		usecase.NewUserUsecase(),
	}
}

func (h *UserHandler) List(c *gin.Context) {
	users := h.u.List()
	c.JSON(http.StatusOK, users)
}
