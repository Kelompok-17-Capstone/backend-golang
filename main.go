package main

import (
	"backend-golang/config"

	"github.com/labstack/echo/v4"
)

func init() {
	config.InitDB()
	config.InitialMigration()

}

func main() {
	e := echo.New()
	e.Logger.Fatal(e.Start(":8080"))
}
