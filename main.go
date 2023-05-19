package main

import (
	"backend-golang/config"
	"backend-golang/models/seeder"
	"backend-golang/routes"
	"backend-golang/util"

	"github.com/go-playground/validator"
)

func init() {
	config.InitDB()
	config.InitialMigration()
	seeder.DBSeed(config.DB)
}

func main() {
	e := routes.New()
	e.Validator = &util.CustomValidator{
		Validator: validator.New(),
	}
	e.Logger.Fatal(e.Start(":8080"))
}
