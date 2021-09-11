package main

import (
	"flag"
	"homesite/routers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

func main() {
	var addr = flag.String("a", ":8080", "listening address")
	flag.Parse()
	engine := gin.Default()

	engine.Use(favicon.New("./favicon.ico"))
	engine.Static("/assets", "./assets/")
	engine.LoadHTMLGlob("templates/*")

	routers.InitRoutes(engine)

	log.Print("web is running...")
	engine.Run(*addr)
}
