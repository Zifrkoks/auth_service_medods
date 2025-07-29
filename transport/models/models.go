package models

type (
	AuthSwag struct {
		Uuid string `json:"uuid" example:"uuid" binding:"required"`
	}
	RefreshSwag struct {
		Jwt     string `json:"jwt" example:"jwt" binding:"required"`
		Refresh string `json:"refresh" example:"refresh" binding:"required"`
	}
	JwtSwag struct {
		Jwt string `json:"jwt" example:"jwt" binding:"required"`
	}
)
