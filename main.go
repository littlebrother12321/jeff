package main

import (
	"log"
	"firstbee/models"
	_ "firstbee/routers"
	
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/beego/beego/v2/server/web/session/postgres"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file!")
	}
	//beego.LoadAppConfig("ini", "conf/app.conf")
	models.InitDB()
	beego.Run()
}
