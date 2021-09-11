package routers

import (
	"homesite/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	saving models.Saving
)

func SetSavingRoutes(router *gin.Engine) {
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
}
