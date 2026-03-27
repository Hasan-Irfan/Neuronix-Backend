package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// In-memory user store
var users = map[string]string{}

func Register(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	users[body.Email] = body.Password

	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
	})
}

func Login(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if users[body.Email] != body.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials",
		})
		return
	}

	token := "abc123"

	c.SetCookie(
		"token",     // name
		token,       // value
		3600,        // maxAge (seconds)
		"/",         // path
		"localhost", // domain (important)
		false,       // secure (true in production with HTTPS)
		true,        // httpOnly (VERY IMPORTANT)
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
	})
}

func Logout(c *gin.Context) {
	// Clear cookie by setting maxAge = -1
	c.SetCookie(
		"token",
		"",
		-1, // expire immediately
		"/",
		"localhost",
		false,
		true,
	)

	c.JSON(200, gin.H{
		"message": "Logged out successfully",
	})
}
