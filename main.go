package main

import (
	"cyeam_post/bot"
	"cyeam_post/dao"
	. "cyeam_post/logs"
	"cyeam_post/parser"
	"fmt"
	"github.com/astaxie/beego/config"
	"time"
)

var AppConfig config.ConfigContainer

var bots_parser map[bot.Bot]parser.Parser

func Process() {
	fmt.Println("Start parse==========")
	bot, err := bot.NewBot("RssBot")
	if err != nil {
		panic(err)
	}
	Dao, err := dao.NewDao("solr", "http://128.199.131.129:8983/solr/post")
	// Dao.Debug(true)
	if err != nil {
		panic(err)
	}
	parser, err := parser.NewParser("RssParser")

	bot.Init(parser, Dao)
	bot.Debug(true)
	bot.Start(AppConfig.String("rss.source"))

	fmt.Println("End parse==========")
	Log.Close()
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
