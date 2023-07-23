package api

import (
	"context"
	"encoding/json"
	"github.com/Khasmag06/gophkeeper/internal/models"
	response "github.com/Khasmag06/gophkeeper/pkg/http"
	"github.com/gin-gonic/gin"
)

func (h *Handler) AddRecord(c *gin.Context, recordReq models.Record, record any) {
	ctx := context.Background()

	err := c.BindJSON(&recordReq)
	if err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	decryptRecord, err := h.decoder.Decrypt(recordReq.EncryptedData)
	if err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	err = json.Unmarshal(decryptRecord, &record)
	if err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	err = h.recordService.AddRecord(ctx, c.GetInt(userIdParam), record)
	if err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	response.WriteSuccessResponse(c, nil)
}

func (h *Handler) AddLoginCredsRecords(c *gin.Context) {
	var recordReq models.Record
	var loginCreds models.LoginCredentials

	h.AddRecord(c, recordReq, &loginCreds)
}

func (h *Handler) AddTextDataRecord(c *gin.Context) {
	var recordReq models.Record
	var textData models.TextData

	h.AddRecord(c, recordReq, &textData)
}

func (h *Handler) AddBinaryDataRecord(c *gin.Context) {
	var recordReq models.Record
	var binaryData models.BinaryData

	h.AddRecord(c, recordReq, &binaryData)
}

func (h *Handler) AddBankCardRecord(c *gin.Context) {
	var recordReq models.Record
	var bankCard models.BankCard

	h.AddRecord(c, recordReq, &bankCard)
}
