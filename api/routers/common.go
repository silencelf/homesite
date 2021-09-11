package routers

import (
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func SetCommonRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "welcome!",
		})
	})

	router.GET("/layout", func(c *gin.Context) {
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"title": "welcome!",
		})
	})

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/ips", func(c *gin.Context) {
		ls := exec.Command("ls")
		b, err := ls.CombinedOutput()
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
		c.String(http.StatusOK, string(b))
	})
}
