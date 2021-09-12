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
	router.POST("/saving/:id/details", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusBadRequest, "invalid id")
		}
		s, err := savings.FindById(id)
		if err != nil {
			c.String(http.StatusNotFound, "not found")
		}
		var add models.AddSavingModel
		if err := c.ShouldBind(&add); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		s.Details = append(s.Details, add)
		s.Achieved += add.Amount

		c.String(http.StatusOK, "ok")
	})

	router.GET("/savings/:id", func(c *gin.Context) {
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
