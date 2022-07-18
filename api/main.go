package main

import (
	"flag"
	"homesite/database"
	"homesite/middlewares"
	"homesite/routers"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/thinkerou/favicon"
)

func main() {
	var addr = flag.String("a", ":8000", "listening address")
	flag.Parse()
	engine := gin.Default()

	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(middlewares.TokenAuthMiddleware())
	engine.Use(favicon.New("./favicon.ico"))
	engine.Static("/assets", "./assets/")
	engine.LoadHTMLGlob("templates/*")

	routers.InitRoutes(engine)

	// init db
	if err := database.Init(); err != nil {
		log.Printf(err.Error())
	}

	log.Print("web is running...")
	engine.Run(*addr)
}
