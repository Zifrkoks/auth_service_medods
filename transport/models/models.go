package models

type (
	AuthSwag struct {
		Uuid string `json:"uuid" example:"uuid" binding:"uuid"`
	}
	RefreshSwag struct {
		Jwt     string `json:"jwt" example:"jwt" binding:"jwt"`
		Refresh string `json:"refresh" example:"refresh" binding:"refresh"`
	}
	JwtSwag struct {
		Jwt string `json:"jwt" example:"jwt" binding:"jwt"`
	}
)
