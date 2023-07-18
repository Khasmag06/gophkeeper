package api

import (
	"context"
	"github.com/Khasmag06/gophkeeper/internal/models"
	response "github.com/Khasmag06/gophkeeper/pkg/http"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetBankCardRecords(c *gin.Context) {
	ctx := context.Background()
	var bankCardType models.BankCard
	bankCards, err := h.recordService.GetAllRecords(ctx, c.GetInt(userIdParam), bankCardType)
	if err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	response.WriteSuccessResponse(c, bankCards)

}

func (h *Handler) GetLoginCredentialsRecords(c *gin.Context) {
	ctx := context.Background()
	var loginCredsType models.LoginCredentials
	loginCreds, err := h.recordService.GetAllRecords(ctx, c.GetInt(userIdParam), loginCredsType)
	if err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	response.WriteSuccessResponse(c, loginCreds)

}

func (h *Handler) GetTextDataRecords(c *gin.Context) {
	ctx := context.Background()
	var textDataType models.TextData
	textData, err := h.recordService.GetAllRecords(ctx, c.GetInt(userIdParam), textDataType)
	if err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	response.WriteSuccessResponse(c, textData)

}

func (h *Handler) GetBinaryDataRecords(c *gin.Context) {
	ctx := context.Background()
	var binaryDataType models.BinaryData
	binaryData, err := h.recordService.GetAllRecords(ctx, c.GetInt(userIdParam), binaryDataType)
	if err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	response.WriteSuccessResponse(c, binaryData)

}
