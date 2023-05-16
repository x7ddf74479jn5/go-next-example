package main

import (
	"backend/app/middleware"
	"net/http"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middleware.Cors())

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "hello world",
			})
		})
	}

	router.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		extension := filepath.Ext(file)

		if file == "" || extension == "" {
			path := filepath.Join(dir, "index.html")
			c.File("../frontend/out" + path)
		} else {
			c.File("../frontend/out" + c.Request.RequestURI)
		}
	})

	router.Run(":8080")
}
