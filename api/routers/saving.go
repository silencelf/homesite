package routers

import (
	"homesite/domains"
	"homesite/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	savings = domains.NewSavingService()
)

func SetSavingRoutes(router *gin.Engine) {
	router.POST("/api/savings", func(c *gin.Context) {
		var saving models.AddSavingModel
		if err := c.ShouldBind(&saving); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		savings.Savings = append(savings.Savings, models.Saving{
			UserId: saving.UserId,
			Desc:   saving.Desc,
			Target: saving.Target,
		})

		c.String(http.StatusCreated, "ok")
	})

	router.POST("/api/savings/:id/details", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		}
		s, err := savings.FindById(id)
		if err != nil {
			c.String(http.StatusNotFound, "not found")
		}
		var add models.AddSavingDetailModel
		if err := c.ShouldBind(&add); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		s.Details = append(s.Details, add)
		s.Achieved += add.Amount

		c.String(http.StatusOK, "ok")
	})

	router.GET("/api/savings", func(c *gin.Context) {
		c.JSON(200, savings)
	})

	router.GET("/api/savings/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusBadRequest, "invalid id")
		}
		s, err := savings.FindById(id)
		if err != nil {
			c.String(http.StatusNotFound, "not found")
		} else {
			c.JSON(200, s)
		}
	})
}
