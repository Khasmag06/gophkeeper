package api

import (
	"context"
	"github.com/Khasmag06/gophkeeper/internal/models"
	"github.com/Khasmag06/gophkeeper/pkg/app_err"
	response "github.com/Khasmag06/gophkeeper/pkg/http"
	"github.com/gin-gonic/gin"
	"strings"
)

func (h *Handler) Login(c *gin.Context) {
	var loginReq models.User
	ctx := context.Background()

	err := c.BindJSON(&loginReq)
	if err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	err = checkRequestData(loginReq)
	if err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	loginReq.Login = strings.TrimSpace(strings.ToLower(loginReq.Login))

	tokenData, err := h.authService.Login(ctx, loginReq)
	if err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	response.WriteSuccessResponse(c, tokenData)
}

func checkRequestData(data models.User) error {
	if data.Login == "" {
		return app_err.NewAuthorizationError("логин не указан")
	}
	if data.Password == "" {
		return app_err.NewBusinessError("пароль не указан")
	}

	return nil
}
