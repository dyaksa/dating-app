package delivery

import (
	"net/http"

	"github.com/dyaksa/dating-app/internal/dto"
	"github.com/dyaksa/dating-app/internal/modules/auth/v1/usecase"
	"github.com/gin-gonic/gin"
)

type RestAuthHandler struct {
	AuthUsecase usecase.Authusecase
}

func NewRestAuthHandler(authUsecase usecase.Authusecase) *RestAuthHandler {
	return &RestAuthHandler{AuthUsecase: authUsecase}
}

func (h *RestAuthHandler) Register(c *gin.Context) {
	var payload dto.RegisterRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.AuthUsecase.Register(c, &payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}

func (h *RestAuthHandler) Login(c *gin.Context) {
	var payload dto.LoginRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.AuthUsecase.Login(c, &payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": res})
}
