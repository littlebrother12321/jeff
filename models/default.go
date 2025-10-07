package models

import (
	"log"
	"os"
	"fmt"
	
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
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
	driverName := os.Getenv("DATABASE_DRIVER")
	dataSource := os.Getenv("DATABASE_URL")
	driver := orm.DRSqlite
	if driverName == "postgres" {
		driver = orm.DRPostgres
	} else {
		driverName = "sqlite3"
		dataSource = "./public.db.db"
	}
	orm.RegisterDriver(driverName, driver)
	orm.RegisterDataBase("default", driverName, dataSource)
	// this function can take a list, e.g. orm.RegisterModel(new(M1), new(M2), ...)
	orm.RegisterModel(new(AboutModel), new(User))
	O = orm.NewOrm()

	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		log.Fatalf("Failed to sync database: %v", err)
	}

	if driverName == "postgres" {
		_, err = O.Raw(`CREATE TABLE IF NOT EXISTS "session") (
			"session_key" CHAR(64) NOT NULL,
			"session_data" BYTEA,
			"session_expiry" BIGINT NOT NULL,
			PRIMARY KEY ("session_key")
		);`).Exec()
		if err != nil {
			fmt.Println("table `session` created")
		} else {
			fmt.Println("`session` table not created, application may not work")
		}
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
	orm.RegisterDataBase("default", "sqlite3", "./public.db.db")
	orm.RegisterModel(new(Race))
}
