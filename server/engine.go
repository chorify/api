package server

import (
	"github.com/chorify/api/server/routes"
	"github.com/gin-gonic/gin"
)

// Register registers routes to a gin engine
func Register(e *gin.Engine) {
	e.GET("/", routes.Status())
}
