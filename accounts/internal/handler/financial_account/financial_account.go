package financialaccount

import (
	"net/http"

	"github.com/TitusW/accounts/internal/entity"
	"github.com/gin-gonic/gin"
)

func (h *Handler) DebitFinancialAccount(ctx *gin.Context) {
	var debitFinancialAccount entity.DebitFinancialAccount
	if err := ctx.ShouldBindJSON(&debitFinancialAccount); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if debitFinancialAccount.Amount < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Insufficient Fund",
		})
		return
	}

	financialAccount, err := h.uc.DebitFinancialAccount(ctx, debitFinancialAccount)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": financialAccount,
	})
}

func (h *Handler) CreditFinancialAccount(ctx *gin.Context) {
	var creditFinancialAccount entity.CreditFinancialAccount
	if err := ctx.ShouldBindJSON(&creditFinancialAccount); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if creditFinancialAccount.Amount < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Insufficient Amount",
		})
		return
	}

	financialAccount, err := h.uc.CreditFinancialAccount(ctx, creditFinancialAccount)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": financialAccount,
	})
}
