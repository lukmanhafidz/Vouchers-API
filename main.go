package main

import (
	"Test/config"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	fmt.Println("==== STARTING PROGRAM ====")
	address := fmt.Sprintf(":%d", config.SERVERPORT)
	e.Logger.Fatal(e.Start(address))
}
