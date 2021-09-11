package routers

import (
	"homesite/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetCommentsRoutes(router *gin.Engine) {
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

}
