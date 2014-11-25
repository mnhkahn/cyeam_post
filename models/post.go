package models

import (
	"time"
)

type Post struct {
	Id          int       `xorm:"not null pk autoincr INT(11)"`
	Title       string    `xorm:"not null VARCHAR(200)"`
	CreateTime  time.Time `xorm:"not null DATETIME"`
	Author      string    `xorm:"VARCHAR(45)"`
	Detail      string    `xorm:"not null LONGTEXT"`
	Category    string    `xorm:"VARCHAR(45)"`
	Tags        string    `xorm:"VARCHAR(45)"`
	Figure      string    `xorm:"VARCHAR(100)"`
	Description string    `xorm:"TINYTEXT"`
	Link        string    `xorm:"not null VARCHAR(100)"`
	Source      string    `xorm:"VARCHAR(40)"`
}
