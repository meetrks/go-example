package main

import (
	"gin-example/migrations"
	"gin-example/models"
	"gin-example/routers"
)

func init() {
	models.Setup()
	migrations.RunMigration()
	routers.InitRouter()
}

func main() {
	//router :=
	//router.Run()
}
