package models

import (
	"log"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/mattn/go-sqlite3"
)

type AboutModel struct {
	Id      uint64 `orm:"auto"`        // this automatically creates an integer primary key
	Name    string `orm:"size(100)"`   // 100 characters max
	Email   string `orm:"size(255)"`   // 255 characters max
	Message string `form:"type(text)"` // any size string
}

var O orm.Ormer

func InitDB() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "./public.db.db")
	// this function can take a list, e.g. orm.RegisterModel(new(M1), new(M2), ...)
	orm.RegisterModel(new(AboutModel), new(User))
	O = orm.NewOrm()

	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		log.Fatalf("Failed to sync database: %v", err)
	}
}

type Race struct {
	Id   int64  `orm:"auto;pk"`
	Name string `orm:"size(100)"`
	Code string `orm:"size(100)"`
}

type Team struct {
	Id   uint64 `orm:"auto"`
	Name string `orm:"size(100)"`
}

type User struct {
	Id       uint64 `orm:"auto"` // this automatically creates an integer primary key
	Name     string `orm:"size(100)"`
	Email    string `orm:"size(255);unique"`
	Password string `orm:"size(255)"`
}


type TeamMember struct {
	Id     uint64 `orm:"auto"`
	Name   string `orm:"size(100)"`
	Gender string `orm:"size(20)`
	Grade  uint8
	Team   *Team `orm:"rel(fk)"`
}

func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "./stopwatch.db")
	orm.RegisterModel(new(Race))
}
