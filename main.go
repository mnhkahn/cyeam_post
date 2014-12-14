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
	Dao, _ := dao.NewDao("solr", "http://128.199.131.129:8983/solr/post")
	Dao.Debug(AppConfig.String("runmode") == "dev")
	rss_list := strings.Split(AppConfig.String("rss.source"), ";")
	for _, rss := range rss_list {
		P, err := parser.NewParser("rss", rss)
		if err != nil {
			panic(err)
		}
		for i := 0; i < P.Len(); i++ {
			Dao.AddOrUpdate(P.Index(i))
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
	Process()
	timer()
}

func init() {
	var err error
	AppConfig, err = config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}
}
