package notification

import (
	"net/http"

	"github.com/TitusW/notifications/internal/entity"
	"github.com/gin-gonic/gin"
)

func (h Handler) CreateNotification(ctx *gin.Context) {
	var createNotification entity.CreateNotification
	if err := ctx.ShouldBindJSON(&createNotification); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err := h.uc.CreateNotification(ctx, createNotification)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Notification created! Dispatch in progress",
	})
}
