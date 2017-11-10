package main

import (
	"github.com/gin-gonic/gin"
	"github.com/konojunya/go-simple-crud/controller"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.Static("/js", "./public/js")
	r.Static("/image", "./public/image")
	r.Static("/css", "./public/css")
	r.LoadHTMLGlob("view/*")

	// api routing
	api := r.Group("/api")
	api.GET("/users", controller.FindUser)
	api.GET("/users/:name", controller.FindUserByName)
	api.POST("/users", controller.CreateUser)
	api.PUT("/users/:name", controller.EditUser)
	api.DELETE("/users/:name", controller.DeleteUser)

	r.Run(":2000")
}
