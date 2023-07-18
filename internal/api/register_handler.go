package api

import (
	"context"
	"fmt"
	"github.com/Khasmag06/gophkeeper/internal/models"
	"github.com/Khasmag06/gophkeeper/pkg/app_err"
	response "github.com/Khasmag06/gophkeeper/pkg/http"
	"github.com/gin-gonic/gin"
	"strings"
	"unicode/utf8"
)

const allowedPasswordLen = 8

func (h *Handler) SignUp(c *gin.Context) {
	var signUpReq models.User
	ctx := context.Background()

	if err := c.BindJSON(&signUpReq); err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	err := checkReqData(signUpReq)
	if err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	signUpReq.Login = strings.TrimSpace(strings.ToLower(signUpReq.Login))

	err = validatePassword(signUpReq.Password)
	if err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	err = h.authService.SignUp(ctx, signUpReq)
	if err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	response.WriteSuccessResponse(c, nil)
}

func checkReqData(data models.User) error {
	if data.Login == "" {
		return app_err.NewAuthorizationError("логин не указан")
	}
	if data.Password == "" {
		return app_err.NewAuthorizationError("пароль не указан")
	}

	return nil
}

func validatePassword(password string) error {
	if utf8.RuneCountInString(password) < allowedPasswordLen {
		return app_err.NewBusinessError(fmt.Sprintf("длина пароля должна быть от %d символов", allowedPasswordLen))
	}

	return nil
}
