package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hal-iosk/hal-cinema/middleware"
	cors "github.com/itsjamie/gin-cors"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.Static("/js", "./public/js")
	r.Static("/image", "./public/image")
	r.Static("/css", "./public/css")
	r.LoadHTMLGlob("view/*")

	// view routing
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", nil)
	})

	// api routing
	api := r.Group("/api")
	api.Use(cors.Middleware(middleware.CorsConfig))

	r.Run(":2000")
}
