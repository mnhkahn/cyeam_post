package main

import (
	"cyeam_post/bot"
	"cyeam_post/conf"
	"cyeam_post/dao"
	. "cyeam_post/logs"
	"cyeam_post/parser"
	"cyeam_post/utils"
	"time"
)

var bots_parser map[bot.Bot]parser.Parser

func Process() {
	utils.ColorLog("[INFO] Start parseing at %s...\n", time.Now().Format(time.RFC3339))
	bot, err := bot.NewBot("CyBot")
	if err != nil {
		panic(err)
	}
	Dao, err := dao.NewDao("solr", conf.String("solr.host"))
	Dao.Debug(conf.Bool("dao.debug"))
	if err != nil {
		panic(err)
	}
	parser, err := parser.NewParser("CyParser")
	parser.Debug(conf.Bool("parser.debug"))

	bot.Init(parser, Dao)
	bot.Debug(conf.Bool("bot.debug"))
	bot.Start(conf.String("root"))
	utils.ColorLog("[SUCC] Parse successful\n")
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
