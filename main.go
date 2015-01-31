package main

import (
	"cyeam_post/bot"
	"cyeam_post/conf"
	"cyeam_post/dao"
	. "cyeam_post/logs"
	"cyeam_post/parser"
	"fmt"
	"time"
)

var bots_parser map[bot.Bot]parser.Parser

func Process() {
	fmt.Println("Start parse==========")
	bot, err := bot.NewBot("CyBot")
	if err != nil {
		panic(err)
	}
	Dao, err := dao.NewDao("solr", conf.String("solr.host"))
	// Dao.Debug(true)
	if err != nil {
		panic(err)
	}
	parser, err := parser.NewParser("CyParser")

	bot.Init(parser, Dao)
	bot.Debug(true)
	bot.Start(conf.String("root"))

	fmt.Println("End parse==========")
	Log.Close()
}

func timer() {
	duration := conf.DefaultInt("parse.duration", 60)
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
