package handlers

import (
	"net/http"
	"strconv"

	"tr-search-back/internal/domain/user"
	"tr-search-back/internal/usecases"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	createUserUseCase   *usecases.UseCase
	deleteUserUseCase   *usecases.DeleteUserUseCase
	updateUserUseCase   *usecases.UpdateUserUseCase
	getUserEmailUseCase *usecases.GetUserEmailUseCase
}

func NewUserHandler(createUserUseCase *usecases.UseCase, deleteUserUseCase *usecases.DeleteUserUseCase, updateUserUseCase *usecases.UpdateUserUseCase, getUserEmailUseCase *usecases.GetUserEmailUseCase) *UserHandler {
	return &UserHandler{createUserUseCase: createUserUseCase, deleteUserUseCase: deleteUserUseCase, updateUserUseCase: updateUserUseCase, getUserEmailUseCase: getUserEmailUseCase}
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

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var u user.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u.ID = uint(id)

	if err := h.updateUserUseCase.Execute(&u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, u)
}

func (h *UserHandler) GetUserEmail(c *gin.Context) {
	email := c.Param("email")
	u, err := h.getUserEmailUseCase.Execute(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, u)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.deleteUserUseCase.Execute(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
