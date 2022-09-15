package main

import (
	"Test/config"
	"Test/factory"
	"Test/features/middlewares"
	"Test/utils/mysql"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.Getconfig()
	db := mysql.InitDB(cfg)

	e := echo.New()

	factory.InitFactory(e, db)

	fmt.Println("==== STARTING PROGRAM ====")
	address := fmt.Sprintf(":%d", config.SERVERPORT)
	middlewares.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(address))
}
