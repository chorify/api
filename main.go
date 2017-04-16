package main

import (
	"github.com/chorify/api/server"
	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	server.Register(e)
	e.Run()
}
