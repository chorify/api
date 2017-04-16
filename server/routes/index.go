package routes

import (
	"github.com/gin-gonic/gin"
)

// Status implements a status endpoint for the api
func Status() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, map[string]string{
			"app":         "chorify-api",
			"version":     "0.01",
			"description": "api for handling backend needs for chorify",
			"author":      "@bphipps",
		})
	}
}
