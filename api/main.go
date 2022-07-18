package main

import (
	"database/sql"
	"flag"
	"homesite/middlewares"
	"homesite/routers"
	"log"
	"time"

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

	db, err := sql.Open("mysql", "root:kkk123@tcp(localhost)/home")
	if err == nil {
		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)
	}

	log.Print("web is running...")
	engine.Run(*addr)
}
