package controllers

import (
	"firstbee/models"
	"firstbee/utils"

	beego "github.com/beego/beego/v2/server/web"
)

type BaseAdminController struct {
	beego.Controller
}

// Prepare runs before every request
func (c *BaseAdminController) Prepare() {
	// Set common template data
	c.setCommonData()
}

// Set common data for all templates
func (c *BaseAdminController) setCommonData() {
	// Check if user is logged in and set template variables
	if c.IsLoggedIn() {
		user := c.GetCurrentUser()
		c.Data["IsLoggedIn"] = true
		c.Data["User"] = user
	} else {
		c.Data["IsLoggedIn"] = false
		c.Data["User"] = nil
	}
}

// Check if user is logged in
func (c *BaseAdminController) IsLoggedIn() bool {
	user := c.GetSession("user")
	return user != nil
}

// Get current user from session
func (c *BaseAdminController) GetCurrentUser() *models.User {
	userSession := c.GetSession("user")
	if userSession == nil {
		return nil
	}

	// Assuming you store user info in session
	if user, ok := userSession.(models.User); ok {
		return &user
	}
	return nil
}

// RequireAuth middleware - add this to controllers that need authentication
func (c *BaseAdminController) RequireAuth() {
	if !c.IsLoggedIn() {
		// Store the current URL for redirect after login
		c.SetSession("redirect_after_login", c.Ctx.Request.URL.Path)
		c.Redirect("/admin/login", 302)
		return
	}
}

type AdminController struct {
	BaseAdminController
}

func (c *AdminController) Get() {
	c.RequireAuth()
	c.Data["Title"] = "Admin"
	c.TplName = "admin/index.tpl"
}

type LogoutController struct {
	BaseAdminController
}

func (c *LogoutController) Get() {
	c.DestroySession()
	c.Redirect("/admin/login", 302)
}

type LoginController struct {
	BaseAdminController
}

func (c *LoginController) Get() {
	c.Data["Title"] = "Login"
	c.TplName = "admin/login.tpl"
}

func (c *LoginController) Post() {
	c.Data["Title"] = "Login"
	c.TplName = "admin/login.tpl"

	loginreq := utils.LoginReq{}
	err := c.Ctx.BindForm(&loginreq)
	if err != nil {
		c.Data["Result"] = "ERROR: " + err.Error()
	} else if loginreq.Email == "" || loginreq.Password == "" {
		c.Data["Result"] = "ERROR: Please enter all values."
	} else {
		// authenticate
		user, err := utils.Authenticate(&loginreq)
		if err != nil {
			c.Data["Result"] = "ERROR: " + err.Error()
		} else {
			// create session
			c.SetSession("user", *user)

			// redirect
			path := c.GetSession("redirect_after_login")
			if path != nil {
				if path, ok := path.(string); ok {
					c.Redirect(path, 302)
				}
			}
			c.Redirect("/admin", 302)
		}
	}
}
