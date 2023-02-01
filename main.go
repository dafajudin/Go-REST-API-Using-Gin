package main

import (
  	"github.com/gin-gonic/gin"
	"go-api/controllers"
	"go-api/models"
	"go-api/middleware"
)

func main() {

	models.ConnectDataBase()
	
	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login",controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middleware.JwtAuthMiddleware())
	protected.GET("/user",controllers.CurrentUser)

	r.Run(":8080")

}