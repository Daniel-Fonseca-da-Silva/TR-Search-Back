package handlers

import (
	"net/http"

	"tr-search-back/internal/domain/user"
	"tr-search-back/internal/usecases"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	createUserUseCase *usecases.UseCase
}

func NewUserHandler(createUserUseCase *usecases.UseCase) *UserHandler {
	return &UserHandler{createUserUseCase: createUserUseCase}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var u user.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.createUserUseCase.Execute(&u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, u)
}
