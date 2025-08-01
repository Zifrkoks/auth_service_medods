package routes

import (
	app "auth_service_medods/internal/app"
	"auth_service_medods/internal/domain/models"
	view "auth_service_medods/transport/models"

	"github.com/gin-gonic/gin"
)

// @Summary	Login
// @Tags		AUTH
// @Accept		json
// @Produce	json
// @Param		Data	body		view.AuthSwag	true	"Login form"
// @Success	200		{string}	string
// @Failure	400		{string}	string
// @Failure	404		{string}	string
// @Router		/auth/login [post]
func Login(c *gin.Context) {
	var authReq view.AuthSwag
	var authdata models.AuthData
	if err := c.BindJSON(&authReq); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	authdata.UserAgent = c.Request.UserAgent()
	authdata.Ip = c.RemoteIP()
	authdata.Id = authReq.Uuid
	service := app.NewAuthService()
	tokens, err := service.Auth(authdata)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"tokens": tokens})

}

// @Summary	Refresh
// @Tags		AUTH
// @Accept		json
// @Produce	json
// @Param		Data	body		view.RefreshSwag	true	"refresh form"
// @Success	200		{string}	string
// @Failure	400		{string}	string
// @Failure	404		{string}	string
// @Router		/auth/refresh [post]
func Refresh(c *gin.Context) {
	var refreshReq view.RefreshSwag
	var refreshdata models.RefreshData
	if err := c.BindJSON(&refreshReq); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	refreshdata.UserAgent = c.Request.UserAgent()
	refreshdata.Ip = c.RemoteIP()
	refreshdata.Jwt = refreshReq.Jwt
	refreshdata.Refresh = refreshReq.Refresh
	service := app.NewAuthService()
	tokens, err := service.RefreshAuthTokens(refreshdata)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"tokens": tokens})

}
