package routes

import (
	"auth_service_medods/internal/app"

	"github.com/gin-gonic/gin"
)

// @Summary	Get user's UUID
// @Tags	USER
// @Accept	json
// @Produce	json
// @Success	200			{string}	string
// @Failure	400			{string}	string
// @Router		/data/me/ 	[get]
// @Security	ApiKeyAuth
func GetUUID(c *gin.Context) {
	data := app.NewDataService()
	resp, err := data.GetUserData(c.GetString("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"user": resp})
}

// @Summary	Logout
// @Tags	USER
// @Accept	json
// @Produce	json
// @Success	200	{string}	string
// @Failure	400	{string}	string
// @Failure	404	{string}	string
// @Router		/data/logout [post]
// @Security	ApiKeyAuth
func Logout(c *gin.Context) {

	service := app.NewAuthService()
	 err := service.Logout(c.GetString("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"result": "ok"})
}
