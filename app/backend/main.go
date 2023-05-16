package main

import (
	"backend/app/middleware"
	"backend/env"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middleware.Cors())
	env.Load()
	port := os.Getenv("APP_PORT")
	staticFilePath := os.Getenv("STATIC_FILE_PATH")

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "hello world",
			})
		})
	}

	router.NoRoute(func(c *gin.Context) {
		var url string
		if c.Request.RequestURI == "/" {
			url = staticFilePath + "/index.html"
		} else {
			url = staticFilePath + strings.TrimSuffix(c.Request.RequestURI, "/") + ".html"
		}

		resp, err := http.Get(url)
		if err != nil || resp.StatusCode == 404 {
			url = staticFilePath + "/404.html"
			resp, err = http.Get(url)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "failed to fetch the file",
					"error":   err.Error(),
				})
				return
			}
		}
		defer resp.Body.Close()

		c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, map[string]string{})
	})

	router.Run(":" + port)
}
