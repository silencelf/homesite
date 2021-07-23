package main

import (
	"flag"
	"homesite/controllers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

func main() {
	var port = flag.String("p", "80", "port")
	flag.Parse()
	router := gin.Default()

	router.Use(favicon.New("./favicon.ico"))
	router.Static("/assets", "./assets/")
	router.LoadHTMLGlob("templates/*")

	controllers.Register(router)

	log.Print("web is running...")
	router.Run(":" + *port)
}
