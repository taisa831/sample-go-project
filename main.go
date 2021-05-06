package main

import (
	"github.com/gin-gonic/gin"
	"github.com/taisa831/sample-go-project/interfaces/handler"
)

func main() {
	r := gin.New()
	userHandler := handler.NewUserHandler()
	r.GET("/users", userHandler.List)
	r.Run()
}
