package bot

import (
	"cyeam_post/dao"
	"cyeam_post/parser"
	"testing"
)

// func TestRss(t *testing.T) {
// 	bot, err := NewBot("RssBot")
// 	if err != nil {
// 		panic(err)
// 	}
// 	Dao, err := dao.NewDao("solr", "http://128.199.131.129:8983/solr/post")
// 	Dao.Debug(true)
// 	if err != nil {
// 		panic(err)
// 	}
// 	parser, err := parser.NewParser("RssParser")

// 	bot.Init(parser, Dao)
// 	bot.Debug(true)
// 	bot.Start("http://blog.cyeam.com/rss.xml")
// }

func TestCy(t *testing.T) {
	bot, err := NewBot("CyBot")
	if err != nil {
		panic(err)
	}
	Dao, err := dao.NewDao("solr", "http://128.199.131.129:8983/solr/post")
	Dao.Debug(true)
	if err != nil {
		panic(err)
	}
	parser, err := parser.NewParser("CyParser")

	bot.Init(parser, Dao)
	bot.Debug(true)
	bot.Start("http://localhost:8080")
}
