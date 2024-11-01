package delivery

import (
	"github.com/dyaksa/dating-app/internal/dto"
	"github.com/dyaksa/dating-app/internal/modules/purchases/v1/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RestPurchasesHandler struct {
	usecase usecase.PurchasesUsecase
}

func NewRestPurchasesHandler(usecase usecase.PurchasesUsecase) *RestPurchasesHandler {
	return &RestPurchasesHandler{
		usecase: usecase,
	}
}

func (h *RestPurchasesHandler) Purchases(ctx *gin.Context) {
	var payload dto.PurchasesRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userID := ctx.MustGet("user_id").(string)
	parseID, err := uuid.Parse(userID)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.Purchases(ctx, parseID, &payload); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "success purchases"})
}
