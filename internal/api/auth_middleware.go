package api

import (
	"github.com/Khasmag06/gophkeeper/pkg/app_err"
	response "github.com/Khasmag06/gophkeeper/pkg/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userIdParam         = "userId"
)

func (h *Handler) authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := strings.Split(c.GetHeader(authorizationHeader), "Bearer ")
		if len(authHeader) != 2 {
			response.WriteErrorResponse(c, h.logger, app_err.NewUnauthorizedError())
			h.logger.Error(`header 'Authorization' невалиден`)
			c.Abort()
			return
		}

		accessToken := authHeader[1]

		claims, err := h.authService.ParseToken(accessToken)
		if err != nil {
			response.WriteErrorResponse(c, h.logger, app_err.NewUnauthorizedError())
			h.logger.Error(err)
			c.Abort()
			return
		}

		userIdBytes, err := h.decoder.Decrypt(claims.UserID)
		if err != nil {
			response.WriteErrorResponse(c, h.logger, err)
			c.Abort()
			return
		}
		userID, _ := strconv.Atoi(string(userIdBytes))
		c.Set(userIdParam, userID)
		c.Next()
	}
}
