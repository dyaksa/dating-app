package delivery

import (
	"github.com/dyaksa/dating-app/internal/dto"
	"github.com/dyaksa/dating-app/internal/modules/swipes/v1/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RestSwipesHandler struct {
	usecase usecase.SwipeUsecase
}

func NewRestSwipesHandler(usecase usecase.SwipeUsecase) *RestSwipesHandler {
	return &RestSwipesHandler{
		usecase: usecase,
	}
}

func (h *RestSwipesHandler) SwipeUser(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(string)
	var payload dto.SwipesRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	parseID, err := uuid.Parse(userID)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err = h.usecase.SwipeUser(ctx, parseID, &payload); err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "swipe success",
	})
}

func (h *RestSwipesHandler) GetSwipeUser(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(string)

	parseID, err := uuid.Parse(userID)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	user, err := h.usecase.GetSwipeUser(ctx, parseID)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": user,
	})
}
