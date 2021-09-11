package routers

import "github.com/gin-gonic/gin"

func InitRoutes(engine *gin.Engine) {
	SetCommonRoutes(engine)
	SetCommentsRoutes(engine)
	SetSavingRoutes(engine)
}
