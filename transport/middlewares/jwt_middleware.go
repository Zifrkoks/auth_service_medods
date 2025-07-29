package middlewares

import (
	app "auth_service_medods/internal/app"
	"auth_service_medods/internal/logger"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Jwt_middleware(c *gin.Context) {
	tokenString, _ := strings.CutPrefix(c.Request.Header.Get("Authorization"), "Bearer ")
	service := app.NewAuthService()
	claims, err := service.ValidateJWT(tokenString)
	if err != nil {
		logger.LogImportant("Unauthorized by middleware")
		c.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}
	c.Set("id", claims.Subject)
	c.Next()
}
