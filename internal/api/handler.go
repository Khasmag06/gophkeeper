package api

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	*gin.Engine
	authService   authService
	recordService recordService
	decoder       decoder
	logger        logger
}

func NewHandler(auth authService, record recordService, decoder decoder, logger logger) *Handler {
	h := &Handler{
		Engine:        gin.New(),
		authService:   auth,
		recordService: record,
		decoder:       decoder,
		logger:        logger,
	}
	h.Use(gin.Recovery())

	user := h.Group("/api/user")

	//user
	user.POST("/signup", h.SignUp)
	user.POST("/login", h.Login)

	api := h.Group("/api")
	api.Use(h.authMiddleware())

	api.GET("/record/login-creds", h.GetLoginCredentialsRecords)
	api.POST("/record/login-creds/add", h.AddLoginCredsRecords)

	api.GET("/record/bank-card", h.GetBankCardRecords)
	api.POST("/record/bank-card/add", h.AddBankCardRecord)

	api.GET("/record/binary", h.GetBinaryDataRecords)
	api.POST("/record/binary/add", h.AddBinaryDataRecord)

	api.GET("/record/text", h.GetTextDataRecords)
	api.POST("/record/text/add", h.AddTextDataRecord)

	return h

}
