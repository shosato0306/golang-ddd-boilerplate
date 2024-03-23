package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/shosato0306/golang-ddd-boilerplate/internal/usecase"
)

type ExampleHandler struct {
	userUseCase usecase.User
}

func NewExampleHandler(userUseCase usecase.User) *ExampleHandler {
	return &ExampleHandler{userUseCase: userUseCase}
}

func (h *ExampleHandler) GetUser(c *gin.Context) {
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
