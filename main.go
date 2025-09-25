package main

import (
	"firstbee/models"
	_ "firstbee/routers"
	
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.LoadAppConfig("ini", "conf/app.conf")
	models.InitDB()
	beego.Run()
}
