package main

import (
	"backend/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://192.168.88.239:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	// Auth routes
	r.POST("/api/register", api.Register)
	r.POST("/api/login", api.Login)
	r.POST("/api/logout", api.Logout)
	r.GET("/api/campaigns", api.GetCampaigns)

	r.Run(":8080")
}
