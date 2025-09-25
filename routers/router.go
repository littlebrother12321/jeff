package routers

import (
	"firstbee/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/about", &controllers.AboutController{})
	beego.Router("/test", &controllers.TestController{})
	beego.Router("/profile", &controllers.ProfileController{})
	beego.Router("/admin", &controllers.AdminController{})
	beego.Router("/admin/login", &controllers.LoginController{})
	beego.Router("/admin/logout", &controllers.LogoutController{})
	beego.Router("/admin/about/:id", &controllers.AboutDetailController{})
	beego.Router("/admin/about", &controllers.AboutListController{})

}
