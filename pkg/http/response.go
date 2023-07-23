package http

import (
	"errors"
	"github.com/Khasmag06/gophkeeper/pkg/app_err"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SuccessResponse struct {
	Status string `json:"status"`
	Data   any    `json:"data,omitempty"`
}

type ErrorResponse struct {
	Status string `json:"status"`
	Error  `json:"error"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type logger interface {
	Info(text ...any)
	Warn(text ...any)
	Error(text ...any)
}

func WriteSuccessResponse(c *gin.Context, data any) {
	successResponse := SuccessResponse{
		Status: "success",
		Data:   data,
	}

	c.JSON(http.StatusOK, successResponse)
}

func WriteErrorResponse(c *gin.Context, logger logger, err error) {
	var bErr app_err.BusinessError
	var sErr app_err.SilentError

	if errors.As(err, &bErr) {
		errorResponse := ErrorResponse{
			Status: "error",
			Error: Error{
				Code:    bErr.Code(),
				Message: bErr.Error(),
			},
		}

		logger.Warn(err)

		c.JSON(http.StatusBadRequest, errorResponse)

	} else if errors.As(err, &sErr) {
		logger.Error(err)

		c.JSON(http.StatusOK, nil)

	} else {
		errorResponse := ErrorResponse{
			Status: "error",
			Error: Error{
				Code:    "InternalServerError",
				Message: "Что-то пошло не так, попробуйте еще раз",
			},
		}

		logger.Error(err)

		c.JSON(http.StatusInternalServerError, errorResponse)
	}
}
