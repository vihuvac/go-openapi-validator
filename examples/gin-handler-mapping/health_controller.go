package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

type ApiResponse struct {
	Message string `json:"message"`
}

func (h *HealthController) CheckLiveness(c *gin.Context) {
	c.JSON(http.StatusOK, ApiResponse{
		Message: "Ok",
	})
}

func (h *HealthController) CheckReadiness(c *gin.Context) {
	c.JSON(http.StatusOK, ApiResponse{
		Message: "Ok",
	})
}
