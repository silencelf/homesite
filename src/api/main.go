package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

func main() {
	router := gin.Default()

	router.Use(favicon.New("./favicon.ico"))
	//router.StaticFile("/favicon.ico", "assets/favicon.ico")
	router.LoadHTMLGlob("templates/*")

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main WebSite",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	log.Print("web is running...")
	router.Run(":8000")
}
