package controllers

import (
	"homesite/models"
	"log"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

var (
	saving   models.Saving
	comments []models.Comment
)

func Register(router *gin.Engine) {
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

	router.GET("/comments", func(c *gin.Context) {
		c.JSON(http.StatusOK, comments)
	})

	router.POST("/comments", func(c *gin.Context) {
		var cm models.Comment
		err := c.BindJSON(&cm)
		if err != nil {
			log.Printf("failed to create comment: %s", err.Error())
		}
		log.Println(cm)
		// of course it's not thread safe
		comments = append(comments, cm)
	})

	router.POST("/saving", func(c *gin.Context) {
		var amount models.AddSavingModel

		if err := c.ShouldBind(&amount); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		saving.Total += amount.Amount

		c.String(http.StatusOK, "ok")
	})

	router.GET("/saving", func(c *gin.Context) {
		c.JSON(200, saving)
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
