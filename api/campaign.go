package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCampaigns(c *gin.Context) {
	c.JSON(http.StatusOK, []gin.H{
		{"id": 1, "name": "Summer Sale", "status": "Active"},
		{"id": 2, "name": "Winter Promo", "status": "Paused"},
	})
}
