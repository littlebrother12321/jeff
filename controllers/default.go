package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"firstbee/models"
	"log"
	"net/smtp"
	"strings"
	"github.com/joho/godotenv"
	"os"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Title"] = "hello"
	c.TplName = "index.tpl"
}

type AboutController struct {
	beego.Controller
}

func (c *AboutController) Get() {
	c.Data["Title"] = "about"
	c.TplName = "about.tpl"
}

type About struct {
        Name string  `form:"name"`
	Email string `form:"email"`
	Message string `form:"message"`
}

func (c *AboutController) Post() {
        c.Data["Title"] = "About"
	c.Data["Result"] = "Thanks for ur data!"
	c.TplName = "about.tpl"

	// handle form data or somethin
	contact := About{}
	err := c.Ctx.BindForm(&contact) //pass pointer to struct
	if err != nil {
		c.Data["Result"] = "ERROR: " + err.Error()
	} else if contact.Message == "" || contact.Name == "" || contact.Email == "" {	
		c.Data["Result"] = "ERROR: Please enter all values."
	} else {
		err = sendEmail(contact)
		if err != nil {
			c.Data["Result"] = "ERROR: Email stoopid"
		} else {
			c.Data["Result"] = "ERROR: Email sent successfully"
		}
		c.Data["Result"] = "Thanks for ur data!"
		log.Default().Println(contact)

		aboutDb := models.AboutModel{
			Name: contact.Name,
			Email: contact.Email,
			Message: contact.Message,
		}
		_, err := models.O.Insert(&aboutDb)
		if err != nil {
			c.Data["Result"] = "ERR: cOOLD NOT SAVE DBASE!:" + err.Error()
		} else {
			c.Data["Result"] = "Contact form recevved!"
		}
	}

}

// sendEmail sends an email using SMTP
func sendEmail(contact About) error {
	err2 := godotenv.Load()
	if err2 != nil {
		log.Fatal("error loading .env!")
	}
	from := os.Getenv("EMAIL") // Replace with your email
	password := os.Getenv("PASSWORD")    // Replace with your email password
	//NEW! using .env file!
	toArr := [...]string{contact.Email} // Replace with the recipient's email(s)
	// slice the array
	to := toArr[:]

	// Set up authentication information.
	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com") // Replace with your SMTP server

	// Compose the email message
	subject := "New Contact Form Submission"
	body := "Name: " + contact.Name + "\nEmail: " + contact.Email + "\nMessage: " + contact.Message
	message := []byte("To: " + strings.Join(to[:], ",") + "\r\n" +
		"From: " + from + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + // This empty line separates the headers from the body
		body)

	// Send the email
	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, message) // Replace with your SMTP server and port
	return err
}


type TestController struct {
     	beego.Controller
}

func (c *TestController) Get() {
        c.Data["Title"] = "test"
	c.TplName = "index.tpl"
}    

type ProfileController struct {
	BaseAdminController
}

func (c *ProfileController) Get() {
	c.RequireAuth()
	c.Data["Title"] = "Profile"
	c.TplName = "profile.tpl"
}

type EditProfile struct {
	Name  string `form:"name"`
	Email string `form:"email"`
}

func (c *ProfileController) Post() {
	c.RequireAuth()
	c.Data["Title"] = "Profile"
	c.TplName = "profile.tpl"

	ep := EditProfile{}
	err := c.Ctx.BindForm(&ep)
	if err != nil {
		c.Data["Result"] = "ERROR: " + err.Error()
	} else if ep.Email == "" || ep.Name == "" {
		c.Data["Result"] = "ERROR: Please enter all values."
	} else {
		user := c.GetCurrentUser()
		user.Name = ep.Name
		user.Email = ep.Email
		_, err := models.O.Update(user, "Name", "Email")
		if err != nil {
			c.Data["Result"] = "ERROR: " + err.Error()
		} else {
			// update the user in the session
			c.SetSession("user", *user)
			c.Data["Result"] = "Profile updated!"
		}
	}

}
