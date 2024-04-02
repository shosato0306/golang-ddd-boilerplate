package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/shosato0306/golang-ddd-boilerplate/internal/usecase"
)

type APIHandler struct {
	userUseCase usecase.User
}

func NewAPIHandler(userUseCase usecase.User) *APIHandler {
	return &APIHandler{userUseCase: userUseCase}
}

func (h *APIHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.userUseCase.GetUser(c, id)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	})
}

func (h *APIHandler) CreateUser(c *gin.Context) {
	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	user, err := h.userUseCase.CreateUser(c, req.Name, req.Email)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	})
}
