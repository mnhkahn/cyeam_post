package main

import (
	"cyeam_post/dao"
	"cyeam_post/parser"
	"fmt"
	"github.com/astaxie/beego/config"
	"strings"
	"time"
)

var AppConfig config.ConfigContainer

func Process() {
	fmt.Println("Start parse============", time.Now())
	Dao, _ := dao.NewDao("db", "cyeam:qwerty@tcp(128.199.131.129:3306)/cyeam?charset=utf8")
	Dao.Debug(AppConfig.String("runmode") == "dev")
	// fmt.Println(Dao.GetPost("mangoer_ys", "", 1, 0))
	rss_list := strings.Split(AppConfig.String("rss.source"), ";")
	for _, rss := range rss_list {
		P, err := parser.NewParser("rss", rss)
		if err != nil {
			panic(err)
		}
		for i := 0; i < P.Len(); i++ {
			if Dao.IsPostUpdate(P.Index(i)) {
				fmt.Println("upate**************", P.Index(i).Title)
				Dao.UpdatePost(P.Index(i))
			} else if Dao.GetPostByLink(P.Index(i).Link).Title == "" {
				fmt.Println("add****************", P.Index(i).Title)
				Dao.AddPost(P.Index(i))
			} else {
				fmt.Println("already exists & not changed**************", P.Index(i).Title)
			}
		}
	}
}

func timer() {
	duration := AppConfig.DefaultInt("parse.duration", 60)
	timer := time.NewTicker(time.Duration(duration) * time.Minute)
	for {
		select {
		case <-timer.C:
			go func() {
				Process()
			}()
		}
	}
}

func main() {
	timer()
	// Process()
}

func init() {
	var err error
	AppConfig, err = config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}
}
